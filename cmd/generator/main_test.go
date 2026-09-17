package main

import (
	"encoding/json"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func writeSchema(t *testing.T, path string, schema Schema) {
	t.Helper()
	data, err := json.Marshal(schema)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, data, 0644); err != nil {
		t.Fatal(err)
	}
}

func TestGeneratedYearData(t *testing.T) {
	root := t.TempDir()
	dataDir := filepath.Join(root, "data")
	outputDir := filepath.Join(root, "pkg", "holiday")
	for _, dir := range []string{dataDir, outputDir} {
		if err := os.MkdirAll(dir, 0755); err != nil {
			t.Fatal(err)
		}
	}
	// Deliberately visit later arrangement years first to test precedence.
	fixtures := []Schema{
		{Year: 2020, Days: []arrangementDay{}},
		{Year: 2019, Days: []arrangementDay{
			{Name: "New Year", Date: "2018-12-29", IsOffDay: false},
			{Name: "New Year", Date: "2018-12-31", IsOffDay: true},
			{Name: "New Year", Date: "2019-01-01", IsOffDay: true},
		}},
		{Year: 2018, Days: []arrangementDay{
			{Name: "Superseded", Date: "2018-12-29", IsOffDay: true},
			{Name: "Current arrangement", Date: "2018-01-03", IsOffDay: false},
		}},
	}
	for i, schema := range fixtures {
		writeSchema(t, filepath.Join(dataDir, string(rune('a'+i))+".json"), schema)
	}
	yearFile := filepath.Join(outputDir, "year_2018.go")
	if err := generate(dataDir, outputDir); err != nil {
		t.Fatal(err)
	}
	first, err := os.ReadFile(yearFile)
	if err != nil {
		t.Fatal(err)
	}
	if err := generate(dataDir, outputDir); err != nil {
		t.Fatal(err)
	}
	second, err := os.ReadFile(yearFile)
	if err != nil {
		t.Fatal(err)
	}
	if string(first) != string(second) {
		t.Fatal("year generation is not deterministic")
	}

	// Compile the generated output with the real public helper implementation.
	for _, name := range []string{"go.mod", "holidaycn.go"} {
		content, err := os.ReadFile(filepath.Join("..", "..", name))
		if err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(root, name), content, 0644); err != nil {
			t.Fatal(err)
		}
	}
	const testSource = `package holidaycn

import (
	"testing"
	"time"

	"github.com/washanhanzi/holiday-cn-go/pkg/holiday"
)

func TestCalendar(t *testing.T) {
	if data := holiday.GetYearData(2020); data != nil {
		t.Errorf("empty source year returned %v, want nil", data)
	}
	if _, err := CheckHoliday(time.Date(2020, 1, 1, 0, 0, 0, 0, cnLocation)); err == nil {
		t.Error("empty source year should be unsupported")
	}
	tests := []struct {
		date string
		off bool
		name string
	}{
		{"2018-12-29", false, "New Year"},
		{"2018-12-31", true, "New Year"},
		{"2018-01-03", false, "Current arrangement"},
		{"2018-12-28", false, ""},
		{"2018-12-30", true, ""},
		{"2019-01-01", true, "New Year"},
		{"2019-01-06", true, ""},
	}
	for _, tt := range tests {
		t.Run(tt.date, func(t *testing.T) {
			date, err := time.ParseInLocation("2006-01-02", tt.date, cnLocation)
			if err != nil { t.Fatal(err) }
			off, name, err := IsRestDay(date)
			if err != nil || off != tt.off || name != tt.name {
				t.Errorf("got (%v, %q, %v), want (%v, %q, nil)", off, name, err, tt.off, tt.name)
			}
		})
	}
	// The existing year API exposes the record merged during generation.
	if off, name, err := holiday.IsHoliday2018("2018-12-29"); err != nil || off || name != "New Year" {
		t.Errorf("missing merged record: (%v, %q, %v)", off, name, err)
	}
	if _, exists := holiday.GetYearData(2018)["2019-01-01"]; exists {
		t.Error("2018 imported a date belonging to 2019")
	}
	for _, year := range []int{2018, 2019} {
		if day := holiday.GetYearData(year)["2018-12-29"]; day.ArrangementYear != 2019 {
			t.Errorf("year %d lost original arrangement year: %+v", year, day)
		}
	}
	if day := holiday.GetYearData(2018)["2018-01-03"]; day.ArrangementYear != 2018 {
		t.Errorf("incorrect arrangement year for original record: %+v", day)
	}
}
`
	if err := os.WriteFile(filepath.Join(root, "calendar_test.go"), []byte(testSource), 0644); err != nil {
		t.Fatal(err)
	}
	cmd := exec.Command("go", "test", "./...")
	cmd.Dir = root
	cmd.Env = append(os.Environ(), "GOWORK=off")
	if output, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("generated package failed: %v\n%s", err, output)
	}
}

func TestGenerateRejectsInvalidDate(t *testing.T) {
	root := t.TempDir()
	writeSchema(t, filepath.Join(root, "2019.json"), Schema{
		Year: 2019,
		Days: []arrangementDay{{Date: "2018-12-32"}},
	})
	err := generate(root, root)
	if err == nil || !strings.Contains(err.Error(), "invalid date") {
		t.Fatalf("generate() = %v, want invalid date error", err)
	}
}

func TestGenerateSkipsEmptyYears(t *testing.T) {
	for _, source := range []string{
		`{"year":2018,"days":[]}`,
		`{"year":2018,"days":null}`,
		`{"year":2018}`,
	} {
		t.Run(source, func(t *testing.T) {
			root := t.TempDir()
			path := filepath.Join(root, "2018.json")
			writeSchema(t, path, Schema{Year: 2018, Days: []arrangementDay{{Date: "2018-01-01"}}})
			if err := generate(root, root); err != nil {
				t.Fatal(err)
			}
			if _, err := os.Stat(filepath.Join(root, "year_2018.go")); err != nil {
				t.Fatal(err)
			}
			if err := os.WriteFile(path, []byte(source), 0644); err != nil {
				t.Fatal(err)
			}
			// A neighboring year's cross-year entry must not revive an empty year.
			writeSchema(t, filepath.Join(root, "2019.json"), Schema{
				Year: 2019, Days: []arrangementDay{{Date: "2018-12-31"}},
			})
			for i := 0; i < 2; i++ {
				if err := generate(root, root); err != nil {
					t.Fatal(err)
				}
				if _, err := os.Stat(filepath.Join(root, "year_2018.go")); !errors.Is(err, os.ErrNotExist) {
					t.Fatalf("empty year file still exists or stat failed: %v", err)
				}
				registry, err := os.ReadFile(filepath.Join(root, "holiday.go"))
				if err != nil {
					t.Fatal(err)
				}
				if strings.Contains(string(registry), "Init2018") || !strings.Contains(string(registry), "Init2019") {
					t.Fatal("registry must include only the nonempty year")
				}
			}
		})
	}
}

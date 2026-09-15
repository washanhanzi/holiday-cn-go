package holiday_test

import (
	"testing"

	"github.com/washanhanzi/holiday-cn-go/pkg/holiday"
)

func TestYearAPIsIncludeCrossYearRecords(t *testing.T) {
	const date = "2018-12-29"
	want := holiday.Day{Name: "元旦", Date: date, IsOffDay: false}
	if got := holiday.GetYearData(2019)[date]; got != want {
		t.Errorf("GetYearData(2019)[%q] = %+v, want %+v", date, got, want)
	}
	if got := holiday.GetYearData(2018)[date]; got != want {
		t.Errorf("GetYearData(2018)[%q] = %+v, want %+v", date, got, want)
	}

	tests := []struct {
		date    string
		wantOff bool
	}{
		{date, false},
		{"2018-12-31", true},
		{"2019-01-01", true},
	}
	for _, tt := range tests {
		t.Run(tt.date, func(t *testing.T) {
			off, name, err := holiday.IsHoliday2019(tt.date)
			if err != nil || off != tt.wantOff || name != "元旦" {
				t.Errorf("IsHoliday2019(%q) = (%v, %q, %v), want (%v, %q, nil)", tt.date, off, name, err, tt.wantOff, "元旦")
			}
		})
	}

	// The existing year-specific API also sees records merged during generation.
	if off, name, err := holiday.IsHoliday2018("2018-12-31"); err != nil || !off || name != "元旦" {
		t.Errorf("IsHoliday2018(2018-12-31) = (%v, %q, %v), want (true, 元旦, nil)", off, name, err)
	}
	// A regular Sunday absent from the file does not get weekend fallback.
	if off, name, err := holiday.IsHoliday2019("2019-01-06"); err != nil || off || name != "" {
		t.Errorf("IsHoliday2019(2019-01-06) = (%v, %q, %v), want (false, empty, nil)", off, name, err)
	}
}

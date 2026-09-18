package holiday

import (
	"reflect"
	"sync"
	"testing"
	"time"
)

func TestMutateDay(t *testing.T) {
	const year = 2000
	date := time.Date(year, 1, 1, 0, 0, 0, 0, time.FixedZone("Test", 8*60*60))
	original := Day{ArrangementYear: 1999, Date: "2000-01-01", Name: "Original", IsOffDay: true}
	other := Day{Date: "2000-01-02", Name: "Other"}
	calls := 0
	yearData[year] = &lazyYearData{init: func() map[string]Day {
		calls++
		return map[string]Day{original.Date: original, other.Date: other}
	}}
	t.Cleanup(func() { delete(yearData, year) })
	var retained *Day
	if err := MutateDay("2000-01-01", func(day *Day) {
		if calls != 1 || *day != original {
			t.Fatal("cache was not populated before mutation")
		}
		retained = day
		day.Name = "Custom"
		day.IsOffDay = false
		if GetYearData(year)[original.Date] != original {
			t.Fatal("unpublished mutation visible to readers")
		}
	}); err != nil {
		t.Fatal(err)
	}
	retained.Name = "Retained pointer"
	want := original
	want.Name = "Custom"
	want.IsOffDay = false
	if day, err := CheckHoliday(date); err != nil || day == nil || *day != want {
		t.Fatalf("updated record = (%v, %v), want %v", day, err, want)
	}
	if GetYearData(year)[other.Date] != other {
		t.Fatal("mutation changed another record")
	}
	if err := MutateDay("2000-01-03", func(day *Day) {
		if *day != (Day{Date: "2000-01-03", ArrangementYear: year}) {
			t.Fatalf("unexpected new record defaults: %+v", day)
		}
		day.IsOffDay = true
	}); err != nil {
		t.Fatal(err)
	}
	if calls != 1 || !GetYearData(year)["2000-01-03"].IsOffDay {
		t.Fatal("insert failed or reinitialized the cache")
	}
	before := GetYearData(year)
	if err := MutateDay("2000-01-01", func(day *Day) { day.Date = "2000-01-02" }); err == nil {
		t.Fatal("date change accepted")
	}
	func() {
		defer func() {
			if recover() == nil {
				t.Error("expected callback panic")
			}
		}()
		MutateDay("2000-01-01", func(day *Day) {
			day.Name = "Aborted"
			panic("abort")
		})
	}()
	if !reflect.DeepEqual(GetYearData(year), before) {
		t.Fatal("failed callback changed cache")
	}
	if err := MutateDay("2000-01-01", func(*Day) {}); err != nil {
		t.Fatal(err)
	}
	if err := MutateDay("2000-01-01", nil); err == nil {
		t.Fatal("nil callback accepted")
	}
	if err := MutateDay("1999-01-01", func(*Day) { t.Fatal("unsupported callback invoked") }); err == nil {
		t.Fatal("unsupported year accepted")
	}
}

func TestMutateDayConcurrent(t *testing.T) {
	const year = 2000
	date := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
	yearData[year] = &lazyYearData{init: func() map[string]Day {
		return map[string]Day{"2000-01-01": {Date: "2000-01-01"}}
	}}
	t.Cleanup(func() { delete(yearData, year) })
	var wg sync.WaitGroup
	for i := 0; i < 64; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := MutateDay("2000-01-01", func(day *Day) {
				day.ArrangementYear++
			}); err != nil {
				t.Error(err)
			}
			GetYearData(year)
			if _, err := CheckHoliday(date); err != nil {
				t.Error(err)
			}
		}()
	}
	wg.Wait()
	if got := GetYearData(year)["2000-01-01"].ArrangementYear; got != 64 {
		t.Fatalf("lost concurrent updates: got %d, want 64", got)
	}
}

func TestMutateDayDateValidation(t *testing.T) {
	const year = 2000
	calls := 0
	yearData[year] = &lazyYearData{init: func() map[string]Day {
		calls++
		return make(map[string]Day)
	}}
	t.Cleanup(func() { delete(yearData, year) })
	for _, date := range []string{
		"", "2000-1-01", "2000-01-1", "2000-02-30", "2001-02-29",
		"2000-13-01", "2000-01-00", " 2000-01-01", "2000-01-01 ",
		"2000-01-01T00:00:00+08:00",
	} {
		t.Run(date, func(t *testing.T) {
			if err := MutateDay(date, func(*Day) { t.Fatal("invalid date callback invoked") }); err == nil {
				t.Fatal("invalid date accepted")
			}
		})
	}
	if calls != 0 {
		t.Fatal("invalid input initialized the cache")
	}
	if err := MutateDay("2000-02-29", func(day *Day) {
		day.Name = "Leap day"
		day.IsOffDay = true
	}); err != nil {
		t.Fatal(err)
	}
	want := Day{Date: "2000-02-29", ArrangementYear: year, Name: "Leap day", IsOffDay: true}
	if got := GetYearData(year)[want.Date]; got != want {
		t.Fatalf("leap day record = %+v, want %+v", got, want)
	}
}

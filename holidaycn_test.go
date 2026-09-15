package holidaycn

import (
	"strings"
	"testing"
	"time"
)

func TestIsRestDay(t *testing.T) {
	tests := []struct {
		name            string
		date            time.Time
		wantIsRestDay   bool
		wantHolidayName string
		wantErr         bool
	}{
		{
			name:            "New Year 2025",
			date:            time.Date(2025, 1, 1, 0, 0, 0, 0, cnLocation),
			wantIsRestDay:   true,
			wantHolidayName: "元旦",
			wantErr:         false,
		},
		{
			name:            "Regular Workday",
			date:            time.Date(2025, 1, 2, 0, 0, 0, 0, cnLocation),
			wantIsRestDay:   false,
			wantHolidayName: "",
			wantErr:         false,
		},
		{
			name:            "Regular Weekend",
			date:            time.Date(2025, 1, 4, 0, 0, 0, 0, cnLocation), // Saturday
			wantIsRestDay:   true,
			wantHolidayName: "",
			wantErr:         false,
		},
		{
			name:    "Unknown date before first arrangement year",
			date:    time.Date(2006, 12, 29, 0, 0, 0, 0, cnLocation),
			wantErr: true,
		},
		{
			name:    "Cross-year entry outside supported years",
			date:    time.Date(2006, 12, 30, 0, 0, 0, 0, cnLocation),
			wantErr: true,
		},
		{
			name:    "Second cross-year entry outside supported years",
			date:    time.Date(2006, 12, 31, 0, 0, 0, 0, cnLocation),
			wantErr: true,
		},
		{
			name:            "Future Year",
			date:            time.Date(2030, 1, 1, 0, 0, 0, 0, cnLocation),
			wantIsRestDay:   false,
			wantHolidayName: "",
			wantErr:         true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIsRestDay, gotHolidayName, err := IsRestDay(tt.date)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsRestDay() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotIsRestDay != tt.wantIsRestDay {
				t.Errorf("IsRestDay() gotIsRestDay = %v, want %v", gotIsRestDay, tt.wantIsRestDay)
			}
			if gotHolidayName != tt.wantHolidayName {
				t.Errorf("IsRestDay() gotHolidayName = %v, want %v", gotHolidayName, tt.wantHolidayName)
			}
		})
	}
}

func TestCrossYearDays(t *testing.T) {
	// These December dates are stored in the following year's arrangements.
	tests := []struct {
		date    string
		wantOff bool
	}{
		{"2007-12-29", false},
		{"2007-12-30", true},
		{"2007-12-31", true},
		{"2011-12-31", false},
		{"2018-12-29", false},
		{"2018-12-30", true},
		{"2018-12-31", true},
		{"2022-12-31", true},
	}

	for _, tt := range tests {
		t.Run(tt.date, func(t *testing.T) {
			date, err := time.ParseInLocation("2006-01-02", tt.date, cnLocation)
			if err != nil {
				t.Fatal(err)
			}
			off, name, err := IsRestDay(date)
			if err != nil || off != tt.wantOff || name != "元旦" {
				t.Errorf("IsRestDay(%s) = (%v, %q, %v), want (%v, %q, nil)", tt.date, off, name, err, tt.wantOff, "元旦")
			}
			work, err := IsWorkday(date)
			if err != nil || work != !tt.wantOff {
				t.Errorf("IsWorkday(%s) = (%v, %v), want (%v, nil)", tt.date, work, err, !tt.wantOff)
			}
		})
	}
}

func TestIsWorkday(t *testing.T) {
	tests := []struct {
		name          string
		date          time.Time
		wantIsWorkday bool
		wantErr       bool
	}{
		{
			name:          "Regular Workday",
			date:          time.Date(2025, 1, 2, 0, 0, 0, 0, cnLocation),
			wantIsWorkday: true,
			wantErr:       false,
		},
		{
			name:          "Holiday",
			date:          time.Date(2025, 1, 1, 0, 0, 0, 0, cnLocation),
			wantIsWorkday: false,
			wantErr:       false,
		},
		{
			name:          "Weekend",
			date:          time.Date(2025, 1, 4, 0, 0, 0, 0, cnLocation), // Saturday
			wantIsWorkday: false,
			wantErr:       false,
		},
		{
			name:          "Future Year",
			date:          time.Date(2030, 1, 1, 0, 0, 0, 0, cnLocation),
			wantIsWorkday: false,
			wantErr:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIsWorkday, err := IsWorkday(tt.date)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsWorkday() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotIsWorkday != tt.wantIsWorkday {
				t.Errorf("IsWorkday() = %v, want %v", gotIsWorkday, tt.wantIsWorkday)
			}
		})
	}
}

func TestAfterWorkdays(t *testing.T) {
	tests := []struct {
		name        string
		startDate   time.Time
		workdays    int
		wantDate    time.Time
		wantErr     bool
		errContains string
	}{
		{
			name:      "One workday, next day is workday",
			startDate: time.Date(2025, 1, 2, 0, 0, 0, 0, cnLocation), // Thursday
			workdays:  1,
			wantDate:  time.Date(2025, 1, 6, 0, 0, 0, 0, cnLocation), // Next workday after counting 1 workday (Friday)
			wantErr:   false,
		},
		{
			name:      "One workday, starting Friday",
			startDate: time.Date(2025, 1, 3, 0, 0, 0, 0, cnLocation), // Friday
			workdays:  1,
			wantDate:  time.Date(2025, 1, 7, 0, 0, 0, 0, cnLocation), // Next workday after counting 1 workday (Monday)
			wantErr:   false,
		},
		{
			name:      "Two workdays starting Friday",
			startDate: time.Date(2025, 1, 3, 0, 0, 0, 0, cnLocation), // Friday
			workdays:  2,
			wantDate:  time.Date(2025, 1, 8, 0, 0, 0, 0, cnLocation), // Next workday after counting 2 workdays (Tue)
			wantErr:   false,
		},
		{
			name:      "One workday during Chinese New Year",
			startDate: time.Date(2025, 1, 27, 0, 0, 0, 0, cnLocation), // Monday before CNY
			workdays:  1,
			wantDate:  time.Date(2025, 2, 6, 0, 0, 0, 0, cnLocation), // Next workday after counting 1 workday during CNY
			wantErr:   false,
		},
		{
			name:      "Next workday is a cross-year makeup Saturday",
			startDate: time.Date(2018, 12, 28, 0, 0, 0, 0, cnLocation),
			workdays:  0,
			wantDate:  time.Date(2018, 12, 29, 0, 0, 0, 0, cnLocation),
		},
		{
			name:      "Skip cross-year New Year break",
			startDate: time.Date(2018, 12, 29, 0, 0, 0, 0, cnLocation),
			workdays:  0,
			wantDate:  time.Date(2019, 1, 2, 0, 0, 0, 0, cnLocation),
		},
		{
			name:      "Count cross-year makeup Saturday",
			startDate: time.Date(2018, 12, 27, 0, 0, 0, 0, cnLocation),
			workdays:  1,
			wantDate:  time.Date(2018, 12, 29, 0, 0, 0, 0, cnLocation),
		},
		{
			name:        "Negative workdays",
			startDate:   time.Date(2025, 1, 2, 0, 0, 0, 0, cnLocation),
			workdays:    -1,
			wantErr:     true,
			errContains: "must be non-negative",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDate, err := AfterWorkdays(tt.startDate, tt.workdays)
			if (err != nil) != tt.wantErr {
				t.Errorf("AfterWorkdays() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if !strings.Contains(err.Error(), tt.errContains) {
					t.Errorf("AfterWorkdays() error = %v, want error containing %v", err, tt.errContains)
				}
				return
			}
			if !gotDate.Equal(tt.wantDate) {
				t.Errorf("AfterWorkdays() = %v, want %v", gotDate, tt.wantDate)
			}
		})
	}
}

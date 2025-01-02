package holidaycn

import (
	"testing"
	"time"
)

func TestIsHoliday(t *testing.T) {
	tests := []struct {
		name           string
		date           time.Time
		wantIsHoliday  bool
		wantHolidayName string
		wantErr        bool
	}{
		{
			name:           "New Year 2025",
			date:           time.Date(2025, 1, 1, 0, 0, 0, 0, cnLocation),
			wantIsHoliday:  true,
			wantHolidayName: "元旦",
			wantErr:        false,
		},
		{
			name:           "Regular Workday",
			date:           time.Date(2025, 1, 2, 0, 0, 0, 0, cnLocation),
			wantIsHoliday:  false,
			wantHolidayName: "",
			wantErr:        false,
		},
		{
			name:           "Regular Weekend",
			date:           time.Date(2025, 1, 4, 0, 0, 0, 0, cnLocation), // Saturday
			wantIsHoliday:  true,
			wantHolidayName: "",
			wantErr:        false,
		},
		{
			name:           "Future Year",
			date:           time.Date(2030, 1, 1, 0, 0, 0, 0, cnLocation),
			wantIsHoliday:  false,
			wantHolidayName: "",
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIsHoliday, gotHolidayName, err := IsHoliday(tt.date)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsHoliday() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotIsHoliday != tt.wantIsHoliday {
				t.Errorf("IsHoliday() gotIsHoliday = %v, want %v", gotIsHoliday, tt.wantIsHoliday)
			}
			if gotHolidayName != tt.wantHolidayName {
				t.Errorf("IsHoliday() gotHolidayName = %v, want %v", gotHolidayName, tt.wantHolidayName)
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

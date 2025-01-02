package holidaycn

import (
	"fmt"
	"time"

	"github.com/washanhanzi/holiday-cn-go/pkg/holiday"
)

var cnLocation *time.Location

func init() {
	var err error
	cnLocation, err = time.LoadLocation("Asia/Shanghai")
	if err != nil {
		// Fallback to UTC+8 if timezone data is not available
		cnLocation = time.FixedZone("CST", 8*60*60)
	}
}

// IsNowHoliday checks if current time in China is a holiday
func IsNowHoliday() (bool, string, error) {
	now := time.Now().In(cnLocation)
	return IsHoliday(now)
}

// IsNowWorkday checks if current time in China is a workday
func IsNowWorkday() (bool, error) {
	now := time.Now().In(cnLocation)
	return IsWorkday(now)
}

// IsHoliday checks if a given date is a holiday
func IsHoliday(date time.Time) (bool, string, error) {
	year := date.Year()

	data := holiday.GetYearData(year)
	if data == nil {
		return false, "", fmt.Errorf("no holiday data for year %d", year)
	}

	dateStr := date.Format("2006-01-02")
	if day, exists := data[dateStr]; exists {
		return day.IsOffDay, day.Name, nil
	}

	// If not found in holiday data, use weekend logic
	return date.Weekday() == time.Saturday || date.Weekday() == time.Sunday, "", nil
}

// IsWorkday checks if a given date is a workday
func IsWorkday(date time.Time) (bool, error) {
	isHoliday, _, err := IsHoliday(date)
	if err != nil {
		return false, err
	}
	return !isHoliday, nil
}

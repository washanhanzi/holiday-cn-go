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
	return IsRestDay(now)
}

// IsNowRestDay checks if current time in China is a rest day (holiday or weekend)
func IsNowRestDay() (bool, error) {
	now := time.Now().In(cnLocation)
	isRest, _, err := IsRestDay(now)
	return isRest, err
}

// CheckHoliday returns a copy of the holiday or makeup workday record for date.
// It uses the date's own timezone and returns (nil, nil) for dates without a
// record. Years without nonempty source data return an error. Regular weekends do not
// receive a fallback record.
func CheckHoliday(date time.Time) (*holiday.Day, error) {
	return holiday.CheckHoliday(date)
}

// IsRestDay checks if a given date is either a holiday or weekend.
// Only years with nonempty source JSON day records are supported; other years return an error,
// even if their dates appear in a supported year's cross-year arrangements.
func IsRestDay(date time.Time) (bool, string, error) {
	day, err := CheckHoliday(date)
	if err != nil {
		return false, "", err
	}
	if day != nil {
		return day.IsOffDay, day.Name, nil
	}

	// If not found in holiday data, use weekend logic
	return date.Weekday() == time.Saturday || date.Weekday() == time.Sunday, "", nil
}

// IsWorkday checks if a given date is a workday
func IsWorkday(date time.Time) (bool, error) {
	isRestDay, _, err := IsRestDay(date)
	if err != nil {
		return false, err
	}
	return !isRestDay, nil
}

// AfterWorkdays returns the next workday after counting the specified number of workdays from the start date
// For example:
// - If workdays=1, it counts one workday and returns the next workday after that
// - If workdays=2, it counts two workdays and returns the next workday after that
// Rest days (holidays and weekends) are skipped in the counting
func AfterWorkdays(startDate time.Time, workdays int) (time.Time, error) {
	if workdays < 0 {
		return time.Time{}, fmt.Errorf("workdays must be non-negative, got %d", workdays)
	}

	currentDate := startDate
	remainingWorkdays := workdays

	for remainingWorkdays >= 0 {
		// Move to next day
		currentDate = currentDate.AddDate(0, 0, 1)

		// Check if it's a workday
		isWorkday, err := IsWorkday(currentDate)
		if err != nil {
			return time.Time{}, err
		}

		// Only decrement remaining workdays if it's a workday
		if isWorkday {
			remainingWorkdays--
		}
	}

	return currentDate, nil
}

# holiday-cn-go

Go package for checking Chinese holidays. Data is sourced from [holiday-cn](https://github.com/NateScarlet/holiday-cn).

## Go Version Support

Requires Go 1.16 or newer.

## Versioning

Release tags use a UTC timestamp, such as `v0.4.20260914103045`
(September 14, 2026 at 10:30:45 UTC).

## Automatic Data Updates

Holiday data is updated automatically each day.

## Code Organization

- `holidaycn.go`: Main API for checking holidays and workdays
- `pkg/holiday/*.go`: Generated code containing holiday data and initialization

## Generate Code

To generate the holiday data code:

```bash
go run cmd/generator/main.go holiday-cn pkg/holiday
```

## Usage

```go
import "github.com/washanhanzi/holiday-cn-go"

func main() {
    // Check current time in China
    isHoliday, name, err := holidaycn.IsNowHoliday()
    if err != nil {
        log.Fatal(err)
    }
    if isHoliday {
        fmt.Printf("Current time in China is a holiday: %s\n", name)
    }

    // Check if current time in China is a rest day (holiday or weekend)
    isRest, err := holidaycn.IsNowRestDay()
    if err != nil {
        log.Fatal(err)
    }
    if isRest {
        fmt.Println("Current time in China is a rest day")
    }

    // Check a specific date
    date := time.Date(2024, 2, 10, 0, 0, 0, 0, time.UTC)
    day, err := holidaycn.CheckHoliday(date)
    if err != nil {
        log.Fatal(err)
    }
    if day != nil {
        fmt.Printf("%s: %s (arrangement year: %d, off day: %t)\n", day.Date, day.Name, day.ArrangementYear, day.IsOffDay)
    }
    isHoliday, name, err = holidaycn.IsRestDay(date)
    if err != nil {
        log.Fatal(err)
    }
    if isHoliday {
        fmt.Printf("%s is a rest day: %s\n", date.Format("2006-01-02"), name)
    }

    // Check if a specific date is a workday
    isWorkday, err = holidaycn.IsWorkday(date)
    if err != nil {
        log.Fatal(err)
    }
    if isWorkday {
        fmt.Printf("%s is a workday\n", date.Format("2006-01-02"))
    }

    // Get the next workday after counting 2 workdays from a specific date
    date = time.Date(2024, 2, 10, 0, 0, 0, 0, time.UTC)
    nextWorkday, err := holidaycn.AfterWorkdays(date, 2)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Next workday after counting 2 workdays from %s is %s\n",
        date.Format("2006-01-02"), nextWorkday.Format("2006-01-02"))
}
```

## Functions

- `CheckHoliday(time.Time) (*holiday.Day, error)`: Return a holiday or makeup workday record, `(nil, nil)` if absent, or `(nil, err)` if the year is unsupported. Uses the supplied date's timezone, without weekend fallback. `Day.ArrangementYear` records the original JSON's top-level `year`, including for merged cross-year entries.
- `IsNowHoliday() (bool, string, error)`: Check if current time in China is a holiday
- `IsNowRestDay() (bool, error)`: Check if current time in China is a rest day (holiday or weekend)
- `IsRestDay(time.Time) (bool, string, error)`: Check if a given date is a rest day
- `IsWorkday(time.Time) (bool, error)`: Check if a given date is a workday
- `AfterWorkdays(time.Time, int) (time.Time, error)`: Get the next workday after counting N workdays from a given date

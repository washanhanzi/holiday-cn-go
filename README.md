# holiday-cn-go

Go package for checking Chinese holidays. Data is sourced from [holiday-cn](https://github.com/NateScarlet/holiday-cn).

## Go Version Support

Requires Go 1.16 or newer.

## Versioning

The active release series is **v0.4**. Tags use `v0.4.YYYYMMDDHHMMSS` in UTC,
for example `v0.4.20260914103045`. This date-based patch number allows multiple
releases per day. Data and maintenance updates stay within v0.4; the project
remains pre-1.0 and does not promise a stable API yet.

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

- `IsNowHoliday() (bool, string, error)`: Check if current time in China is a holiday
- `IsNowRestDay() (bool, error)`: Check if current time in China is a rest day (holiday or weekend)
- `IsRestDay(time.Time) (bool, string, error)`: Check if a given date is a rest day
- `IsWorkday(time.Time) (bool, error)`: Check if a given date is a workday
- `AfterWorkdays(time.Time, int) (time.Time, error)`: Get the next workday after counting N workdays from a given date

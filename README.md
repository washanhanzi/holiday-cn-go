# holiday-cn-go

Go package for checking Chinese holidays. Data is sourced from [holiday-cn](https://github.com/NateScarlet/holiday-cn).

## Code Organization

- `holiday.go`: Main API for checking holidays and workdays
- `pkg/holiday/*.go`: Generated code containing holiday data and initialization

## Generate Code

To generate the holiday data code:

```bash
go run cmd/generator/main.go holiday-cn pkg/holiday
```

This will:
1. Read the JSON files from the `holiday-cn` submodule
2. Generate year-specific Go files in the `pkg/holiday` directory
3. Generate `holiday.go` with data structures and initialization code

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

    // Check if current time in China is a workday
    isWorkday, err := holidaycn.IsNowWorkday()
    if err != nil {
        log.Fatal(err)
    }
    if isWorkday {
        fmt.Println("Current time in China is a workday")
    }

    // Check a specific date
    date := time.Date(2024, 2, 10, 0, 0, 0, 0, time.UTC)
    isHoliday, name, err = holidaycn.IsHoliday(date)
    if err != nil {
        log.Fatal(err)
    }
    if isHoliday {
        fmt.Printf("%s is a holiday: %s\n", date.Format("2006-01-02"), name)
    }

    // Check if a specific date is a workday
    isWorkday, err = holidaycn.IsWorkday(date)
    if err != nil {
        log.Fatal(err)
    }
    if isWorkday {
        fmt.Printf("%s is a workday\n", date.Format("2006-01-02"))
    }
}
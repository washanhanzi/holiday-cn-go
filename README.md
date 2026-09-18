# holiday-cn-go

Check Chinese holidays, makeup workdays, and weekends in Go. Requires Go 1.16 or newer.

Holiday data comes from [holiday-cn](https://github.com/NateScarlet/holiday-cn) and
is updated daily in this repository. Each year's data loads into memory when
first used. The library does not download updates at runtime.

## Quick start

```sh
go get github.com/washanhanzi/holiday-cn-go
```

```go
package main

import (
    "fmt"
    "log"
    "time"

    "github.com/washanhanzi/holiday-cn-go"
)

func main() {
    date := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
    off, name, err := holidaycn.IsRestDay(date)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(off, name) // true 元旦
}
```

`IsRestDay` checks holiday and makeup workday records first. If there is no
record, Saturday and Sunday are rest days; Monday through Friday are workdays.
Queries for unsupported years return an error, even for weekends.

Functions that accept `time.Time` use its calendar date in its own timezone.
They do not convert it to China time. The `IsNow…` functions use China time (UTC+8).

## Functions

| Function | What it does |
| --- | --- |
| `CheckHoliday(date)` | Returns a `*holiday.Day`, or `nil` when no record exists. Does not apply weekend rules. |
| `IsRestDay(date)` | Returns whether the date is a rest day, its holiday name (if any), and an error. |
| `IsWorkday(date)` | Returns whether the date is a workday and an error. |
| `IsNowRestDay()` | Checks whether today in China is a rest day. |
| `IsNowHoliday()` | Like `IsNowRestDay`, but also returns the holiday name. Includes regular weekends. |
| `AfterWorkdays(date, n)` | Skips `n` workdays after the date, then returns the following workday. Use `0` for the next workday. |
| `MutateDay("YYYY-MM-DD", callback)` | Adds or updates one cached day. |

A `holiday.Day` contains `Date`, `Name`, `IsOffDay`, and `ArrangementYear`.
`ArrangementYear` identifies the original holiday arrangement's year, which can
differ from the date's year when an arrangement spans New Year.

## Customize a day

Use `MutateDay` to add a company holiday or override an existing record. Import
`github.com/washanhanzi/holiday-cn-go/pkg/holiday` for the `holiday.Day` type.

```go
err := holidaycn.MutateDay("2025-01-02", func(day *holiday.Day) {
    day.Name = "Company holiday"
    day.IsOffDay = true
})
if err != nil {
    log.Fatal(err)
}
```

The year loads automatically before your callback runs. If the day already
exists, you get its current values. Otherwise, you get these defaults:

```go
holiday.Day{
    Date:            "2025-01-02",
    ArrangementYear: 2025,
    Name:            "",
    IsOffDay:        false,
}
```

Set the fields you want to change; the record is saved when your callback
returns. The date string has no timezone conversion. Invalid dates, unsupported
years, nil callbacks, or changes to `day.Date` return an error without saving.

Changes affect subsequent holiday checks and workday calculations in the running
process. They are not saved to disk. Only the selected year's cache changes,
including its `holiday.IsHolidayYYYY` lookups; copies of the same date in other
years' caches stay unchanged. `MutateDay` is also available in the `holiday` package.

Concurrent updates to the same year run one at a time. Readers see the old or
new data, never a partial update. Inside the callback, you may read the cache,
but must not call `MutateDay` again or continue editing from another goroutine
after returning. Changing a saved pointer later will not update the cache.
If the callback panics, its changes are discarded and the panic propagates.

## Development

- `holidaycn.go` contains the public helpers.
- `pkg/holiday/` contains generated data and cache APIs.
- `cmd/generator/` contains the generator. Edit its templates to change generated code.

Regenerate data from the local `holiday-cn` directory:

```sh
go run cmd/generator/main.go holiday-cn pkg/holiday
```

Years with no source records are not generated or supported.

Release tags use a UTC timestamp. For example, `v0.4.20260914103045` represents
September 14, 2026 at 10:30:45 UTC.

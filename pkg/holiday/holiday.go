// Code generated by generator; DO NOT EDIT.
package holiday

import (
	"sync"
)

// Day represents a single holiday or workday entry
type Day struct {
	Name     string
	Date     string
	IsOffDay bool
}

var (
	// yearDataCache stores initialized year data
	yearDataCache sync.Map

	// yearInitFuncs maps years to their initialization functions
	yearInitFuncs = map[int]func() map[string]Day{
		2007: Init2007,
		2008: Init2008,
		2009: Init2009,
		2010: Init2010,
		2011: Init2011,
		2012: Init2012,
		2013: Init2013,
		2014: Init2014,
		2015: Init2015,
		2016: Init2016,
		2017: Init2017,
		2018: Init2018,
		2019: Init2019,
		2020: Init2020,
		2021: Init2021,
		2022: Init2022,
		2023: Init2023,
		2024: Init2024,
		2025: Init2025,
		2026: Init2026,
	}
)

// GetYearData returns the holiday data for a specific year, initializing it if necessary
func GetYearData(year int) map[string]Day {
	if data, ok := yearDataCache.Load(year); ok {
		return data.(map[string]Day)
	}

	// Check if we have an init function for this year
	initFunc, ok := yearInitFuncs[year]
	if !ok {
		return nil
	}

	// Initialize the data
	data := initFunc()
	
	// Store in cache
	actualData, _ := yearDataCache.LoadOrStore(year, data)
	return actualData.(map[string]Day)
}
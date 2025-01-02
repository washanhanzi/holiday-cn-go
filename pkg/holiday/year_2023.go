// Code generated by generator; DO NOT EDIT.
package holiday

// Year2023Data contains holiday data for year 2023
var Year2023Data map[string]Day

// Init2023 initializes the holiday data for year 2023
func Init2023() map[string]Day {
	data := make(map[string]Day, 34)
	data["2022-12-31"] = Day{Name: "元旦", Date: "2022-12-31", IsOffDay: true}
	data["2023-01-01"] = Day{Name: "元旦", Date: "2023-01-01", IsOffDay: true}
	data["2023-01-02"] = Day{Name: "元旦", Date: "2023-01-02", IsOffDay: true}
	data["2023-01-21"] = Day{Name: "春节", Date: "2023-01-21", IsOffDay: true}
	data["2023-01-22"] = Day{Name: "春节", Date: "2023-01-22", IsOffDay: true}
	data["2023-01-23"] = Day{Name: "春节", Date: "2023-01-23", IsOffDay: true}
	data["2023-01-24"] = Day{Name: "春节", Date: "2023-01-24", IsOffDay: true}
	data["2023-01-25"] = Day{Name: "春节", Date: "2023-01-25", IsOffDay: true}
	data["2023-01-26"] = Day{Name: "春节", Date: "2023-01-26", IsOffDay: true}
	data["2023-01-27"] = Day{Name: "春节", Date: "2023-01-27", IsOffDay: true}
	data["2023-01-28"] = Day{Name: "春节", Date: "2023-01-28", IsOffDay: false}
	data["2023-01-29"] = Day{Name: "春节", Date: "2023-01-29", IsOffDay: false}
	data["2023-04-05"] = Day{Name: "清明节", Date: "2023-04-05", IsOffDay: true}
	data["2023-04-23"] = Day{Name: "劳动节", Date: "2023-04-23", IsOffDay: false}
	data["2023-04-29"] = Day{Name: "劳动节", Date: "2023-04-29", IsOffDay: true}
	data["2023-04-30"] = Day{Name: "劳动节", Date: "2023-04-30", IsOffDay: true}
	data["2023-05-01"] = Day{Name: "劳动节", Date: "2023-05-01", IsOffDay: true}
	data["2023-05-02"] = Day{Name: "劳动节", Date: "2023-05-02", IsOffDay: true}
	data["2023-05-03"] = Day{Name: "劳动节", Date: "2023-05-03", IsOffDay: true}
	data["2023-05-06"] = Day{Name: "劳动节", Date: "2023-05-06", IsOffDay: false}
	data["2023-06-22"] = Day{Name: "端午节", Date: "2023-06-22", IsOffDay: true}
	data["2023-06-23"] = Day{Name: "端午节", Date: "2023-06-23", IsOffDay: true}
	data["2023-06-24"] = Day{Name: "端午节", Date: "2023-06-24", IsOffDay: true}
	data["2023-06-25"] = Day{Name: "端午节", Date: "2023-06-25", IsOffDay: false}
	data["2023-09-29"] = Day{Name: "中秋节、国庆节", Date: "2023-09-29", IsOffDay: true}
	data["2023-09-30"] = Day{Name: "中秋节、国庆节", Date: "2023-09-30", IsOffDay: true}
	data["2023-10-01"] = Day{Name: "中秋节、国庆节", Date: "2023-10-01", IsOffDay: true}
	data["2023-10-02"] = Day{Name: "中秋节、国庆节", Date: "2023-10-02", IsOffDay: true}
	data["2023-10-03"] = Day{Name: "中秋节、国庆节", Date: "2023-10-03", IsOffDay: true}
	data["2023-10-04"] = Day{Name: "中秋节、国庆节", Date: "2023-10-04", IsOffDay: true}
	data["2023-10-05"] = Day{Name: "中秋节、国庆节", Date: "2023-10-05", IsOffDay: true}
	data["2023-10-06"] = Day{Name: "中秋节、国庆节", Date: "2023-10-06", IsOffDay: true}
	data["2023-10-07"] = Day{Name: "中秋节、国庆节", Date: "2023-10-07", IsOffDay: false}
	data["2023-10-08"] = Day{Name: "中秋节、国庆节", Date: "2023-10-08", IsOffDay: false}
	return data
}

// IsHoliday2023 checks if a given date in 2023 is a holiday
func IsHoliday2023(dateStr string) (bool, string, error) {
	data := GetYearData(2023)
	if day, exists := data[dateStr]; exists {
		return day.IsOffDay, day.Name, nil
	}
	return false, "", nil
}
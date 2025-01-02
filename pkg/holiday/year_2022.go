// Code generated by generator; DO NOT EDIT.
package holiday

// Year2022Data contains holiday data for year 2022
var Year2022Data map[string]Day

// Init2022 initializes the holiday data for year 2022
func Init2022() map[string]Day {
	data := make(map[string]Day, 38)
	data["2022-01-01"] = Day{Name: "元旦", Date: "2022-01-01", IsOffDay: true}
	data["2022-01-02"] = Day{Name: "元旦", Date: "2022-01-02", IsOffDay: true}
	data["2022-01-03"] = Day{Name: "元旦", Date: "2022-01-03", IsOffDay: true}
	data["2022-01-29"] = Day{Name: "春节", Date: "2022-01-29", IsOffDay: false}
	data["2022-01-30"] = Day{Name: "春节", Date: "2022-01-30", IsOffDay: false}
	data["2022-01-31"] = Day{Name: "春节", Date: "2022-01-31", IsOffDay: true}
	data["2022-02-01"] = Day{Name: "春节", Date: "2022-02-01", IsOffDay: true}
	data["2022-02-02"] = Day{Name: "春节", Date: "2022-02-02", IsOffDay: true}
	data["2022-02-03"] = Day{Name: "春节", Date: "2022-02-03", IsOffDay: true}
	data["2022-02-04"] = Day{Name: "春节", Date: "2022-02-04", IsOffDay: true}
	data["2022-02-05"] = Day{Name: "春节", Date: "2022-02-05", IsOffDay: true}
	data["2022-02-06"] = Day{Name: "春节", Date: "2022-02-06", IsOffDay: true}
	data["2022-04-02"] = Day{Name: "清明节", Date: "2022-04-02", IsOffDay: false}
	data["2022-04-03"] = Day{Name: "清明节", Date: "2022-04-03", IsOffDay: true}
	data["2022-04-04"] = Day{Name: "清明节", Date: "2022-04-04", IsOffDay: true}
	data["2022-04-05"] = Day{Name: "清明节", Date: "2022-04-05", IsOffDay: true}
	data["2022-04-24"] = Day{Name: "劳动节", Date: "2022-04-24", IsOffDay: false}
	data["2022-04-30"] = Day{Name: "劳动节", Date: "2022-04-30", IsOffDay: true}
	data["2022-05-01"] = Day{Name: "劳动节", Date: "2022-05-01", IsOffDay: true}
	data["2022-05-02"] = Day{Name: "劳动节", Date: "2022-05-02", IsOffDay: true}
	data["2022-05-03"] = Day{Name: "劳动节", Date: "2022-05-03", IsOffDay: true}
	data["2022-05-04"] = Day{Name: "劳动节", Date: "2022-05-04", IsOffDay: true}
	data["2022-05-07"] = Day{Name: "劳动节", Date: "2022-05-07", IsOffDay: false}
	data["2022-06-03"] = Day{Name: "端午节", Date: "2022-06-03", IsOffDay: true}
	data["2022-06-04"] = Day{Name: "端午节", Date: "2022-06-04", IsOffDay: true}
	data["2022-06-05"] = Day{Name: "端午节", Date: "2022-06-05", IsOffDay: true}
	data["2022-09-10"] = Day{Name: "中秋节", Date: "2022-09-10", IsOffDay: true}
	data["2022-09-11"] = Day{Name: "中秋节", Date: "2022-09-11", IsOffDay: true}
	data["2022-09-12"] = Day{Name: "中秋节", Date: "2022-09-12", IsOffDay: true}
	data["2022-10-01"] = Day{Name: "国庆节", Date: "2022-10-01", IsOffDay: true}
	data["2022-10-02"] = Day{Name: "国庆节", Date: "2022-10-02", IsOffDay: true}
	data["2022-10-03"] = Day{Name: "国庆节", Date: "2022-10-03", IsOffDay: true}
	data["2022-10-04"] = Day{Name: "国庆节", Date: "2022-10-04", IsOffDay: true}
	data["2022-10-05"] = Day{Name: "国庆节", Date: "2022-10-05", IsOffDay: true}
	data["2022-10-06"] = Day{Name: "国庆节", Date: "2022-10-06", IsOffDay: true}
	data["2022-10-07"] = Day{Name: "国庆节", Date: "2022-10-07", IsOffDay: true}
	data["2022-10-08"] = Day{Name: "国庆节", Date: "2022-10-08", IsOffDay: false}
	data["2022-10-09"] = Day{Name: "国庆节", Date: "2022-10-09", IsOffDay: false}
	return data
}

// IsHoliday2022 checks if a given date in 2022 is a holiday
func IsHoliday2022(dateStr string) (bool, string, error) {
	data := GetYearData(2022)
	if day, exists := data[dateStr]; exists {
		return day.IsOffDay, day.Name, nil
	}
	return false, "", nil
}
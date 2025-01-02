// Code generated by generator; DO NOT EDIT.
package holiday

// Year2019Data contains holiday data for year 2019
var Year2019Data map[string]Day

// Init2019 initializes the holiday data for year 2019
func Init2019() map[string]Day {
	data := make(map[string]Day, 31)
	data["2018-12-29"] = Day{Name: "元旦", Date: "2018-12-29", IsOffDay: false}
	data["2018-12-30"] = Day{Name: "元旦", Date: "2018-12-30", IsOffDay: true}
	data["2018-12-31"] = Day{Name: "元旦", Date: "2018-12-31", IsOffDay: true}
	data["2019-01-01"] = Day{Name: "元旦", Date: "2019-01-01", IsOffDay: true}
	data["2019-02-02"] = Day{Name: "春节", Date: "2019-02-02", IsOffDay: false}
	data["2019-02-03"] = Day{Name: "春节", Date: "2019-02-03", IsOffDay: false}
	data["2019-02-04"] = Day{Name: "春节", Date: "2019-02-04", IsOffDay: true}
	data["2019-02-05"] = Day{Name: "春节", Date: "2019-02-05", IsOffDay: true}
	data["2019-02-06"] = Day{Name: "春节", Date: "2019-02-06", IsOffDay: true}
	data["2019-02-07"] = Day{Name: "春节", Date: "2019-02-07", IsOffDay: true}
	data["2019-02-08"] = Day{Name: "春节", Date: "2019-02-08", IsOffDay: true}
	data["2019-02-09"] = Day{Name: "春节", Date: "2019-02-09", IsOffDay: true}
	data["2019-02-10"] = Day{Name: "春节", Date: "2019-02-10", IsOffDay: true}
	data["2019-04-05"] = Day{Name: "清明节", Date: "2019-04-05", IsOffDay: true}
	data["2019-04-28"] = Day{Name: "劳动节", Date: "2019-04-28", IsOffDay: false}
	data["2019-05-01"] = Day{Name: "劳动节", Date: "2019-05-01", IsOffDay: true}
	data["2019-05-02"] = Day{Name: "劳动节", Date: "2019-05-02", IsOffDay: true}
	data["2019-05-03"] = Day{Name: "劳动节", Date: "2019-05-03", IsOffDay: true}
	data["2019-05-04"] = Day{Name: "劳动节", Date: "2019-05-04", IsOffDay: true}
	data["2019-05-05"] = Day{Name: "劳动节", Date: "2019-05-05", IsOffDay: false}
	data["2019-06-07"] = Day{Name: "端午节", Date: "2019-06-07", IsOffDay: true}
	data["2019-09-13"] = Day{Name: "中秋节", Date: "2019-09-13", IsOffDay: true}
	data["2019-09-29"] = Day{Name: "国庆节", Date: "2019-09-29", IsOffDay: false}
	data["2019-10-01"] = Day{Name: "国庆节", Date: "2019-10-01", IsOffDay: true}
	data["2019-10-02"] = Day{Name: "国庆节", Date: "2019-10-02", IsOffDay: true}
	data["2019-10-03"] = Day{Name: "国庆节", Date: "2019-10-03", IsOffDay: true}
	data["2019-10-04"] = Day{Name: "国庆节", Date: "2019-10-04", IsOffDay: true}
	data["2019-10-05"] = Day{Name: "国庆节", Date: "2019-10-05", IsOffDay: true}
	data["2019-10-06"] = Day{Name: "国庆节", Date: "2019-10-06", IsOffDay: true}
	data["2019-10-07"] = Day{Name: "国庆节", Date: "2019-10-07", IsOffDay: true}
	data["2019-10-12"] = Day{Name: "国庆节", Date: "2019-10-12", IsOffDay: false}
	return data
}

// IsHoliday2019 checks if a given date in 2019 is a holiday
func IsHoliday2019(dateStr string) (bool, string, error) {
	data := GetYearData(2019)
	if day, exists := data[dateStr]; exists {
		return day.IsOffDay, day.Name, nil
	}
	return false, "", nil
}
// Code generated by generator; DO NOT EDIT.
package holiday

// Year2011Data contains holiday data for year 2011
var Year2011Data map[string]Day

// Init2011 initializes the holiday data for year 2011
func Init2011() map[string]Day {
	data := make(map[string]Day, 34)
	data["2011-01-01"] = Day{Name: "元旦", Date: "2011-01-01", IsOffDay: true}
	data["2011-01-02"] = Day{Name: "元旦", Date: "2011-01-02", IsOffDay: true}
	data["2011-01-03"] = Day{Name: "元旦", Date: "2011-01-03", IsOffDay: true}
	data["2011-01-30"] = Day{Name: "春节", Date: "2011-01-30", IsOffDay: false}
	data["2011-02-02"] = Day{Name: "春节", Date: "2011-02-02", IsOffDay: true}
	data["2011-02-03"] = Day{Name: "春节", Date: "2011-02-03", IsOffDay: true}
	data["2011-02-04"] = Day{Name: "春节", Date: "2011-02-04", IsOffDay: true}
	data["2011-02-05"] = Day{Name: "春节", Date: "2011-02-05", IsOffDay: true}
	data["2011-02-06"] = Day{Name: "春节", Date: "2011-02-06", IsOffDay: true}
	data["2011-02-07"] = Day{Name: "春节", Date: "2011-02-07", IsOffDay: true}
	data["2011-02-08"] = Day{Name: "春节", Date: "2011-02-08", IsOffDay: true}
	data["2011-02-12"] = Day{Name: "春节", Date: "2011-02-12", IsOffDay: false}
	data["2011-04-02"] = Day{Name: "清明节", Date: "2011-04-02", IsOffDay: false}
	data["2011-04-03"] = Day{Name: "清明节", Date: "2011-04-03", IsOffDay: true}
	data["2011-04-04"] = Day{Name: "清明节", Date: "2011-04-04", IsOffDay: true}
	data["2011-04-05"] = Day{Name: "清明节", Date: "2011-04-05", IsOffDay: true}
	data["2011-04-30"] = Day{Name: "劳动节", Date: "2011-04-30", IsOffDay: true}
	data["2011-05-01"] = Day{Name: "劳动节", Date: "2011-05-01", IsOffDay: true}
	data["2011-05-02"] = Day{Name: "劳动节", Date: "2011-05-02", IsOffDay: true}
	data["2011-06-04"] = Day{Name: "端午节", Date: "2011-06-04", IsOffDay: true}
	data["2011-06-05"] = Day{Name: "端午节", Date: "2011-06-05", IsOffDay: true}
	data["2011-06-06"] = Day{Name: "端午节", Date: "2011-06-06", IsOffDay: true}
	data["2011-09-10"] = Day{Name: "中秋节", Date: "2011-09-10", IsOffDay: true}
	data["2011-09-11"] = Day{Name: "中秋节", Date: "2011-09-11", IsOffDay: true}
	data["2011-09-12"] = Day{Name: "中秋节", Date: "2011-09-12", IsOffDay: true}
	data["2011-10-01"] = Day{Name: "国庆节", Date: "2011-10-01", IsOffDay: true}
	data["2011-10-02"] = Day{Name: "国庆节", Date: "2011-10-02", IsOffDay: true}
	data["2011-10-03"] = Day{Name: "国庆节", Date: "2011-10-03", IsOffDay: true}
	data["2011-10-04"] = Day{Name: "国庆节", Date: "2011-10-04", IsOffDay: true}
	data["2011-10-05"] = Day{Name: "国庆节", Date: "2011-10-05", IsOffDay: true}
	data["2011-10-06"] = Day{Name: "国庆节", Date: "2011-10-06", IsOffDay: true}
	data["2011-10-07"] = Day{Name: "国庆节", Date: "2011-10-07", IsOffDay: true}
	data["2011-10-08"] = Day{Name: "国庆节", Date: "2011-10-08", IsOffDay: false}
	data["2011-10-09"] = Day{Name: "国庆节", Date: "2011-10-09", IsOffDay: false}
	return data
}

// IsHoliday2011 checks if a given date in 2011 is a holiday
func IsHoliday2011(dateStr string) (bool, string, error) {
	data := GetYearData(2011)
	if day, exists := data[dateStr]; exists {
		return day.IsOffDay, day.Name, nil
	}
	return false, "", nil
}
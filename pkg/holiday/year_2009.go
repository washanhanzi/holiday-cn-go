// Code generated by generator; DO NOT EDIT.
package holiday

// Year2009Data contains holiday data for year 2009
var Year2009Data map[string]Day

// Init2009 initializes the holiday data for year 2009
func Init2009() map[string]Day {
	data := make(map[string]Day, 33)
	data["2009-01-01"] = Day{Name: "元旦", Date: "2009-01-01", IsOffDay: true}
	data["2009-01-02"] = Day{Name: "元旦", Date: "2009-01-02", IsOffDay: true}
	data["2009-01-03"] = Day{Name: "元旦", Date: "2009-01-03", IsOffDay: true}
	data["2009-01-04"] = Day{Name: "元旦", Date: "2009-01-04", IsOffDay: false}
	data["2009-01-24"] = Day{Name: "春节", Date: "2009-01-24", IsOffDay: false}
	data["2009-01-25"] = Day{Name: "春节", Date: "2009-01-25", IsOffDay: true}
	data["2009-01-26"] = Day{Name: "春节", Date: "2009-01-26", IsOffDay: true}
	data["2009-01-27"] = Day{Name: "春节", Date: "2009-01-27", IsOffDay: true}
	data["2009-01-28"] = Day{Name: "春节", Date: "2009-01-28", IsOffDay: true}
	data["2009-01-29"] = Day{Name: "春节", Date: "2009-01-29", IsOffDay: true}
	data["2009-01-30"] = Day{Name: "春节", Date: "2009-01-30", IsOffDay: true}
	data["2009-01-31"] = Day{Name: "春节", Date: "2009-01-31", IsOffDay: true}
	data["2009-02-01"] = Day{Name: "春节", Date: "2009-02-01", IsOffDay: false}
	data["2009-04-04"] = Day{Name: "清明节", Date: "2009-04-04", IsOffDay: true}
	data["2009-04-05"] = Day{Name: "清明节", Date: "2009-04-05", IsOffDay: true}
	data["2009-04-06"] = Day{Name: "清明节", Date: "2009-04-06", IsOffDay: true}
	data["2009-05-01"] = Day{Name: "劳动节", Date: "2009-05-01", IsOffDay: true}
	data["2009-05-02"] = Day{Name: "劳动节", Date: "2009-05-02", IsOffDay: true}
	data["2009-05-03"] = Day{Name: "劳动节", Date: "2009-05-03", IsOffDay: true}
	data["2009-05-28"] = Day{Name: "端午节", Date: "2009-05-28", IsOffDay: true}
	data["2009-05-29"] = Day{Name: "端午节", Date: "2009-05-29", IsOffDay: true}
	data["2009-05-30"] = Day{Name: "端午节", Date: "2009-05-30", IsOffDay: true}
	data["2009-05-31"] = Day{Name: "端午节", Date: "2009-05-31", IsOffDay: false}
	data["2009-09-27"] = Day{Name: "国庆节、中秋节", Date: "2009-09-27", IsOffDay: false}
	data["2009-10-01"] = Day{Name: "国庆节、中秋节", Date: "2009-10-01", IsOffDay: true}
	data["2009-10-02"] = Day{Name: "国庆节、中秋节", Date: "2009-10-02", IsOffDay: true}
	data["2009-10-03"] = Day{Name: "国庆节、中秋节", Date: "2009-10-03", IsOffDay: true}
	data["2009-10-04"] = Day{Name: "国庆节、中秋节", Date: "2009-10-04", IsOffDay: true}
	data["2009-10-05"] = Day{Name: "国庆节、中秋节", Date: "2009-10-05", IsOffDay: true}
	data["2009-10-06"] = Day{Name: "国庆节、中秋节", Date: "2009-10-06", IsOffDay: true}
	data["2009-10-07"] = Day{Name: "国庆节、中秋节", Date: "2009-10-07", IsOffDay: true}
	data["2009-10-08"] = Day{Name: "国庆节、中秋节", Date: "2009-10-08", IsOffDay: true}
	data["2009-10-10"] = Day{Name: "国庆节、中秋节", Date: "2009-10-10", IsOffDay: false}
	return data
}

// IsHoliday2009 checks if a given date in 2009 is a holiday
func IsHoliday2009(dateStr string) (bool, string, error) {
	data := GetYearData(2009)
	if day, exists := data[dateStr]; exists {
		return day.IsOffDay, day.Name, nil
	}
	return false, "", nil
}
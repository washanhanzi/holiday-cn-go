// Code generated by generator; DO NOT EDIT.
package holiday

// Year2015Data contains holiday data for year 2015
var Year2015Data map[string]Day

// Init2015 initializes the holiday data for year 2015
func Init2015() map[string]Day {
	data := make(map[string]Day, 31)
	data["2015-01-01"] = Day{Name: "元旦", Date: "2015-01-01", IsOffDay: true}
	data["2015-01-02"] = Day{Name: "元旦", Date: "2015-01-02", IsOffDay: true}
	data["2015-01-03"] = Day{Name: "元旦", Date: "2015-01-03", IsOffDay: true}
	data["2015-01-04"] = Day{Name: "元旦", Date: "2015-01-04", IsOffDay: false}
	data["2015-02-15"] = Day{Name: "春节", Date: "2015-02-15", IsOffDay: false}
	data["2015-02-18"] = Day{Name: "春节", Date: "2015-02-18", IsOffDay: true}
	data["2015-02-19"] = Day{Name: "春节", Date: "2015-02-19", IsOffDay: true}
	data["2015-02-20"] = Day{Name: "春节", Date: "2015-02-20", IsOffDay: true}
	data["2015-02-21"] = Day{Name: "春节", Date: "2015-02-21", IsOffDay: true}
	data["2015-02-22"] = Day{Name: "春节", Date: "2015-02-22", IsOffDay: true}
	data["2015-02-23"] = Day{Name: "春节", Date: "2015-02-23", IsOffDay: true}
	data["2015-02-24"] = Day{Name: "春节", Date: "2015-02-24", IsOffDay: true}
	data["2015-02-28"] = Day{Name: "春节", Date: "2015-02-28", IsOffDay: false}
	data["2015-04-05"] = Day{Name: "清明节", Date: "2015-04-05", IsOffDay: true}
	data["2015-04-06"] = Day{Name: "清明节", Date: "2015-04-06", IsOffDay: true}
	data["2015-05-01"] = Day{Name: "劳动节", Date: "2015-05-01", IsOffDay: true}
	data["2015-06-20"] = Day{Name: "端午节", Date: "2015-06-20", IsOffDay: true}
	data["2015-06-22"] = Day{Name: "端午节", Date: "2015-06-22", IsOffDay: true}
	data["2015-09-03"] = Day{Name: "抗日战争暨世界反法西斯战争胜利70周年纪念日", Date: "2015-09-03", IsOffDay: true}
	data["2015-09-04"] = Day{Name: "抗日战争暨世界反法西斯战争胜利70周年纪念日", Date: "2015-09-04", IsOffDay: true}
	data["2015-09-05"] = Day{Name: "抗日战争暨世界反法西斯战争胜利70周年纪念日", Date: "2015-09-05", IsOffDay: true}
	data["2015-09-06"] = Day{Name: "抗日战争暨世界反法西斯战争胜利70周年纪念日", Date: "2015-09-06", IsOffDay: false}
	data["2015-09-27"] = Day{Name: "中秋节", Date: "2015-09-27", IsOffDay: true}
	data["2015-10-01"] = Day{Name: "国庆节", Date: "2015-10-01", IsOffDay: true}
	data["2015-10-02"] = Day{Name: "国庆节", Date: "2015-10-02", IsOffDay: true}
	data["2015-10-03"] = Day{Name: "国庆节", Date: "2015-10-03", IsOffDay: true}
	data["2015-10-04"] = Day{Name: "国庆节", Date: "2015-10-04", IsOffDay: true}
	data["2015-10-05"] = Day{Name: "国庆节", Date: "2015-10-05", IsOffDay: true}
	data["2015-10-06"] = Day{Name: "国庆节", Date: "2015-10-06", IsOffDay: true}
	data["2015-10-07"] = Day{Name: "国庆节", Date: "2015-10-07", IsOffDay: true}
	data["2015-10-10"] = Day{Name: "国庆节", Date: "2015-10-10", IsOffDay: false}
	return data
}

// IsHoliday2015 checks if a given date in 2015 is a holiday
func IsHoliday2015(dateStr string) (bool, string, error) {
	data := GetYearData(2015)
	if day, exists := data[dateStr]; exists {
		return day.IsOffDay, day.Name, nil
	}
	return false, "", nil
}
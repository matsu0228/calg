package calg

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var today time.Time

// Calender is interface for each contry calender
// type Calender interface {
// 	Sho() string
// }

// JpCalender is calender of japan
type JpCalender struct {
	startWithMonday bool
	verbos          bool
	weekdayIndex    []int
}

// NewCalender is constractor
func NewCalender(now time.Time, swMonday, v bool) *JpCalender { //*Calender {

	// switch calType {
	// case "jp":
	// }
	today = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	wdIndex := []int{0, 1, 2, 3, 4, 5, 6} //defalt
	if swMonday {
		wdIndex = []int{1, 2, 3, 4, 5, 6, 0}
	}
	if v {
		log.Print("[INFO]", wdIndex)
	}
	return &JpCalender{startWithMonday: swMonday, verbos: v, weekdayIndex: wdIndex}
}

// ShowMonthly :display monthly calender
func (c JpCalender) ShowMonthly() {

	// display label
	fmt.Println(c.ShowCalenderLabel())
	fmt.Println(c.ShowWeekLabel())

	date := c.BeginDay()

	for {
		fmt.Println(c.ShowWeek(date))
		date = c.NextWeekDay(date)
		if c.LastDay().Before(date) {
			break
		}
	}
}

func (c JpCalender) isLargerDay(i, startWeekDayInt int) bool {
	if c.startWithMonday {
		if i == 0 {
			return true
		}
		if c.verbos {
			log.Printf("i=%v, startWeekDayInt=%v res=%v", i, startWeekDayInt, i >= startWeekDayInt)
		}
		return i >= (startWeekDayInt)
	}
	return i >= startWeekDayInt
}

func (c JpCalender) weekdayInt(d time.Time) int {
	return c.weekdayIndex[d.Weekday()]
}

// ShowCalenderLabel return calender label(year, month)
func (c JpCalender) ShowCalenderLabel() string {
	if c.verbos {
		log.Printf("BeginDay:%v, LastDay:%v", c.BeginDay(), c.LastDay())
	}
	return fmt.Sprintf("     %v年 %v月",
		today.Year(),
		int(today.Month()),
	)
}

// ShowWeekLabel return weekday label
func (c JpCalender) ShowWeekLabel() string {
	weekdays := []string{}
	for _, i := range c.weekdayIndex {
		weekdays = append(weekdays, c.convertWeekday(i))
	}
	return strings.Join(weekdays, " ")
}

// BeginDay return time of begin day of this month
func (c JpCalender) BeginDay() time.Time {
	return time.Date(today.Year(), today.Month(), 1, 0, 0, 0, 0, time.Local)
}

// LastDay return time of last day of this month
func (c JpCalender) LastDay() time.Time {
	return time.Date(today.Year(), today.Month()+1, 1, 0, 0, 0, 0, time.Local).AddDate(0, 0, -1)
}

// NextWeekDay return time of next week's start day.
func (c JpCalender) NextWeekDay(d time.Time) time.Time {
	weekDayInt := int(d.Weekday())
	return d.AddDate(0, 0, 7-weekDayInt)
}

// fillDate is fill space for days less than 10th.
func (c JpCalender) fillDate(d time.Time) string {

	dStr := fmt.Sprint(d.Day())
	if c.startWithMonday {
		dStr = fmt.Sprint(d.AddDate(0, 0, 1).Day())
	}

	// a digit or more
	isSmallDate := false
	if len(dStr) == 1 {
		isSmallDate = true
	}

	// heiright for today
	if fmt.Sprint(today.Day()) == dStr {
		dStr = colorable(dStr)
	}
	if isSmallDate {
		return " " + dStr
	}
	return dStr
}

// fillWeekDay is fill days for `begin of month` or `last of month`
func (c JpCalender) fillWeekDay(startDay time.Time) string {
	startWeekDayInt := int(startDay.Weekday())
	days := []string{}

	// begin of month
	if c.isBeginWeek(startDay) {
		// for i := 0; i < 7; i++ {
		for index, i := range c.weekdayIndex {
			if c.verbos {
				fmt.Fprintf(os.Stderr, "[DEBUG] loop=%v %v %v", i, index-startWeekDayInt, c.isLargerDay(i, startWeekDayInt))
			}
			if c.isLargerDay(i, startWeekDayInt) {
				days = append(days, c.fillDate(startDay.AddDate(0, 0, index-startWeekDayInt)))
			} else {
				days = append(days, "  ")
			}
		}
		return fmt.Sprint(strings.Join(days, " "))
	}

	// last of month
	for index, i := range c.weekdayIndex {
		lastWeekDayInt := int(c.LastDay().Weekday())
		if c.verbos {
			fmt.Fprintf(os.Stderr, "[DEGUB] loop=%v %v %v", i, index-startWeekDayInt, c.isLargerDay(i, lastWeekDayInt))
		}
		// for i := 0; i < 7; i++ {
		if !c.isLargerDay(i, lastWeekDayInt+1) {
			// if i <= startWeekDayInt {
			days = append(days, c.fillDate(startDay.AddDate(0, 0, index)))
		} else {
			days = append(days, "  ")
		}
	}
	return fmt.Sprint(strings.Join(days, " "))
}

func (c JpCalender) isBeginWeek(d time.Time) bool {
	subFromBegin := d.Sub(c.BeginDay())
	return subFromBegin.String() == "0s"
}

func (c JpCalender) isLastWeek(d time.Time) bool {
	subToLast := fmt.Sprint(c.LastDay().Sub(d.AddDate(0, 0, 6)))
	return (subToLast == "0s" || subToLast[:1] == "-")
}

// ShowWeek return string of days
func (c JpCalender) ShowWeek(startDay time.Time) string {
	days := []string{}

	if c.verbos {
		fmt.Fprintf(os.Stderr, "[DEGUB] isBegin:%v, isLast:%v", c.isBeginWeek(startDay), c.isLastWeek(startDay))
	}

	if c.isBeginWeek(startDay) || c.isLastWeek(startDay) {
		return c.fillWeekDay(startDay)
	}
	for i := 0; i < 7; i++ {
		days = append(days, c.fillDate(startDay.AddDate(0, 0, i)))
	}
	return fmt.Sprint(strings.Join(days, " "))
}

// Weekday return string of weekday
func (c JpCalender) Weekday(d time.Time) string {
	return c.convertWeekday(int(d.Weekday()))
}
func (c JpCalender) convertWeekday(i int) string {
	wdays := []string{"日", "月", "火", "水", "木", "金", "土"}
	return wdays[i]
}

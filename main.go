package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"time"
)

var verbos bool
var today time.Time

// // Calender :
// type Calender interface {
// 	Weekday() string
// }
//
type jpCalender struct {
}

//
// // NewCalender is constractor
// func NewCalender(calType string) *Calender {
// 	switch calType {
// 	case "jp":
// 		return &jpCalender{today: time.Now(), wday: time.Weekday()}
// 	}
// 	return &jpCalender{today: time.Now(), wday: time.Weekday()}
// }

func (c jpCalender) ShowCalenderLabel() string {
	return fmt.Sprintf("     %v年 %v月",
		today.Year(),
		int(today.Month()),
	)
}

func (c jpCalender) ShowWeekLabel() string {
	return fmt.Sprintf("%v %v %v %v %v %v %v",
		c.convertWeekday(0),
		c.convertWeekday(1),
		c.convertWeekday(2),
		c.convertWeekday(3),
		c.convertWeekday(4),
		c.convertWeekday(5),
		c.convertWeekday(6),
	)
}

func (c jpCalender) BeginDay() time.Time {
	return time.Date(today.Year(), today.Month(), 1, 0, 0, 0, 0, time.Local)
}
func (c jpCalender) LastDay() time.Time {
	return time.Date(today.Year(), today.Month()+1, 1, 0, 0, 0, 0, time.Local).AddDate(0, 0, -1)
}

func (c jpCalender) NextWeekDay(d time.Time) time.Time {
	weekDayInt := int(d.Weekday())
	return d.AddDate(0, 0, 7-weekDayInt)
}

// fillDate is fill space for days less than 10th.
func (c jpCalender) fillDate(d time.Time) string {
	dStr := fmt.Sprint(d.Day())
	// heiright for today
	if fmt.Sprint(today.Day()) == dStr {
		dStr = colorable(dStr)
	}
	if len(dStr) == 1 {
		return " " + dStr
	}
	return dStr
}

func colorable(s string) string {
	return fmt.Sprintf("\x1b[31m%s\x1b[0m", s)
}

// fillWeekDay is fill days for `begin of month` or `last of month`
func (c jpCalender) fillWeekDay(startDay time.Time) string {
	startWeekDayInt := int(startDay.Weekday())
	days := []string{}

	// begin of month
	if c.isBeginWeek(startDay) {
		for i := 0; i < 7; i++ {
			if i >= startWeekDayInt {
				days = append(days, c.fillDate(startDay.AddDate(0, 0, i-startWeekDayInt)))
			} else {
				days = append(days, "  ")
			}
		}
		return fmt.Sprint(strings.Join(days, " "))
	}

	// last of month
	for i := 0; i < 7; i++ {
		if i <= startWeekDayInt {
			days = append(days, c.fillDate(startDay.AddDate(0, 0, i)))
		} else {
			days = append(days, "  ")
		}
	}
	return fmt.Sprint(strings.Join(days, " "))
}

func (c jpCalender) isBeginWeek(d time.Time) bool {
	subFromBegin := d.Sub(c.BeginDay())
	return subFromBegin.String() == "0s"
}

func (c jpCalender) isLastWeek(d time.Time) bool {
	subToLast := fmt.Sprint(c.LastDay().Sub(d.AddDate(0, 0, 6)))
	return (subToLast == "0s" || subToLast[:1] == "-")
}

func (c jpCalender) ShowWeek(startDay time.Time) string {
	days := []string{}

	if verbos {
		log.Printf("isBegin:%v, isLast:%v", c.isBeginWeek(startDay), c.isLastWeek(startDay))
	}

	if c.isBeginWeek(startDay) || c.isLastWeek(startDay) {
		return c.fillWeekDay(startDay)
	}
	for i := 0; i < 7; i++ {
		days = append(days, c.fillDate(startDay.AddDate(0, 0, i)))
	}
	return fmt.Sprint(strings.Join(days, " "))
}

func (c jpCalender) Weekday(d time.Time) string {
	return c.convertWeekday(int(d.Weekday()))
}
func (c jpCalender) convertWeekday(i int) string {
	wdays := []string{"日", "月", "火", "水", "木", "金", "土"}
	return wdays[i]
}

func init() {
	flag.BoolVar(&verbos, "v", false, "デバッグ用文言の表示")
	// flag.StringVar(&convCommand, "f", "jpg:png", "変換元の画像形式:変換後の画像形式 (png/jpg/gifのみサポート)")
	flag.Parse()
}

func main() {
	now := time.Now()
	today = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	c := jpCalender{}

	if verbos {
		log.Printf("today: %v, last: %v", today, c.LastDay())
	}

	fmt.Println(c.ShowCalenderLabel())
	fmt.Println(c.ShowWeekLabel())

	date := c.BeginDay()
	fmt.Println(c.ShowWeek(date))

	date = c.NextWeekDay(date)
	fmt.Println(c.ShowWeek(date))

	date = c.NextWeekDay(date)
	fmt.Println(c.ShowWeek(date))

	date = c.NextWeekDay(date)
	fmt.Println(c.ShowWeek(date))

	date = c.NextWeekDay(date)
	fmt.Println(c.ShowWeek(date))

	date = c.NextWeekDay(date)
	fmt.Println(c.ShowWeek(date))

	// fmt.Println(c.ShowWeek(today))

}

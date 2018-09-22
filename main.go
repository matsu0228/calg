package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/matsu0228/calg/calg"
)

var verbos bool
var startWithMonday bool

func init() {
	flag.BoolVar(&verbos, "v", false, "show debug strings")
	flag.BoolVar(&startWithMonday, "m", false, "weekday with starting monday")
	flag.Parse()
}

func main() {
	now := time.Now()
	// now := time.Now().AddDate(1, 1, 0) //debug
	c := calg.NewCalender(now, startWithMonday, verbos)

	// if v {
	// 	log.Printf("today: %v, last: %v", today, c.LastDay())
	// }
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

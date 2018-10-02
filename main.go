package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/matsu0228/calg/calg"
)

var verbos bool
var startWithSunday bool

func init() {
	flag.BoolVar(&verbos, "v", false, "show debug strings")
	flag.BoolVar(&startWithSunday, "s", false, "weekday with starting sunday(default: start with monday)")
	flag.Parse()
}

func main() {
	now := time.Now()
	// now := time.Now().AddDate(1, 1, 0) //debug

	// default: start with monday
	c := calg.NewCalender(now, !startWithSunday, verbos)

	if verbos {
		fmt.Fprintf(os.Stderr, "now: %v, last: %v", now, c.LastDay())
	}

	// display monthly calender
	c.ShowMonthly()

}

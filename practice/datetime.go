package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	var dateString string

	if len(os.Args) >= 2 {
		dateString = os.Args[1]
	} else {
		d := time.Now()
		tenHoursAgo := d.Add(-time.Hour * 10)
		dateString = tenHoursAgo.Format(time.RFC850)
		fmt.Println("RFC850 format:", dateString)
		dateString = d.Format(time.RFC1123Z)
		fmt.Println("RFC1123Z format:", dateString)
		dateString = d.Format(time.RFC3339)
		fmt.Println("RFC3339 format:", dateString)
		loc, err := time.LoadLocation("Europe/London")
		if err == nil {
			fmt.Println(loc)
			dateString = tenHoursAgo.In(loc).Format(time.RFC3339)
			fmt.Println("Ten hours ago in London:", dateString)
		}

		// RFC850 format: Wednesday, 25-Aug-21 01:36:38 IDT
		// RFC1123Z format: Wed, 25 Aug 2021 11:36:38 +0300
		// RFC3339 format: 2021-08-25T11:36:38+03:00
		// Europe/London
		// Ten hours ago in London: 2021-08-24T23:36:38+01:00
		return
	}

	// Is this a date only?
	d, err := time.Parse("02 January 2006", dateString)
	if err == nil {
		fmt.Println("Full:", d)
		fmt.Println("Time:", d.Day(), d.Month(), d.Year())
	}

	// Is this a date + time?
	d, err = time.Parse("02 January 2006 15:04", dateString)
	if err == nil {
		fmt.Println("Full:", d)
		fmt.Println("Date:", d.Day(), d.Month(), d.Year())
		fmt.Println("Time:", d.Hour(), d.Minute())
	}

	// Is this a date + time with month represented as a number?
	d, err = time.Parse("02-01-2006 15:04", dateString)
	if err == nil {
		fmt.Println("Full:", d)
		fmt.Println("Date:", d.Day(), d.Month(), d.Year())
		fmt.Println("Time:", d.Hour(), d.Minute())
	}

	// Is it time only?
	d, err = time.Parse("15:04", dateString)
	if err == nil {
		fmt.Println("Full:", d)
		fmt.Println("Time:", d.Hour(), d.Minute())
	}

	t := time.Now().Unix()
	fmt.Println("Epoch time:", t)
	// Convert Epoch time to time.Time
	d = time.Unix(t, 0)
	fmt.Println("Date:", d.Day(), d.Month(), d.Year())
	fmt.Printf("Time: %d:%d\n", d.Hour(), d.Minute())

	duration := time.Since(start)
	fmt.Println("Execution time:", duration)
}

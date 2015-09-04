package main

import (
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"os"
	"regexp"
	"strconv"
	"time"
)

const (
	DefaultNumOfResultToShow = 3
)

type Time struct {
	hour, minute int
}

func (time *Time) toString() string {
	if time.minute < 10 {
		return fmt.Sprintf("%d:0%d ", time.hour, time.minute)
	} else {
		return fmt.Sprintf("%d:%d ", time.hour, time.minute)
	}
}

func departureTime(departure string) (int, int) {
	ret, _ := regexp.MatchString("^[0-9]{1,2}:[0-9]{1,2}$", departure)
	var hour int
	var minute int
	if ret {
		re := regexp.MustCompile("^([0-9]{1,2}):([0-9]{1,2})$")
		bs := []byte(departure)
		group := re.FindSubmatch(bs)
		h, _ := strconv.Atoi(string(group[1]))
		m, _ := strconv.Atoi(string(group[2]))
		if h >= 0 && h < 24 && m >= 0 && m < 60 {
			hour = h
			minute = m
		} else {
			now := time.Now()
			hour = now.Hour()
			minute = now.Minute()
		}
	} else {
		now := time.Now()
		hour = now.Hour()
		minute = now.Minute()
	}
	return hour, minute
}

func getTimeTableUrl() string {
	weekday := time.Now().Weekday().String()
	if weekday == "Saturday" || weekday == "Sunday" {
		return "http://www.jreast-timetable.jp/1509/timetable/tt0413/0413021.html"
	} else {
		return "http://www.jreast-timetable.jp/1509/timetable/tt0413/0413020.html"
	}
}

func createTimetable() []Time {
	var timetable = make([]Time, 0)
	url := getTimeTableUrl()
	doc, _ := goquery.NewDocument(url)
	doc.Find(".timetable .result_03").Each(func(_ int, s *goquery.Selection) {
		s.Find("tbody tr").Each(func(_ int, s *goquery.Selection) {
			key, err := strconv.Atoi(s.Find("td:nth-child(1)").Text())
			if err == nil {
				s.Find("td:nth-child(2) div.timetable_time[data-dest='府']").Each(func(_ int, s *goquery.Selection) {
					value, _ := strconv.Atoi(s.Find(".minute").Text())
					timetable = append(timetable, Time{key, value})
				})
			}
		})
	})
	return timetable
}

func printTimes(times []string) {
	for _, v := range times {
		fmt.Println(v)
	}
}

func main() {
	var (
		departure   string
		numOfResult int
		isLast      bool
	)
	flag.StringVar(&departure, "t", "", "specify departure time.")
	flag.BoolVar(&isLast, "l", false, "show last bus of the day.")
	flag.IntVar(&numOfResult, "n", DefaultNumOfResultToShow, "specify amount of result.")
	flag.Parse()

	if numOfResult < 0 {
		fmt.Fprintf(os.Stderr, "parameter for -n must be greater than 0.\n")
		os.Exit(2)
	}

	hour, minute := departureTime(departure)
	timetable := createTimetable()

	if isLast {
		result := []string{timetable[len(timetable)-1].toString()}
		printTimes(result)
		return
	}

	result := make([]string, 0, numOfResult)
	for i := 0; i < len(timetable); i++ {
		v := timetable[i]
		if v.minute > minute && v.hour >= hour {
			timeStr := v.toString()
			result = append(result, timeStr)
			if len(result) >= numOfResult {
				break
			}
		}
	}
	printTimes(result)
}

package main

import (
	"testing"
	"time"
)

func TestDepartureTime(t *testing.T) {
	deptime := departureTime("10:30")
	expected_time := Time{10, 30}
	if deptime != expected_time {
		t.Errorf("got %v\nwant %v", deptime, expected_time)
	}
}

func TestDepartureTime2(t *testing.T) {
	deptime := departureTime("0:0")
	expected_time := Time{0, 0}
	if deptime != expected_time {
		t.Errorf("got %v\nwant %v", deptime, expected_time)
	}
}

func TestDepartureTime3(t *testing.T) {
	deptime := departureTime("00:00")
	expected_time := Time{0, 0}
	if deptime != expected_time {
		t.Errorf("got %v\nwant %v", deptime, expected_time)
	}
}

func TestDepartureTime4(t *testing.T) {
	deptime := departureTime("23:59")
	expected_time := Time{23, 59}
	if deptime != expected_time {
		t.Errorf("got %v\nwant %v", deptime, expected_time)
	}
}

func TestDepartureTime5(t *testing.T) {
	deptime := departureTime("44:30")
	now := time.Now()
	expected_time := Time{now.Hour(), now.Minute()}
	if deptime != expected_time {
		t.Errorf("got %v\nwant %v", deptime, expected_time)
	}
}

func TestDepartureTime6(t *testing.T) {
	deptime := departureTime("24:30")
	now := time.Now()
	expected_time := Time{now.Hour(), now.Minute()}
	if deptime != expected_time {
		t.Errorf("got %v\nwant %v", deptime, expected_time)
	}
}

func TestDepartureTime7(t *testing.T) {
	deptime := departureTime("12:60")
	now := time.Now()
	expected_time := Time{now.Hour(), now.Minute()}
	if deptime != expected_time {
		t.Errorf("got %v\nwant %v", deptime, expected_time)
	}
}

func TestToString(t *testing.T) {
	value := Time{hour: 8, minute: 8}
	ret := value.toString()
	if ret != "8:08 " {
		t.Errorf("got %v, wrong Format", ret)
	}
}

func TestToString2(t *testing.T) {
	value := Time{10, 10}
	ret := value.toString()
	if ret != "10:10 " {
		t.Errorf("got %v, wrong Format", ret)
	}
}

func TestCreateTimetable(t *testing.T) {
	timetable := createTimetable()
	if len(timetable) == 0 {
		t.Error("Timetable is nil. Current URL may be out of date.")
	}
}

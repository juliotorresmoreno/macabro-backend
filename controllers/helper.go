package controllers

import (
	"strconv"
	"time"
)

var zone, _ = time.Now().Local().Zone()

// Response s
type Response map[string]interface{}

func newValueString(n, a string) string {
	if n != "" {
		return n
	}
	return a
}
func newValueTime(n, a time.Time) time.Time {
	hours, _ := strconv.Atoi(zone)
	if n.Year() > 0 {
		return n.Add(time.Duration(hours*-1) * time.Hour)
	}
	return a
}

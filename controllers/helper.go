package controllers

import "time"

// Response s
type Response map[string]interface{}

func newValueString(n, a string) string {
	if n != "" {
		return n
	}
	return a
}
func newValueTime(n, a time.Time) time.Time {
	if n.Year() > 0 {
		return n
	}
	return a
}

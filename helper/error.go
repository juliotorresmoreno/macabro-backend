package helper

import (
	"errors"
	"fmt"
	"strings"
)

// ParseError s
func ParseError(err error) error {
	errString := err.Error()
	if errString[:17] == "pq: duplicate key" {
		s := strings.Split(errString, "\"")
		s = strings.Split(s[1], "_")
		f := strings.Join(s[2:], "_")
		return fmt.Errorf("%v: %v already exists", f, f)
	}
	if errString[:8] == "dial tcp" {
		return errors.New("Database is not running")
	}
	return err
}

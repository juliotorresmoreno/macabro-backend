package models

import (
	"fmt"
	"regexp"
	"unicode"

	"github.com/asaskevich/govalidator"
)

func verifyPassword(s string) (sevenOrMore, number, upper, special bool) {
	letters := 0
	for _, c := range s {
		switch {
		case unicode.IsNumber(c):
			number = true
		case unicode.IsUpper(c):
			upper = true
			letters++
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			special = true
		case unicode.IsLetter(c) || c == ' ':
			letters++
		default:
			//return false, false, false, false
		}
	}
	sevenOrMore = letters >= 7
	return
}

func init() {
	pname, _ := regexp.Compile("^[a-zA-Z]+(([' ][a-zA-Z ])?[a-zA-Z]*)*$")
	govalidator.CustomTypeTagMap.Set("name", func(i, o interface{}) bool {
		v := fmt.Sprintf("%v", i)
		return pname.MatchString(v)
	})

	pusername, _ := regexp.Compile("^[a-zA-Z][a-zA-Z0-9\\.]{3,}$")
	govalidator.CustomTypeTagMap.Set("username", func(i, o interface{}) bool {
		v := fmt.Sprintf("%v", i)
		return pusername.MatchString(v)
	})

	govalidator.CustomTypeTagMap.Set("password", func(i, o interface{}) bool {
		v := fmt.Sprintf("%v", i)
		sevenOrMore, number, upper, special := verifyPassword(v)
		return sevenOrMore && number && upper && special
	})
}

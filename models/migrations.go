package models

import (
	"log"

	"github.com/juliotorresmoreno/macabro/db"
)

// Migrate s
func Migrate() {
	conn, err := db.NewEngigne()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	err = conn.Sync2(User{})
	if err != nil {
		log.Fatal(err)
	}
}

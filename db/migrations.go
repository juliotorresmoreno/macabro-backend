package db

import (
	"log"

	"github.com/juliotorresmoreno/macabro/models"
)

// Migrate s
func Migrate() {
	conn, err := NewEngigne()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	err = conn.Sync2(models.User{})
	if err != nil {
		log.Fatal(err)
	}
	err = conn.Sync2(models.Business{})
	if err != nil {
		log.Fatal(err)
	}
	err = conn.Sync2(models.PaymentMethods{})
	if err != nil {
		log.Fatal(err)
	}
	err = conn.Sync2(models.Instance{})
	if err != nil {
		log.Fatal(err)
	}
}

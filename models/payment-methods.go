package models

import (
	"encoding/json"
	"time"

	creditcard "github.com/durango/go-credit-card"
	"github.com/juliotorresmoreno/macabro/helper"
)

type PaymentMethods struct {
	ID     int `xorm:"id integer not null autoincr pk"                    valid:""`
	UserID int `xorm:"user_id integer not null index"                     valid:""`

	Name            string `xorm:"name varchar(100) not null"             valid:""`
	AliasNumber     string `xorm:"alias_number varchar(100) not null"     valid:""`
	Type            string `xorm:"type varchar(100) not null"             valid:""`
	Default         int    `xorm:"'default' integer not null"               valid:""`
	Number          string `xorm:"number varchar(100) not null"           valid:""`
	ExpirationMonth string `xorm:"expiration_month varchar(100) not null" valid:""`
	ExpirationYear  string `xorm:"expiration_year varchar(100) not null"  valid:""`
	CVV             string `xorm:"cvv varchar(100) not null"              valid:""`

	ACL       ACL       `xorm:"acl json not null" valid:"required"`
	CreatedAt time.Time `xorm:"created_at created"`
	UpdatedAt time.Time `xorm:"updated_at updated"`
	Version   int       `xorm:"version version"`
}

func (that PaymentMethods) TableName() string {
	return "payment-methods"
}

type paymentMethods struct {
	ID              int    `json:"id"`
	UserID          int    `json:"user_id"`
	Name            string `json:"name"`
	AliasNumber     string `json:"alias_number"`
	Type            string `json:"type"`
	Default         int    `json:"default"`
	Number          string `json:"number"`
	ExpirationMonth string `json:"expiration_month"`
	ExpirationYear  string `json:"expiration_year"`
	CVV             string `json:"cvv"`
}

func (that PaymentMethods) MarshalJSON() ([]byte, error) {
	u := &paymentMethods{
		ID:          that.ID,
		UserID:      that.UserID,
		Name:        that.Name,
		AliasNumber: that.AliasNumber,
		Type:        that.Type,
		Default:     that.Default,
	}
	return json.Marshal(u)
}

func (that *PaymentMethods) UnmarshalJSON(data []byte) error {
	u := &paymentMethods{}
	err := json.Unmarshal(data, u)
	if err != nil {
		return err
	}
	that.ID = u.ID
	that.UserID = u.UserID
	that.Name = u.Name

	lastFour := ""
	if u.Number != "" {
		card := creditcard.Card{
			Number: u.Number,
			Cvv:    u.CVV,
			Month:  u.ExpirationMonth,
			Year:   u.ExpirationYear,
		}

		err = card.Method()
		if err != nil {
			return err
		}

		lastFour, err = card.LastFour()
		if err != nil {
			return err
		}

		err = card.Validate(true)
		if err != nil {
			return err
		}
		that.Type = card.Company.Short
	}

	that.AliasNumber = "**** **** **** " + lastFour

	that.Number = u.Number
	that.ExpirationYear = u.ExpirationYear
	that.ExpirationMonth = u.ExpirationMonth
	that.CVV = u.CVV
	that.Default = u.Default
	return nil
}

func (that *PaymentMethods) Decrypt(secret string) error {
	key := helper.GetAesKey(secret)
	that.Number, _ = helper.Decrypt(key, that.Number)
	that.CVV, _ = helper.Decrypt(key, that.CVV)
	that.ExpirationYear, _ = helper.Decrypt(key, that.ExpirationYear)
	that.ExpirationMonth, _ = helper.Decrypt(key, that.ExpirationMonth)
	return nil
}

func (that *PaymentMethods) Encrypt(secret string) error {
	key := helper.GetAesKey(secret)
	that.Number, _ = helper.Encrypt(key, that.Number)
	that.CVV, _ = helper.Encrypt(key, that.CVV)
	that.ExpirationYear, _ = helper.Encrypt(key, that.ExpirationYear)
	that.ExpirationMonth, _ = helper.Encrypt(key, that.ExpirationMonth)
	return nil
}

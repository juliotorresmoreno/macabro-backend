package models

import (
	"encoding/json"
	"time"
)

type PaymentMethods struct {
	ID              int    `xorm:"id integer not null autoincr pk"        valid:""`
	UserID          int    `xorm:"user_id integer not null"               json:"user_id"`
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
	Number          string `json:"number"`
	ExpirationMonth string `json:"expiration_month"`
	ExpirationYear  string `json:"expiration_year"`
	CVV             string `json:"cvv"`
}

func (that PaymentMethods) MarshalJSON() ([]byte, error) {
	u := &paymentMethods{
		ID:              that.ID,
		UserID:          that.UserID,
		Number:          that.Number,
		ExpirationYear:  that.ExpirationYear,
		ExpirationMonth: that.ExpirationMonth,
		CVV:             that.CVV,
	}
	return json.Marshal(u)
}

func (that PaymentMethods) UnmarshalJSON(data []byte) error {
	u := &paymentMethods{
		ID:              that.ID,
		UserID:          that.UserID,
		Number:          that.Number,
		ExpirationYear:  that.ExpirationYear,
		ExpirationMonth: that.ExpirationMonth,
		CVV:             that.CVV,
	}
	err := json.Unmarshal(data, u)
	if err != nil {
		return err
	}
	return nil
}

func (that *PaymentMethods) Decrypt() error {
	return nil
}

func (that *PaymentMethods) Encrypt() error {
	return nil
}

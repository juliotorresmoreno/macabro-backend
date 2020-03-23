package models

import (
	"encoding/json"
	"time"
)

type Business struct {
	ID                  int    `xorm:"id integer not null autoincr pk <-" json:"id"`
	UserID              int    `xorm:"user_id integer not null unique"    json:"user_id"`
	Name                string `xorm:"name varchar(255)"                  json:"name"`
	Nit                 string `xorm:"nit varchar(20)"                    json:"nit"`
	LegalRepresentative string `xorm:"legal_representative varchar(255)"  json:"legal_representative"`
	Website             string `xorm:"website text"                       json:"website"`
	Address             string `xorm:"address varchar(255)"               json:"address"`
	Country             string `xorm:"country varchar(2)"                 json:"country"`
	City                string `xorm:"city varchar(100)"                  json:"city"`
	EconomicActivity    string `xorm:"economic_activity varchar(500)"     json:"economic_activity"`
	ImgSrc              string `xorm:"imgSrc text"                        json:"imgSrc"`

	ACL       ACL       `xorm:"acl json not null"  json:"-" valid:"required"`
	CreatedAt time.Time `xorm:"created_at created" json:"-"`
	UpdatedAt time.Time `xorm:"updated_at updated" json:"-"`
	Version   int       `xorm:"version version"    json:"-"`
}

type business struct {
	ID                  int    `json:"id"`
	UserID              int    `json:"user_id"`
	Name                string `json:"name"`
	Nit                 string `json:"nit"`
	LegalRepresentative string `json:"legal_representative"`
	Website             string `json:"website"`
	Address             string `json:"address"`
	Country             string `json:"country"`
	City                string `json:"city"`
	EconomicActivity    string `json:"economic_activity"`
	ImgSrc              string `json:"imgSrc"`
}

func (el Business) TableName() string {
	return "business"
}

func (el Business) Check() error {
	return nil
}

// UnmarshalJSON s
func (that *Business) UnmarshalJSON(b []byte) error {
	u := &business{}
	err := json.Unmarshal(b, u)
	if err != nil {
		return err
	}
	that.ID = u.ID
	that.UserID = u.UserID
	that.Name = u.Name
	that.Nit = u.Nit
	that.LegalRepresentative = u.LegalRepresentative
	that.Website = u.Website
	that.Address = u.Address
	that.Country = u.Country
	that.City = u.City
	that.EconomicActivity = u.EconomicActivity
	that.ImgSrc = u.ImgSrc

	return nil
}

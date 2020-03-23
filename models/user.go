package models

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/asaskevich/govalidator"
	"golang.org/x/crypto/bcrypt"
)

type Group struct {
	Read  bool `json:"read"`
	Write bool `json:"write"`
}

// ACL s
type ACL struct {
	Owner  string           `json:"owner"`
	Group  string           `json:"group"`
	Groups map[string]Group `json:"groups"`
}

func (that ACL) IsAdmin() bool {
	s := strings.Split(that.Group, ",")
	for _, g := range s {
		if g == "admin" {
			return true
		}
	}
	return false
}

func NewACL(user string, groups ...string) ACL {
	group := ""
	if len(groups) > 0 {
		group = strings.Join(groups, ",")
	}
	ACL := ACL{
		Owner: user,
		Group: group,
		Groups: map[string]Group{
			"user": Group{
				Read:  true,
				Write: false,
			},
			"admin": Group{
				Read:  true,
				Write: true,
			},
		},
	}
	return ACL
}

// User s
type User struct {
	ID       int    `xorm:"id integer not null autoincr pk"      valid:""`
	Username string `xorm:"username varchar(20) not null unique" valid:"username,required"`
	Email    string `xorm:"email varchar(200) not null unique"   valid:"email,required"`
	Name     string `xorm:"name varchar(50) not null"            valid:"name,required"`
	LastName string `xorm:"lastname varchar(50) not null"        valid:"name,required"`

	DocumentType string    `xorm:"document_type varchar(2)"      valid:"in(CC|CE|PA|RC|TI)"`
	Expedite     time.Time `xorm:"expedite date"`
	Document     string    `xorm:"document varchar(20)"          valid:"int"`
	DateBirth    time.Time `xorm:"date_birth date"`
	ImgSrc       string    `xorm:"imgSrc text"`
	Country      string    `xorm:"country varchar(2)"`
	Nationality  string    `xorm:"nationality varchar(2)"`
	Facebook     string    `xorm:"facebook varchar(255)"`
	Linkedin     string    `xorm:"linkedin varchar(255)"`

	Password      string `xorm:"password varchar(100) not null"`
	ValidPassword string `xorm:"-"                              valid:"password"`
	RecoveryToken string `xorm:"recovery_token varchar(100) not null"`

	ACL       ACL       `xorm:"acl json not null"                    valid:"required"`
	CreatedAt time.Time `xorm:"created_at created"`
	UpdatedAt time.Time `xorm:"updated_at updated"`
	Version   int       `xorm:"version version"`
}

// TableName s
func (that User) TableName() string {
	return "users"
}

type user struct {
	ID       int    `json:"id"`
	ACL      ACL    `json:"acl"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Name     string `json:"name"`
	LastName string `json:"lastname"`

	DocumentType string    `json:"document_type"`
	Expedite     time.Time `json:"expedite"`
	Document     string    `json:"document"`
	DateBirth    time.Time `json:"date_birth"`
	ImgSrc       string    `json:"imgSrc"`
	Country      string    `json:"country"`
	Nationality  string    `json:"nationality"`
	Facebook     string    `json:"facebook"`
	Linkedin     string    `json:"linkedin"`

	Password string `json:"password"`
}

type userWithowPassword struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Name     string `json:"name"`
	LastName string `json:"lastname"`

	DocumentType string    `json:"document_type"`
	Expedite     time.Time `json:"expedite"`
	Document     string    `json:"document"`
	DateBirth    time.Time `json:"date_birth"`
	ImgSrc       string    `json:"imgSrc"`
	Country      string    `json:"country"`
	Nationality  string    `json:"nationality"`
	Facebook     string    `json:"facebook"`
	Linkedin     string    `json:"linkedin"`
}

// Check s
func (that *User) Check() error {
	_, err := govalidator.ValidateStruct(that)
	return err
}

// SetPassword s
func (that *User) SetPassword(password string) error {
	s, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}
	that.Password = string(s)
	return nil
}

// UnmarshalJSON s
func (that *User) UnmarshalJSON(b []byte) error {
	u := &user{}
	err := json.Unmarshal(b, u)
	if err != nil {
		return err
	}
	that.ID = 0
	that.Username = u.Username
	that.Email = u.Email
	that.Name = u.Name
	that.LastName = u.LastName
	that.ValidPassword = u.Password

	that.DocumentType = u.DocumentType
	that.Expedite = u.Expedite
	that.Document = u.Document
	that.DateBirth = u.DateBirth
	that.ImgSrc = u.ImgSrc
	that.Country = u.Country
	that.Nationality = u.Nationality
	that.Facebook = u.Facebook
	that.Linkedin = u.Linkedin

	that.ACL = NewACL(u.Username, "user")
	err = that.SetPassword(u.Password)
	if err != nil {
		return err
	}
	return nil
}

// MarshalJSON s
func (that User) MarshalJSON() ([]byte, error) {
	u := userWithowPassword{
		ID:       that.ID,
		Email:    that.Email,
		Username: that.Username,
		Name:     that.Name,
		LastName: that.LastName,

		DocumentType: that.DocumentType,
		Expedite:     that.Expedite,
		Document:     that.Document,
		DateBirth:    that.DateBirth,
		ImgSrc:       that.ImgSrc,
		Country:      that.Country,
		Nationality:  that.Nationality,
		Facebook:     that.Facebook,
		Linkedin:     that.Linkedin,
	}
	return json.Marshal(u)
}

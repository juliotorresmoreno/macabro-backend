package models

import (
	"encoding/json"

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
	Groups map[string]Group `json:"groups"`
}

// User s
type User struct {
	ID            uint   `xorm:"id integer not null autoincr pk"      valid:""`
	ACL           ACL    `xorm:"acl json not null"                    valid:"required"`
	Username      string `xorm:"username varchar(20) not null unique" valid:"username,required"`
	Email         string `xorm:"email varchar(200) not null unique"   valid:"email,required"`
	Name          string `xorm:"name varchar(50) not null"            valid:"name,required"`
	LastName      string `xorm:"lastname varchar(50) not null"        valid:"name,required"`
	Password      string `xorm:"password varchar(100) not null"`
	ValidPassword string `xorm:"-"                                    valid:"password"`
	RecoveryToken string `xorm:"recovery_token varchar(100) not null"`
}

// TableName s
func (that User) TableName() string {
	return "users"
}

type user struct {
	ID       uint   `json:"id"`
	ACL      ACL    `json:"acl"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Name     string `json:"name"`
	LastName string `json:"lastname"`
	Password string `json:"password"`
}

type userWithowPassword struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Name     string `json:"name"`
	LastName string `json:"lastname"`
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
	that.ACL = ACL{
		Owner:  u.Username,
		Groups: map[string]Group{},
	}
	that.ACL.Groups["admin"] = Group{
		Read:  true,
		Write: true,
	}
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
	}
	return json.Marshal(u)
}

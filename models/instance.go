package models

import "time"

// Instance _
type Instance struct {
	ID int `xorm:"id integer not null autoincr pk <-" json:"id"`

	User              *User  `xorm:"user_id"            `
	Name              string `xorm:"name varchar(255) not null"          valid:"name,required"`
	Type              string `xorm:"type varchar(20) not null"           valid:"in(small|micro|medium|large|xlarge)"`
	Replicas          int    `xorm:"replicas integer not null"           valid:"min(1),max(20)"`
	AutoScaling       int    `xorm:"auto_scaling integer not null"       valid:"min(0),max(1)"`
	AllowDeletion     int    `xorm:"allow_deletion integer not null"     valid:"min(0),max(1)"`
	BackupPeriodicity int    `xorm:"backup_periodicity integer not null" valid:"in(12|24|72|168|336|720)"`
	URL               string `xorm:"url string not null"                 valid:"name,required"`
	Username          string `xorm:"username string not null"            valid:"username,required"`
	Password          string `xorm:"password string not null"            valid:"password,required"`

	ACL       ACL       `xorm:"acl json not null"  json:"-" valid:"required"`
	CreatedAt time.Time `xorm:"created_at created" json:"-"`
	UpdatedAt time.Time `xorm:"updated_at updated" json:"-"`
	Version   int       `xorm:"version version"    json:"-"`
}

func TableName() string {
	return "instances"
}

func (el Instance) Check() error {
	return nil
}

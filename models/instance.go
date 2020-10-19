package models

import "time"

// Instance _
type Instance struct {
	ID                int       `xorm:"id integer not null autoincr pk <-"                                              json:"id"                `
	User              *User     `xorm:"user_id"                                                                         json:"-"                 `
	IsCloud           bool      `xorm:"is_cloud"                                                                        json:"is_cloud"          `
	Name              string    `xorm:"name varchar(255) not null"          valid:"name,required"                       json:"name"              `
	Type              string    `xorm:"type varchar(20) not null"           valid:"in(small|micro|medium|large|xlarge)" json:"type"              `
	Replicas          int       `xorm:"replicas integer not null"           valid:"min(1),max(20)"                      json:"replicas"          `
	AutoScaling       int       `xorm:"auto_scaling integer not null"       valid:"min(0),max(1)"                       json:"auto_scaling"      `
	AllowDeletion     int       `xorm:"allow_deletion integer not null"     valid:"min(0),max(1)"                       json:"allow_deletion"    `
	BackupPeriodicity int       `xorm:"backup_periodicity integer not null" valid:"in(12|24|72|168|336|720)"            json:"backup_periodicity"`
	URL               string    `xorm:"url varchar(255) not null"                 valid:"name,required"                 json:"url"         `
	Username          string    `xorm:"username varchar(100) not null"            valid:"username,required"             json:"username"    `
	Password          string    `xorm:"password varchar(100) not null"            valid:"password,required"             json:"password"    `
	ACL               ACL       `xorm:"acl json not null"                                                               json:"-" valid:"required"`
	CreatedAt         time.Time `xorm:"created_at created"                                                              json:"-"                 `
	UpdatedAt         time.Time `xorm:"updated_at updated"                                                              json:"-"                 `
	Version           int       `xorm:"version version"                                                                 json:"-"                 `
}

// InstanceWithDate .
type InstanceWithDate struct {
	ID                int       `xorm:"id integer not null autoincr pk <-"  json:"id"                `
	User              *User     `xorm:"user_id"                             json:"-"                 `
	IsCloud           bool      `xorm:"is_cloud"                            json:"is_cloud"          `
	Name              string    `xorm:"name varchar(255) not null"          json:"name"              `
	Type              string    `xorm:"type varchar(20) not null"           json:"type"              `
	Replicas          int       `xorm:"replicas integer not null"           json:"replicas"          `
	AutoScaling       int       `xorm:"auto_scaling integer not null"       json:"auto_scaling"      `
	AllowDeletion     int       `xorm:"allow_deletion integer not null"     json:"allow_deletion"    `
	BackupPeriodicity int       `xorm:"backup_periodicity integer not null" json:"backup_periodicity"`
	URL               string    `xorm:"url varchar(255) not null"           json:"url"               `
	Username          string    `xorm:"username varchar(100) not null"      json:"username"          `
	Password          string    `xorm:"password varchar(100) not null"      json:"password"          `
	ACL               ACL       `xorm:"acl json not null"                   json:"-" valid:"required"`
	CreatedAt         time.Time `xorm:"created_at created"                  json:"created_at"        `
	UpdatedAt         time.Time `xorm:"updated_at updated"                  json:"updated_at"        `
	Version           int       `xorm:"version version"                     json:"version"           `
}

// TableName _
func (el Instance) TableName() string {
	return "instances"
}

// TableName _
func (el InstanceWithDate) TableName() string {
	return "instances"
}

// Check _
func (el Instance) Check() error {
	return nil
}

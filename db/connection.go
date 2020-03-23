package db

import (
	"github.com/go-xorm/xorm"
	"github.com/juliotorresmoreno/macabro/config"
	"github.com/lib/pq"
)

// NewEngigne s
func NewEngigne() (*xorm.Engine, error) {
	conf := config.GetConfig()
	dsn := conf.Database.DSN
	if conf.Database.Driver == "postgres" {
		dsn, _ = pq.ParseURL(conf.Database.DSN)
	}
	conn, err := xorm.NewEngine(conf.Database.Driver, dsn)
	return conn, err
}

// NewEngigneWithSession s
func NewEngigneWithSession(user, group string) (*Engine, error) {
	conn, err := NewEngigne()

	r := "(acl->>'owner' = '%v' or (acl->'groups'->'%v'->>'read')::boolean is true)"
	w := "(acl->>'owner' = '%v' or (acl->'groups'->'%v'->>'write')::boolean is true)"

	engine := &Engine{Engine: conn}
	engine.permisionQueryRead = r
	engine.permisionQueryWrite = w
	engine.user = user
	engine.group = group
	engine.ShowSQL(true)

	return engine, err
}

package db

import (
	"errors"

	"github.com/go-xorm/xorm"
	"github.com/juliotorresmoreno/macabro/config"
	"github.com/lib/pq"
)

// NewEngigne s
func NewEngigne() (*xorm.Engine, error) {
	conf := config.GetConfig()
	if conf.Database.Driver == "postgres" {
		dsn, err := pq.ParseURL(conf.Database.DSN)
		if err != nil {
			return &xorm.Engine{}, err
		}
		conn, err := xorm.NewEngine(conf.Database.Driver, dsn)
		conn.ShowSQL(true)
		return conn, err
	}
	return &xorm.Engine{}, errors.New("No implementado")
}

// NewEngigneWithSession s
func NewEngigneWithSession(user, group string) (*Engine, error) {
	conn, err := NewEngigne()

	pQueryRead := "(acl->>'owner' = '%v' or (acl->'groups'->'%v'->>'read')::boolean is true)"
	pQueryWrite := "(acl->>'owner' = '%v' or (acl->'groups'->'%v'->>'write')::boolean is true)"

	engine := &Engine{Engine: conn}
	engine.permisionQueryRead = pQueryRead
	engine.permisionQueryWrite = pQueryWrite
	engine.user = user
	engine.group = group
	engine.ShowSQL(true)

	return engine, err
}

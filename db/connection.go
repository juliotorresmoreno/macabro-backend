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
	return xorm.NewEngine(conf.Database.Driver, dsn)
}

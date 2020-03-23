package bootstrap

import (
	"github.com/juliotorresmoreno/macabro/config"
	"github.com/juliotorresmoreno/macabro/db"
)

// Init s
func Init() {
	config.Init("./config.yml")
	db.Migrate()
}

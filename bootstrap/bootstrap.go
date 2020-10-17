package bootstrap

import (
	"time"

	"github.com/juliotorresmoreno/macabro/config"
	"github.com/juliotorresmoreno/macabro/db"
)

// Init s
func Init() {
	config.Init("./config.yml")
	time.LoadLocation("America/Bogota")
	db.Migrate()
}

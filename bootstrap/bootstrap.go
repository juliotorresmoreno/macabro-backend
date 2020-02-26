package bootstrap

import (
	"github.com/juliotorresmoreno/macabro/config"
	"github.com/juliotorresmoreno/macabro/models"
)

// Init s
func Init() {
	config.Init("./config.yml")
	models.Migrate()
}

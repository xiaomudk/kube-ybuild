package config

import (
	kconfig "github.com/xiaomudk/kube-ybuild/pkg/config"
	logger "github.com/xiaomudk/kube-ybuild/pkg/logs"
	"github.com/xiaomudk/kube-ybuild/pkg/orm"
)

var (
	// conf conf var
	Conf *Config
)

type Config struct {
	Address  string         `json:"address" env:"ADDR"`
	Database *orm.Config    `json:"database"`
	Log      *logger.Config `json:"log"`
}

// Init initialize the routing of this application.
func Init(filename string) {
	c := kconfig.New("./")
	if err := c.Load(filename, &Conf); err != nil {
		panic(err)
	}
}

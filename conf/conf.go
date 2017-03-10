package conf

import (
	"flag"

	"github.com/BurntSushi/toml"
)

// Conf global variable.
var (
	Conf     *Config
	confPath string
)

// Config struct of conf.
type Config struct {
	User *User
}

type User struct {
	username string
	passwd   string
}

func init() {
	flag.StringVar(&confPath, "conf", "", "default config path")
}

// Init create config instance.
func Init() (err error) {
	_, err = toml.DecodeFile(confPath, &Conf)
	return
}

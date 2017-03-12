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
	User  *User
	Mysql *MySQL
}

// MySQL config.
type MySQL struct {
	Addr   string // for trace
	DSN    string // data source name
	Active int    // pool
	Idle   int    // pool
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

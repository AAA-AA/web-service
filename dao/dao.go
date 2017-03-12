package dao

import (
	"database/sql"

	"web-service/conf"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ngaut/log"
)

type Dao struct {
	db             *sql.DB
	getSecUserStmt *sql.Stmt
	c              *conf.Config
}

func New(c *conf.Config) (d *Dao) {
	var err error
	d = &Dao{
		db: NewMysql(c.Mysql),
		c:  c,
	}
	d.getSecUserStmt, err = d.db.Prepare(_getSecUserByNameSQL)
	if err != nil {
		log.Errorf("d.db.Prepare error %v", err)
		return
	}
	return
}

func (d *Dao) Close() {
	if d.db != nil {
		d.db.Close()
	}
}

func (d *Dao) Ping()  {
	var err error
	if err =  d.db.Ping();err != nil {
		log.Errorf("d.db.Ping error:%v",err)
		return
	}
	log.Info("d.db.Ping success!")
}

// NewMySQL new db and retry connection when has error.
func NewMysql(c *conf.MySQL) (db *sql.DB) {
	var err error
	if db, err = sql.Open("mysql", c.DSN); err != nil {
		log.Error("open mysql error(%v)", err)
		panic(err)
	}
	db.SetMaxOpenConns(c.Active)
	db.SetMaxIdleConns(c.Idle)
	return
}

package dao

import (
	"context"
	"testing"
	"web-service/conf"
	"web-service/model"
)

func TestDao_SecUser(t *testing.T) {
	initConf(t)
	dao := New(conf.Conf)

	var (
		name    = "%M%"
		secUser = &model.SecUser{}
		c       = context.TODO()
		err     error
	)
	if err = dao.db.Ping(); err != nil {
		t.Errorf("connected to db error :%v", err)
	}
	secUser, err = dao.SecUser(c, name)
	if err != nil {
		t.Errorf("secUser error :%v", err)
	}
	t.Logf("test succed :%v", secUser)

}

func initConf(t *testing.T) {
	if err := conf.Init(); err != nil {
		t.Errorf("conf.Init() error(%v)", err)
		t.FailNow()
	}
}

package service

import (
	"web-service/conf"
	"web-service/dao"
)

type Service struct {
	dao *dao.Dao
	c   *conf.Config
}

func New(c *conf.Config) (s *Service) {
	s = &Service{
		dao: dao.New(c),
	}
	return
}

func (s *Service) Close() {
	s.dao.Close()
	return
}

package service

import (
	"fmt"
	"web-service/model"
)

func (s *Service) Login(userName string) (ok bool, err error) {
	var (
		secUser *model.SecUser
	)
	secUser, err = s.dao.SecUser(userName)
	if err != nil {
		ok = false
		return
	}
	fmt.Println("userName is :" + userName)
	if secUser != nil && secUser.Name == userName {
		ok = true
		fmt.Printf("user is %v", secUser)
		return
	}
	ok = false
	return
}

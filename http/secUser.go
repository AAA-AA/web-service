package http

import (
	"fmt"
	"net/http"
	"regexp"
	"text/template"

	"web-service/conf"
	"web-service/service"

	"github.com/ngaut/log"
)

func login(w http.ResponseWriter, r *http.Request) {
	var (
		service  = service.New(conf.Conf)
		userName string
	)
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		t, _ := template.ParseFiles("html/template/index.html")
		log.Info(t.Execute(w, nil))
	} else {
		r.ParseForm()
		log.Info("post!")
		//请求的是登陆数据，那么执行登陆的逻辑判断
		fmt.Println("username:", r.Form["userName"])
		userName = r.FormValue("userName")
		fmt.Printf("len of username:%d", len(r.Form["userName"][0]))
		fmt.Println()
		if len(r.Form["userName"][0]) < 3 {
			fmt.Fprintf(w, "用户名过短")
			return
		}
		fmt.Println("password:", r.Form["password"])
		if m, _ := regexp.MatchString("^[0-9]+|[a-zA-Z]+$", r.Form.Get("password")); !m {
			fmt.Fprintf(w, "密码必须为数字或字母")
			return
		}
	}
	if ok, _ := service.Login(userName); ok {
		fmt.Fprintf(w, "登录成功")
		return
	}
	t, _ := template.ParseFiles("html/template/index.html")
	log.Info(t.Execute(w, nil))

}

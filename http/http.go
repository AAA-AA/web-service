package http

import (
	"fmt"
	"html/template"
	"net/http"

	"web-service/conf"

	"github.com/ngaut/log"
)

func Init(c *conf.Config) {
	rooter()
}
func rooter() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Error("ListenAndServe error ", err)
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("request path is :", r.URL.Path)
	fmt.Println("the scheme is :", r.URL.Scheme)
	for k, v := range r.Form {
		fmt.Println("key is :", k)
		fmt.Println("value is ", v)
	}
	if r.Method == "GET" {
		t, _ := template.ParseFiles("html/template/login.html")
		log.Info(t.Execute(w, nil))
	}

}

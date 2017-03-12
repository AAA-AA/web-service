package model

import "time"

type SecUser struct {
	Id              int64     `json:"id"`
	Name            string    `json:"name"`
	Version         int64     `json:"version"`
	Mark            string    `json:"mark"`
	Operator        string    `json:"operator"`
	Last_login_date time.Time `json:"status"`
	Ctime           time.Time `json:"ctime"`
	Mtime           time.Time `json:"mtime"`
	Status          int8      `json:"status"`
}

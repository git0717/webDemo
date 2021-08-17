package main

import (
	"time"
)

type Todo struct {

	Id int `json:"id"`
	Name string `json:"name"`
	Completed   bool    `json:"completed"`
	Due time.Time `json:"due"`
}

type Todos []Todo

type BaseJsonBean struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func NewBaseJsonBean() *BaseJsonBean {
	return &BaseJsonBean{}
}
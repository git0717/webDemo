package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"time"
)

func Index(w http.ResponseWriter,r *http.Request){
	fmt.Fprint(w,"Welcome")
}
func TodoIndex(w http.ResponseWriter,r *http.Request){
	todos :=Todos{
		Todo{Name:"Write presentation"},
		Todo{Name:"Host meetup"},
		Todo{Id:22,Name: "joel",Completed:true,Due:time.Now()},
	}
	if err:=json.NewEncoder(w).Encode(todos); err!=nil{
		panic(err)
	}
}
func TodoShow(w http.ResponseWriter,r *http.Request){
	vars := mux.Vars(r)
	todoId :=vars["todoId"]
	fmt.Fprintln(w,"Todo show:",todoId)
}
func TodoPostJson(w http.ResponseWriter,req *http.Request){
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {

	} else {
		fmt.Fprint(w, bytes.NewBuffer(data).String())
	}
	var user map[string]interface{}
	//body, _ := ioutil.ReadAll(req.Body)
	json.Unmarshal(data, &user)
	fmt.Println("获取json中的username:", user["username"])
	fmt.Println("获取json中的password:", user["password"].(string)) //转字符串通过len(password)!=0判断长度

	result := NewBaseJsonBean()
	userName :=user["username"]
	password :=user["password"]
	//var s = "userName:" + userName + ",password:" + password
	//fmt.Println(s)

	if userName == "zhangsan" && password == "123456" {
		result.Code = 100
		result.Message = "登录成功"
	} else {
		result.Code = 101
		result.Message = "用户名或密码不正确"
	}

	//向客户端返回JSON数据
	bytes, _ := json.Marshal(result)
	fmt.Fprint(w, string(bytes))

}
func TodoPostForm(w http.ResponseWriter,req *http.Request){
	//vars := mux.Vars(req)
	//todoId :=vars["todoId"]
	//fmt.Fprintln(w,"Todo show:",todoId)
	//if req.Method == "POST" {
	//	var (
	//		key   string = req.PostFormValue("key")
	//		value string = req.PostFormValue("value")
	//	)
	//	fmt.Printf("key is  : %s\n", key)
	//	fmt.Printf("value is: %s\n", value)
	//}
	//获取客户端通过GET/POST方式传递的参数
	body := req.GetBody
	fmt.Println(body)

	closer := req.Body
	fmt.Println(closer)
	req.ParseForm()
	param_userName, found1 := req.Form["userName"]
	param_password, found2 := req.Form["password"]

	if !(found1 && found2) {
		fmt.Fprint(w, "请勿非法访问")
		return
	}

	result := NewBaseJsonBean()
	userName := param_userName[0]
	password := param_password[0]

	s := "userName:" + userName + ",password:" + password
	fmt.Println(s)

	if userName == "zhangsan" && password == "123456" {
		result.Code = 100
		result.Message = "登录成功"
	} else {
		result.Code = 101
		result.Message = "用户名或密码不正确"
	}

	//向客户端返回JSON数据
	bytes, _ := json.Marshal(result)
	fmt.Fprint(w, string(bytes))

}

package main

import (
	_ "fmt"
	_ "github.com/gin-gonic/gin"
	_ "github.com/gorilla/mux"
	_ "gorm.io/gorm"
	_ "html"
	"log"
	_ "log"
	"net/http"
	_ "net/http"
)
func main() {
	//http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
	//	fmt.Fprintln(writer, "Hello, ", html.EscapeString(request.URL.Path))
	//})
	//log.Fatal(http.ListenAndServe(":8080",nil))

	//router :=mux.NewRouter().StrictSlash(true)
	//router.HandleFunc("/",Index);
	//router.HandleFunc("/todos",TodoShow)
	//router.HandleFunc("/todo/{todoId}",TodoShow)
	//log.Fatal(http.ListenAndServe(":8080",router))
	router :=NewRouter()
	log.Fatal(http.ListenAndServe(":8080",router))

}
//func TodoShow(writer http.ResponseWriter,request *http.Request){
//	vars :=mux.Vars(request)
//	todoId :=vars["todoId"]
//	fmt.Fprintln(writer,"Todo show:",todoId)
//}
//
//func todoIndex(writer http.ResponseWriter,request *http.Request)  {
// fmt.Fprintln(writer,"Todo Index!",request.URL.Path)
//}
//
//func Index(writer http.ResponseWriter,request *http.Request){
//	fmt.Fprintln(writer,"welcome")

//}
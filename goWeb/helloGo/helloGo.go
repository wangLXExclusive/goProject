package main

//func main(){
//	http.HandleFunc("/111",handler)
//	http.ListenAndServe(":8888",nil)
//}
//
//func handler(writer http.ResponseWriter,request *http.Request){
//	fmt.Fprintf(writer,"hello World,%s!",request.URL.Path[1:])
//}
import (
	"fmt"
	"net/http"
)
type MyHandler struct {}
func (h *MyHandler)ServeHTTP(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"Hello World!")
}

type helloHandler struct {}
func (h *helloHandler)ServeHTTP(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"Hello")
}

type worldHandler struct {}
func (word  *worldHandler)ServeHTTP(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"World")
}

func helloWorld(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"Hello World!")
}
func main(){
	hello:=helloHandler{}
	world:=worldHandler{}
	//将Handler与DefaultServeMux（默认多路复用器）绑定
	http.Handle("/hello",&hello)//Handler接口需要包含一个ServeHTTP方法
	http.Handle("/world",&world)
	http.HandleFunc("/helloWorld",helloWorld)//可以直接把一个func函数设置成一个Handler
	http.ListenAndServe(":8080",nil)
}


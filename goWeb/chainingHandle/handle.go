package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

func main(){
	http.HandleFunc("/hello",log(hello))
	http.ListenAndServe(":8080",nil)
}

func hello(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"hello!")
}

func log(h http.HandlerFunc)http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {//将log和hello串联在一起了
		name:=runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("Handler function called - "+name)
		h(w,r)
	}
}
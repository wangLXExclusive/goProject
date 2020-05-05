package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func writerExample(w http.ResponseWriter,r *http.Request)  {
	str:=`<html>
<head><title>Go Web ProGramming</title></head>
<body><h1>Hello World</h1></body>
</html>`
	w.Write([]byte(str))
}
func writeHeaderExample(w http.ResponseWriter,r *http.Request)  {
	w.WriteHeader(501)
	fmt.Fprintln(w,"No such service,try next door!")
}
func headerExample(w http.ResponseWriter,r *http.Request)  {
	w.Header().Set("Location","http://baidu.com")  //重定向网址
	w.WriteHeader(302)
}

type Post struct {
	User string
	Threads []string
}

func jsonExample(w http.ResponseWriter,r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	post:=&Post{
		User:    "Sau sheong",
		Threads: []string{"first","second","third"},
	}
	json1,_:=json.Marshal(post)
	w.Write(json1)
}
func main() {
	http.HandleFunc("/write",writerExample)
	http.HandleFunc("/writeheader",writeHeaderExample)
	http.HandleFunc("/redirect",headerExample)
	http.HandleFunc("/json",jsonExample)
	http.ListenAndServe(":8080",nil)
}

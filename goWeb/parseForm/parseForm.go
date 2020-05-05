package main

import (
	"fmt"
	"net/http"
)

func process(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprintln(w,r.FormValue("hello"))
	r.ParseForm()		//   只支持application/x-www-form-urlencoded编码
	//fmt.Fprintln(w,r.Form)
	//fmt.Fprintln(w,r.Form["hello"])    //hello 同时出现在表单和URL两个地方的键 返回一个同时
										//包含了键的表单值和URL值的切片（并且表单值在切片中总在URL值前)
	fmt.Fprintln(w,r.PostForm["hello"]) //如果只想获得表单键值对，可以访问PostForm

	//fmt.Fprintln(w,r.Form["post"])
}

func main()  {
	http.HandleFunc("/process",process)
	http.ListenAndServe(":8080",nil)
}

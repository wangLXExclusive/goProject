//*********************************************************
//          			http测试
//
//              Date: 2020/4/25 14:14
//*********************************************************

package main

import "net/http"

func index(w http.ResponseWriter,r *http.Request)  {

}
func first(w http.ResponseWriter,r *http.Request)  {
	w.Write([]byte(`hello`))
}
func main()  {
	http.HandleFunc("/",index)

}
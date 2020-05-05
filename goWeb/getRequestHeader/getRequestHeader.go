package main

import (
	"fmt"
	"net/http"
)

func headerHandle(w http.ResponseWriter,r *http.Request){
	fmt.Fprintln(w,r.Header)
}
func acceptEncodingHandle(w http.ResponseWriter,r *http.Request){
	fmt.Fprintln(w,r.Header["Accept-Encoding"])
}
func ConnectionHandle(w http.ResponseWriter,r *http.Request){
	fmt.Fprintln(w,r.Header.Get("Connection"))
}
func main(){
	http.HandleFunc("/header",headerHandle)
	http.HandleFunc("/header/accept_encoding",acceptEncodingHandle)
	http.HandleFunc("/header/connection",ConnectionHandle)
	http.ListenAndServe(":8080",nil)
}
package main

import (
	"fmt"
	"log"
	"net/http"
)

func process(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprintln(w,r.FormValue("hello"))
	r.ParseMultipartForm(1024)
	fmt.Fprintln(w,r.MultipartForm)

}

func main() {
	http.HandleFunc("/process",process)
	err:=http.ListenAndServe(":8080",nil)
	if err!=nil{
		log.Fatal("ListenAndServer: ",err)
	}
}

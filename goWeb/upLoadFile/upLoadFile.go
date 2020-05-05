package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func process(w http.ResponseWriter,r *http.Request)  {
	r.ParseMultipartForm(1024)
	fileHeader:=r.MultipartForm.File["uploaded"][0]
	file,err:=fileHeader.Open()
	if err==nil{
		data,err:=ioutil.ReadAll(file)
		if err==nil{
			fmt.Fprintln(w,string(data))
		}
	}
}


func main() {
	http.HandleFunc("/process",process)
	http.ListenAndServe(":8080",nil)
}

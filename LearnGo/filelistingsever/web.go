package main

import (
	"LearnGo/filelistingsever/filelisting"
	"log"
	"net/http"
	"os"
)

type userErr interface {
	error
	Message()string
}
type appHandle func (writer http.ResponseWriter,request *http.Request)  error
func errWrapper(handler appHandle)func (writer http.ResponseWriter,request *http.Request){
	return func(writer http.ResponseWriter, request *http.Request) {
		//defer func() {
		//	if r:=recover();r!=nil {
		//		log.Printf("Panic:%v", r)
		//		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		//	}
		//}()
		err:=handler(writer,request)
		if err!=nil{
			log.Printf("Error occurred"+"handling request:%s",err.Error())
			if userErr,ok:=err.(userErr);ok{
				http.Error(writer,userErr.Message(),http.StatusBadRequest)
				return
			}
			code:=http.StatusOK
			switch  {
			case os.IsNotExist(err):
				code=http.StatusNotFound
			default:
				code=http.StatusInternalServerError
			}
			http.Error(writer,http.StatusText(code),code)
		}
	}
}
func main(){
	http.HandleFunc("/",errWrapper(filelisting.HandleFileList))
	err:=http.ListenAndServe(":8888",nil)
	if err!=nil{
		panic(err)
	}
}

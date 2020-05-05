package filelisting

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type userErr string
func (e userErr) Error()string{
	return  e.Message()
}
func (e userErr)Message()string{
	return string(e)
}
func HandleFileList(writer http.ResponseWriter,request *http.Request)error{
	if strings.Index(request.URL.Path, "/list/")!=0{
		return userErr("path must start with "+"/list/")
	}
	path:=request.URL.Path[len("/list/"):]
	file, err:=os.Open(path)
	if err!=nil{
		return err
	}
	defer file.Close()
	all, err:=ioutil.ReadAll(file)
	if err!=nil{
		return  err
	}
	writer.Write(all)
	return  nil
}
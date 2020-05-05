package main

import (
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/list/web.go",
		func(writer http.ResponseWriter, requst *http.Request) {
			path := requst.URL.Path[len("/list/"):]
			file, err := os.Open(path)
			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}
			defer file.Close()
			all, err := ioutil.ReadAll(file)
			if err != nil {
				panic(err)
			}
			writer.Write(all)
		})
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}

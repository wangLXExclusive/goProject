package main

import (
	"fmt"
)

func tryRecover (){
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Error occurred", err)
		} else {
			panic(r)
			//fmt.Sprint("I do not know")
		}
	}()
	b:=0
	a:=5/b
	fmt.Println(a)
}
func main(){
	tryRecover()
}
package main

import (
	"fmt"
	"math/rand"
	"time"
)
func generator()chan int{
	out:=make(chan int)
	i:=0
	go func() {
		for{
			time.Sleep(time.Duration(rand.Intn(1500))*time.Millisecond)
			out<-i
			i++
		}
	}()
	return out
}
func main() {
	var c1,c2 =generator(),generator()
	for {
		select {
		case n := <-c1:
			fmt.Println("Recived from c1", n)
		case n := <-c2:
			fmt.Println("Recived from c2", n)
		//default:
		//	fmt.Println("No vaule recived")
		}
	}
}

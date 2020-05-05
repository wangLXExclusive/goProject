package main

import (
	"fmt"
	"time"
)

func main() {
	chan1,chan2:=make(chan int,10),make(chan int,10)
	go addNumberToChan(chan1)
	go addNumberToChan(chan2)

	for{
		select {
		case e:=<-chan1:
			fmt.Println("From chan1:",e)
		case e:=<-chan2:
			fmt.Println("From chan2:",e)
		default:
			fmt.Println("No element!")
			time.Sleep(1*time.Second)
		}
	}
}
func readChan(chanName <-chan int) int {
	return <-chanName
}
func writeChan(chanName chan<- int,val int){
	chanName<-val
}
func addNumberToChan(chanName chan int){
	for{
		writeChan(chanName,1)
		time.Sleep(1*time.Second)
	}
}
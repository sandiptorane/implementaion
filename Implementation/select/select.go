package main

import (
	"fmt"
	"time"
)

func process(ch chan string){
	time.Sleep(10500*time.Millisecond)
	ch <-"Process successful"
}

func main(){
	ch := make(chan string)
	go process(ch)
	for {
		time.Sleep(1000*time.Millisecond)
		select {
		case value := <-ch:
			fmt.Println("value received ", value)
			return
		default:
			fmt.Println("value not received")
		}
	}
}

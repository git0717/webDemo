package main

var counter int =0
var ch = make(chan int ,1)

func Worker()  {
	ch<-1
	counter++
	<-ch
}
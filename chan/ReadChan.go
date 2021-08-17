package _chan

import "fmt"

func readChan(chanName <-chan int){
	<-chanName
}

func writeChan(chanName chan<-int)  {
	chanName<-1
}

func main()  {
	var  mychan = make(chan int,10)
	writeChan(mychan)
	readChan(mychan)
}

func chanRange(chanName chan int){
	for e:=range chanName{
		fmt.Println("get element from chan:d%\n",e)
	}
}

func MapCURD(){
	m :=make(map[string]string,10)
	m["apple"] = "red"
	m["apple"] = "green"
	delete(m,"apple")
	v,exist :=m["apple"]
	if exist{
		fmt.Println("apple-s%\n",v)
	}
}
package main

import(
	"time"
	"fmt"
)



func p(j int){
	var temp int
	for i:=0;i<10;i++{
		temp=n
		time.Sleep(10*time.Duration(i%5)*time.Nanosecond)
		n=temp+1
		fmt.Printf("Soy el proceso %d\n",j)
		time.Sleep(10*time.Duration(i%5)*time.Nanosecond)
	}
	

}

var n int

func main(){
	n=0
	go p(1)
	go p(2)

	time.Sleep(time.Second)
	fmt.Println(n)

}
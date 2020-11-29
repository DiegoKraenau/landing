package main

import (
	"fmt"
)

var turn int=1

//Exclusion mutua

func p(){
	for{
		fmt.Println("P,SNC")
		for turn!=1{
			//Esperar
		}
		fmt.Println("P,Section Crytical")
		turn=2
	}
} 

func q(){
	for{
		fmt.Println("Q,SNC")
		for turn!=2{
			//Esperar
		}
		fmt.Println("Q,Section Crytical")
		turn=1
	}
}

func main(){
	go p()
	q()
}
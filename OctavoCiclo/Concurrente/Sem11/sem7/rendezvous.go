package main

import(
	"fmt"
	"sync"
)

func Process( nombre string,procede sync.Mutex){
	if(nombre=="a1"){
		fmt.Println(nombre," Esperando a B2!!")
		procede.Lock()
	}else if(nombre=="a2"){
		procede.Lock()
		fmt.Println(nombre," Esperando a B1!!")
	}else if(nombre=="b1"){
		procede.Lock()
		fmt.Println(nombre," Esperando a B1!!")
	}
}



func main(){
	
}
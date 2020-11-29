package main
/*

import("fmt")

func ProcesoOrdena(id int,canalIn,canalOut chan int){
	min:=1000000 //valor min inicial
	for num:=range canalIn{
		if(num<min){
			canalOut<-min
			min=num
		}else{
			canalOut<-num
		}
	}
	
	fmt.Printf("Proceso %d:  %d\n",id,min)
	close(canalOut)
}


func main(){

	numeros:=[]int {4,3,5,7,8,2,9,1,10,6}
	n:=len(numeros)
	ch:=make([]chan int,n+1)

	//canalas sincronos
	for i:=range ch{
		ch[i]=make(chan int)
	}

	//lanzar los procesos que ordenan la lista de numeros
	for i:=range numeros{
		go ProcesoOrdena(i,ch[i],ch[i+1])
	}

	//proceso que envia los numero a ordenar
	go func() {
		for _,valor:= range numeros{
			ch[0]<-valor
		}
		close(ch[0])
	}()

	for range ch[n]{
		
	}
}
*/
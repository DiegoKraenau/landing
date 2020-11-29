package main
/*
import("fmt")

func generarMensajes(ch chan string,texto string){
	for{
		
	ch<-texto
	}
}

func main(){
	ch:=make(chan string)//Sincrono(Solo permite 1 flujo)

	//generar procesos
	go generarMensajes(ch,"Hola")
	go generarMensajes(ch,"Mundo")
	go generarMensajes(ch,"Chau")

	for {
		fmt.Println(<-ch)
	}
}
*/
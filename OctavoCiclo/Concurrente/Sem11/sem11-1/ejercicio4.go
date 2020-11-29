package main
/*

import("fmt")

func filosofo(nombre string,tenedorDer,tenedorIzq chan bool){
	for{
		fmt.Printf("%s esta pensando\n",nombre)
		//sincronización
		<-tenedorDer//Recepciona
		<-tenedorIzq
		fmt.Printf("%s esta comiendo\n",nombre)
		tenedorDer<-true//Envia
		tenedorIzq<-true
	}
	
}

func inicio(tenedorIni chan bool){
	for{
		tenedorIni<-true
		<-tenedorIni	
	}
}	


func main(){

	n:=5
	tenedor:=make([]chan bool,n)
	nombres:=[]string{"Platon","Aristoteles","Socrates","Descartes"}
	//Inicializar los canales
	for i:=range tenedor{
		tenedor[i]=make(chan bool,1)
	
	}
	//Generar los procesos de los filosofos
	for i:=0;i<n-1;i++{
		go filosofo(nombres[i],tenedor[i],tenedor[i+1])
		go inicio(tenedor[i])
	}

	//Ultimo filosofo
	go inicio(tenedor[n-1])
	filosofo("Tales de Mileto",tenedor[n-1],tenedor[0])

}
*/
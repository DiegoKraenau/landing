
package main
/*
import(
	"fmt"
	"sync"
)

func filosofo(nombre string,izquierda,derecha sync.Mutex){
	for{
		fmt.Println(nombre," Pensando!!")
		izquierda.Lock()
		derecha.Lock()
		fmt.Println(nombre,"Comiendo!!")
		izquierda.Unlock()
		derecha.Unlock()
	}
}

func main(){ 
	fork:=make([] sync.Mutex,5)
	go filosofo("Aristoteles",fork[0],fork[1])
	go filosofo("Platon",fork[1],fork[2])
	go filosofo("Socrates",fork[2],fork[3])
	go filosofo("Pitagoras",fork[3],fork[4])
	filosofo("Tales de Mileto",fork[4],fork[0])
}
*/
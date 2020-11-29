package main
import(
	"net"
	"fmt"
)


func main(){

	// establecer la conexion con el marter
	con,_:=net.Dial("tcp","localhost:8000")
	defer con.Close()

	//Que enviamos??
	fmt.Fprintln(con,"Nos compunicamos desde el cliente APOLO!!!")


}
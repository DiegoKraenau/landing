package main 

import(
	"fmt"
	"net"
	"os"
	"bufio"
)

func main(){
	//Rol de aplicacion servidor
	//escuchar por el puerto 8000

	ln,err:=net.Listen("tcp","localhost:8000")

	if err!=nil{
		fmt.Println("Falla al momento de establecer la comunicacion por el puerto 8000")
		os.Exit(1)//finalizo con error
	}

	defer ln.Close()
	con,err:=ln.Accept()

	if err!=nil{
		fmt.Println("Falla al momento de aceptar cliente",err.Error)
	}

	defer con.Close()
	// leer los datos de los clientes que son aceptados
	r:=bufio.NewReader(con)
	mensaje,_:=r.ReadString('\n')
	fmt.Println(mensaje)

}
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

var remotehost string

func main() {
	//host remoto y puerto
	rIn := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese el puerto de destino: ")
	port, _ := rIn.ReadString('\n')
	port = strings.TrimSpace(port)
	remotehost = fmt.Sprintf("localhost:%s", port)
	enviar(54)


}

func enviar(num int) {
	conn, _ := net.Dial("tcp", remotehost)
	defer conn.Close()
	fmt.Fprintf(conn, "%d\n", num)
}

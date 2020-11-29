package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"math"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
)

var (
	mutex sync.Mutex
)

type LinearRegression struct {
	x, y []float64
	xTotal, yTotal, xSqTotal, xYTotal, numerator, denominator,
	slope, yInt float64
}

type Ecuacion struct {
	M, B float64
}

func (e *Ecuacion) ReturnEcuacionParts(c chan []float64,num int) {

	var l LinearRegression
	l.LinearRegressionInit(c)
	e.M = l.slope
	e.B = l.yInt
	sayEcuation(e.M, e.B,num)

}

func (l *LinearRegression) LinearRegressionInit(c chan []float64) {

	l.x = <-c
	l.y = <-c
	//fmt.Println(len(l.X))
	for i, _ := range l.x {
		l.xTotal = l.xTotal + l.x[i]
		l.yTotal = l.yTotal + l.y[i]
		l.xSqTotal = l.xSqTotal + math.Pow(l.x[i], 2)
		l.xYTotal = l.xYTotal + (l.x[i] * l.y[i])
	}

	l.numerator = (float64(len(l.x)) * l.xYTotal) - (l.xTotal * l.yTotal)
	l.denominator = (float64(len(l.x)) * l.xSqTotal) - (math.Pow(l.xTotal, 2))
	yIntNumerator := (l.yTotal * l.xSqTotal) - (l.xTotal * l.xYTotal)
	l.yInt = yIntNumerator / l.denominator
	l.slope = l.numerator / l.denominator
	fmt.Println("Datos para MathScore:", l.x)
	fmt.Println("Datos para WritingScore", l.y)

}

func leerCSV(c chan []float64) {
	mutex.Lock()
	fileUrl := "https://github.com/DiegoKraenau/TA2Concurrente/raw/main/StudentsPerformance.csv"
	err := DescargarArchivo("StudentsPerformance.csv", fileUrl)
	if err != nil {
		panic(err)
	}

	path := flag.String("f", "./StudentsPerformance.csv", "CSV filepath")
	dataFile, err := os.Open(*path)

	reader := csv.NewReader(dataFile)

	var x []float64
	var y []float64

	for {
		record, err := reader.Read()

		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("Error al leer el csv: %v", err)

		}

		// convertir todas las filas a float64
		for i, val := range record {
			float, err := strconv.ParseFloat(val, 64)

			if err != nil {
				fmt.Printf("Error al convertir el valor <%v> afloat64: %v", val, err)
				break
			}

			if i < len(record)-1 {
				x = append(x, float)
			} else {
				y = append(y, float)
			}
		}

	}
	c <- x
	c <- y

}

func sayEcuation(x float64, y float64,num int) {

	fmt.Println("Ecuacion de probabilidades generada: ", `y = `, x, `x + `, y)
	fmt.Println("Donde X es la variable MathScore")
	fmt.Println("Donde Y es la variable WritingScore")
	var numconver float64 = float64(num) 
	fmt.Println("El posible resultado del WritingScore es: ",numconver*x+y)

}

func DescargarArchivo(filepath string, url string) error {

	// Obtiene la data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Crea el archivo
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Lo escribe al archivo
	_, err = io.Copy(out, resp.Body)
	return err
}



var remotehost string
var n int
var chContador chan int

func main() {


	

	//host y puerto del nodo local
	rIn := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese Puerto: ")
	port, _ := rIn.ReadString('\n')
	port = strings.TrimSpace(port)
	hostname := fmt.Sprintf("localhost:%s", port)

	//host y puerto del nodo remoto
	fmt.Print("Ingrese el puerto remoto: ")
	port, _ = rIn.ReadString('\n')
	port = strings.TrimSpace(port)
	remotehost = fmt.Sprintf("localhost:%s", port)

	//Aqui
	fmt.Print("Cantidad de numero a recibir: ")
	cant, _ := rIn.ReadString('\n')
	cant = strings.TrimSpace(cant)

	n, _ = strconv.Atoi(cant)

	//sincronizacion
	chContador = make(chan int, 1)
	chContador <- 0

	//escuchar
	ln, _ := net.Listen("tcp", hostname)
	defer ln.Close()
	for {
		conn, _ := ln.Accept()
		go manejador(conn)
	}

}

func manejador(conn net.Conn) {
	defer conn.Close()
	rIn := bufio.NewReader(conn)
	dato, _ := rIn.ReadString('\n')
	num, _ := strconv.Atoi(strings.TrimSpace(dato))

	fmt.Printf("Llegó el número %d\n", num)

		
	ch := make(chan []float64, 2)

	leerCSV(ch)

	var Ecuacion Ecuacion
	Ecuacion.ReturnEcuacionParts(ch,num)

	var input string
	fmt.Scanln(&input)
	fmt.Println("Fin ...")

	cont := <-chContador
	//lógica del ordenamiento

	enviar(num)

	cont++
	if cont == n {
		fmt.Printf("El numero menor es: %d\n", num)
		cont = 0
		
	}
	chContador <- cont
}
func enviar(num int) {
	conn, _ := net.Dial("tcp", remotehost)
	defer conn.Close()
	fmt.Fprintf(conn, "%d\n", num)
}

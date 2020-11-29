package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

const clusters = 3
const threshold = 0.5

var mutex []sync.Mutex

type Iris struct {
	ID          int     `json:"id"`
	SepalLength float32 `json:"sepal_length"`
	SepalWidth  float32 `json:"sepal_width"`
	PetalLength float32 `json:"petal_length"`
	PetalWidth  float32 `json:"petal_width"`
	Specie      string  `json:"specie"`
}

type Kluster struct {
	X    float32
	Y    float32
	Iris []Iris
}

func loadDataSet() []Iris {
	irisData, _ := os.Open("./data/iris.csv")
	reader := csv.NewReader(irisData)
	var irises []Iris
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal("Error al abrir el archivo: ", err)
		}

		id, _ := strconv.Atoi(row[0])
		sepalLength, _ := strconv.ParseFloat(row[1], 32)
		sepalWidth, _ := strconv.ParseFloat(row[2], 32)
		petalLength, _ := strconv.ParseFloat(row[3], 32)
		petalWidth, _ := strconv.ParseFloat(row[4], 32)

		irises = append(irises, Iris{
			ID:          id,
			SepalLength: float32(sepalLength),
			SepalWidth:  float32(sepalWidth),
			PetalLength: float32(petalLength),
			PetalWidth:  float32(petalWidth),
			Specie:      row[5],
		})
	}
	//fmt.Println(irises)
	return irises
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func calcularDistancia(kluster *Kluster, x float32, y float32) float32 {
	return float32(math.Sqrt(float64((kluster.X*kluster.X - x*x) + (kluster.Y*kluster.Y - y*y))))
}

func promediarDistanciasDePetalos(kluster *Kluster) {
	var x, y float32
	long := len(kluster.Iris)
	for i := range kluster.Iris {
		x += kluster.Iris[i].PetalLength
		y += kluster.Iris[i].PetalLength
	}
	kluster.X = x / float32(long)
	kluster.Y = y / float32(long)
}

func promediarDistanciasDeSepalos(kluster *Kluster) {
	var x, y float32
	long := len(kluster.Iris)
	for i := range kluster.Iris {
		x += kluster.Iris[i].SepalLength
		y += kluster.Iris[i].SepalWidth
	}
	kluster.X = x / float32(long)
	kluster.Y = y / float32(long)
}

func sepalClustering() Kluster {
	data := loadDataSet()
	pasos := 1
	kluster := Kluster{
		1.0 + rand.Float32() + float32(rand.Intn(7)),
		1.0 + rand.Float32() + float32(rand.Intn(7)),
		make([]Iris, len(data)),
	}

	for {
		fmt.Printf("Dimensiones del sepalo del cluster (largo x ancho): %f cm x %f cm\n", kluster.X, kluster.Y)
		fmt.Println("Paso: ", pasos)
		for i := range data {
			kluster.Iris = append(kluster.Iris, data[i])
		}
		xTemp := kluster.X
		yTemp := kluster.Y
		promediarDistanciasDeSepalos(&kluster)
		if calcularDistancia(&kluster, xTemp, yTemp) <= threshold {
			break
		}
		pasos++
	}

	return kluster
}

func petalClustering() Kluster {
	data := loadDataSet()
	pasos := 1
	kluster := Kluster{
		1.0 + rand.Float32() + float32(rand.Intn(7)),
		1.0 + rand.Float32() + float32(rand.Intn(7)),
		make([]Iris, len(data)),
	}

	for {
		fmt.Printf("Dimensiones del sepalo del cluster (largo x ancho): %f cm x %f cm\n", kluster.X, kluster.Y)
		fmt.Println("Paso: ", pasos)
		for i := range data {
			kluster.Iris = append(kluster.Iris, data[i])
		}
		xTemp := kluster.X
		yTemp := kluster.Y
		promediarDistanciasDePetalos(&kluster)
		if calcularDistancia(&kluster, xTemp, yTemp) <= threshold {
			break
		}
		pasos++
	}

	return kluster
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "GO Rest API")
}

func getAllIris(w http.ResponseWriter, r *http.Request) {
	data := loadDataSet()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	const port = ":5000"
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", index)
	router.HandleFunc("/iris", getAllIris)

	petalClustering()

	fmt.Println("Servidor en localhost" + port)
	log.Fatal(http.ListenAndServe(port, router))
}

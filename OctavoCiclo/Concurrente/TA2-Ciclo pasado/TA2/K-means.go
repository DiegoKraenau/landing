package main 
 
import (
	"fmt"
	"os"
	"encoding/csv"
	"io"
	"strconv"
	"math/rand"
)

type Persona struct {
	Nombre string
	Anios int
	Ganan_Mensual int
	Gastan_Mensual int
	Num_tarjetas int
	Deudas int
}

var personas []Persona
 
func leerData() {
    f, err := os.Open("Dataset.csv")
    if err != nil {
        panic(err)
	}
	
	defer f.Close()

	r := csv.NewReader(f)
	r.Comma = ';'
	r.Comment = '#'
	r.FieldsPerRecord = -1

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("Error leyendo informacion", err)
		}

		p := Persona {
			Nombre: record[0],
		}

		Anios, err1 := strconv.Atoi(record[1])
		Ganan_Mensual, err2 := strconv.Atoi(record[2])
		Gastan_Mensual, err3 := strconv.Atoi(record[3])
		Num_tarjetas, err4 := strconv.Atoi(record[4])
		Deudas, err5 := strconv.Atoi(record[5])
		if err1 != nil { fmt.Println(err1,err2,err3,err4,err5)}

		p.Anios = Anios
		p.Ganan_Mensual  = Ganan_Mensual
		p.Gastan_Mensual = Gastan_Mensual
		p.Num_tarjetas = Num_tarjetas
		p.Deudas = Deudas

		personas = append(personas, p)
	}
}

func main() {
	leerData()
	
	k_cluster := 3

	var cluster1 []int
	var cluster2 []int
	var cluster3 []int

	for i := 0; i < k_cluster; i++ {
		cluster1 = append(cluster1, rand.Intn(len(personas)))
		cluster2 = append(cluster2, rand.Intn(len(personas)))
		cluster3 = append(cluster3, rand.Intn(len(personas)))
	}

	var centroide1 []int
	//var centroide2 []int
	//var centroide3 []int


	fmt.Println(personas[5])
	fmt.Println(personas[6])
	fmt.Println(personas[7])

	Anios:=0
	Ganan_Mensual:=0
	Gastan_Mensual:=0
	Num_tarjetas:=0
	Deudas:=0

	for i:=0 ; i<len(cluster1); i++ {
		Anios += personas[cluster1[i]].Anios
		Ganan_Mensual += personas[cluster1[i]].Ganan_Mensual
		Gastan_Mensual += personas[cluster1[i]].Gastan_Mensual
		Num_tarjetas += personas[cluster1[i]].Num_tarjetas
		Deudas += personas[cluster1[i]].Deudas
	}

	centroide1 = append(centroide1, Anios/3)
	centroide1 = append(centroide1, Ganan_Mensual/3)
	centroide1 = append(centroide1, Gastan_Mensual/3)
	centroide1 = append(centroide1, Num_tarjetas/3)
	centroide1= append(centroide1, Deudas/3)

}
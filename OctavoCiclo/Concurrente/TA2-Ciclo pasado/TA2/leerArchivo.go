package main 
 
import (
	"fmt"
	"os"
	"encoding/csv"
	"io"
	"strconv"
)

type Persona struct {
	Nombre string
	Anios int
	Ganan_Mensual int
	Gastan_Mensual int
	Num_tarjetas int
	Deudas int
}
 
func main() {
    f, err := os.Open("Dataset.csv")
    if err != nil {
        panic(err)
	}
	
	defer f.Close()

	r := csv.NewReader(f)
	r.Comma = ';'
	r.Comment = '#'
	r.FieldsPerRecord = -1

	var personas []Persona

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
		if err1 != nil { fmt.Println("Error leyendo anios", err1) }
		Ganan_Mensual, err2 := strconv.Atoi(record[2])
		if err2 != nil { fmt.Println("Error leyendo ganan mensual", err2) }
		Gastan_Mensual, err3 := strconv.Atoi(record[3])
		if err3 != nil { fmt.Println("Error leyendo gastan mensual", err3) }
		Num_tarjetas, err4 := strconv.Atoi(record[4])
		if err4 != nil { fmt.Println("Error leyendo numero de tarjetas", err4) }
		Deudas, err5 := strconv.Atoi(record[5])
		if err5 != nil { fmt.Println("Error leyendo deudas", err5) }

		p.Anios = Anios
		p.Ganan_Mensual  = Ganan_Mensual
		p.Gastan_Mensual = Gastan_Mensual
		p.Num_tarjetas = Num_tarjetas
		p.Deudas = Deudas

		personas = append(personas, p)
		fmt.Println(record)
	}
}
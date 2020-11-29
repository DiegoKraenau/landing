package main

import (
	"encoding/json"
  	"log"
  	"net/http"
  	"github.com/gorilla/mux"
)

type Persona struct {
	ID string `json:"id,omitempty"`
  	Nombre string `json:"nombre,omitempty"`
  	Anios int `json:"anios,omitempty"`
  	Ganan_Mensual int `json:"ganan_mensual,omitempty"`
  	Gastan_Mensual int `json:"gastan_mensual,omitempty"`
  	Num_tarjetas int `json:"num_tarjetas,omitempty"`
	Deudas int `json:"deudas,omitempty"`
}

var personas []Persona

func GetPersonasEndpoint(w http.ResponseWriter, req *http.Request){
  	json.NewEncoder(w).Encode(personas)
}

func CreatePersonaEndpoint(w http.ResponseWriter, req *http.Request){
	params := mux.Vars(req)
	var persona Persona
	_ = json.NewDecoder(req.Body).Decode(&persona)
	persona.ID = params["id"]
	personas = append(personas, persona)
	json.NewEncoder(w).Encode(personas)
}


func main() {
	router := mux.NewRouter()

	personas = append(personas, Persona{ID: "1", Nombre:"Mario", Anios:20, Ganan_Mensual:3500, Gastan_Mensual:2800, Num_tarjetas:2, Deudas: 100})
	
	router.HandleFunc("/personas", GetPersonasEndpoint).Methods("GET")
	router.HandleFunc("/personas", CreatePersonaEndpoint).Methods("POST")

	log.Fatal(http.ListenAndServe(":3000", router))
}
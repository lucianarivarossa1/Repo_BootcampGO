package main

import (
	"fmt"
	//"encoding/json"
)

type Preferencias struct {
	Comidas string
}
type Persona struct {
	Nombre string `json: "nombre"` //informacion extra que le puede servir a alguien
	Edad   int    `json: "edad"`
	Gustos Preferencias
}

func (p Persona) consultar() string {
	return p.Nombre
}
func main() {

	p1 := Persona{
		Nombre: "luciana",
		Edad:   30,
		Gustos: Preferencias{Comidas: "milanesa"},
	}
	fmt.Println(p1.Nombre, p1.Edad, p1.Gustos.Comidas)
	fmt.Println("funcion: ", p1.consultar())
}

package main

import "fmt"

type Compania struct {
	Nombre string
	Puesto string
}
type Empleado struct {
	Nombre   string
	Apellido string
	Compania Compania
}

func (e Empleado) InformacionEmpleado() {
	fmt.Printf("Empleado: %s %s\n Puesto: %s\n Compania: %s \n", e.Nombre, e.Apellido, e.Compania.Puesto, e.Compania.Nombre)
}
func (c *Compania) CambiarPuesto(puesto string) { //*Compania trae la direccion de memoria de todo el struct Compania y lo guarda en c
	fmt.Println("se produce cambio", &c)
	c.Puesto = puesto
}
func main() {
	e := Empleado{
		Nombre:   "Juan",
		Apellido: "Lopez",
		Compania: Compania{
			Nombre: "IT",
			Puesto: "programador",
		},
	}
	e.InformacionEmpleado()
	e.Compania.CambiarPuesto("programador junior")
	e.InformacionEmpleado()
}

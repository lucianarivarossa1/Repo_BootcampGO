package main

import "fmt"

type Alumno struct {
	Nombre   string
	Apellido string
	DNI      string
	Fecha    string
}

func (a Alumno) DetalleAlumno() {
	fmt.Printf("Alumno: \n Nombre: %s\n Apellido: %s \n DNI: %s \n Fecha: %s \n", a.Nombre, a.Apellido, a.DNI, a.Fecha)
}
func (a *Alumno) modificarAlumno(nombre, apellido, dni, fecha string) {
	a.Nombre = nombre
	a.Apellido = apellido
	a.DNI = dni
	a.Fecha = fecha
}
func main() {
	a := Alumno{
		Nombre:   "Juan",
		Apellido: "Lopez",
		DNI:      "34343434",
		Fecha:    "4/4/4",
	}
	a.DetalleAlumno()
	a.modificarAlumno("Pedro", "gomez", "3423232", "3/3/3")
	a.DetalleAlumno()

}

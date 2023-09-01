package main

import "fmt"

func main() {
	var (
		edad, antiguedad int     = 25, 2
		empleado         bool    = true
		sueldo           float64 = 150000.0
	)

	switch {
	case edad <= 22:
		fmt.Println("usted es menor de 23 años")
	case !empleado: //no esta empleado
		fmt.Println("usted esta desempleado")
	case antiguedad < 1:
		fmt.Println("su antiguedad es menor a un año")
	case sueldo < 100000:
		fmt.Println("el sueldo no es suficiente")
	default:
		fmt.Println("puede recibir el prestamo")
	}

}

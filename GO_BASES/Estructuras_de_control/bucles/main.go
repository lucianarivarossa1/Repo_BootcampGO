package main

import "fmt"

func main() {
	meses := []string{"enero", "febrero", "junio", "agosto", "abril"}
	obtenerEstacion(meses)
}

func obtenerEstacion(meses []string) {
	for _, mes := range meses {
		if mes == "enero" || mes == "febrero" || mes == "marzo" {
			fmt.Println("verano")
		} else if mes == "abril" || mes == "mayo" || mes == "junio" {
			fmt.Println("otoño")
		} else if mes == "julio" || mes == "agosto" || mes == "septiembre" {
			fmt.Println("invierno")
		} else if mes == "octubre" || mes == "noviembre" || mes == "diciembre" {
			fmt.Println("primavera")
		} else {
			fmt.Println("no existe ese mes")
		}
	}
}

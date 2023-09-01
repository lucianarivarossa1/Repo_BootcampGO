package main

import "fmt"

func main() {
	fmt.Println("el promedio es: ", calcularPromedio(1, 2, 3, 4))
}
func calcularPromedio(notas ...int) int {
	var promedio int
	for i, nota := range notas {
		nota += nota
		promedio = nota / i
	}

	return promedio

}

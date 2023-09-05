package main

import (
	"fmt"
	"os"
)

func EstadoFinal() {
	fmt.Println("Ejecución finalizada")
}

func main() {
	defer EstadoFinal()
	_, err := os.Open("customers.txt")
	if err != nil {
		panic(err.Error())
		fmt.Println("Error")
	}
	fmt.Println("leido ok")
}

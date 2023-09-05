package main

import (
	"fmt"
	"os"
)

var file, err = os.Open("customers.txt")

func EstadoFinal() {
	fmt.Println("Ejecución finalizada")
}
func CrearArchivo() {
	os.Create("customers.txt")
	os.WriteFile("customers.txt", []byte("Luciana"), 0644)
}
func AbrirArchivo() {
	if err != nil {
		panic(err.Error())
		fmt.Println("Error")
	}

}
func main() {

	err1 := file.Close()
	if err1 != nil {
		panic(err1.Error())
		fmt.Println("error")
	} else {
		file.Close()

	}
	defer EstadoFinal()
	//	CrearArchivo()
	AbrirArchivo()

}

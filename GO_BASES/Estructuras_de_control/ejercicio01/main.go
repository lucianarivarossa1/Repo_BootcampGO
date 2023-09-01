package main

import (
	"fmt"
)

func main() {
	var phrase string = "Hola"
	length := len(phrase)
	fmt.Println("la cantidad de caracteres es: ", length)
	for _, letter := range phrase {
		//letter genera variable de tipo rune , codigo unicode
		fmt.Println("la letra es: ", string(letter))
	}

}

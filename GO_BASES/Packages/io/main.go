package main

import (
	"fmt"
	"strings"
)

func main() {
	//creo una cadena de texto: (ahora str es un Reader)

	str := strings.NewReader("hello world ")

	//Reader es una interface que tiene un metodo Read, lo que nos va a pedir es un slice de bytes
	// devuelve n y err (n es el ultimo byte que leyo y err es error en la lectura)

	//Primero tenemos que dejar en algun lado los bytes, creamos un buffer (slice de bytes):
	buffer := make([]byte, 1)

	//empezamos a leer:
	for {
		n, err := str.Read(buffer)

		if err != nil {
			switch err.Error() {
			case "EOF":
				println("EOF\n") //Esto significa que ya leyo el ultimo byte
			default:
				println(err.Error())
			}
			return

		}
		//si no hay error, muestro info en pantalla. n tiene cada uno de los bytes
		fmt.Printf("Mensaje: %s - byte %d\n", buffer[:n], n)
	}
}

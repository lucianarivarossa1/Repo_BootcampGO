package main

import "fmt"

func main() {
	fmt.Println("hello world")
	//hola mundo

	var username string = "Luciana"

	var username2 string
	username2 = "Luciana 2"

	var username3 = "Luciana3" //GO es fuertemente tipado

	username4 := "Luciana 4"

	var (
		username5 = "Luciana 5"
		age       = 30
	)

	var age2, age3 int = 30, 31

	fmt.Println("Hello", username)
	fmt.Println(username2)
	fmt.Println(username3)
	fmt.Println(username4)
	fmt.Println(username5)
	fmt.Println(age)
	fmt.Println(age2)
	fmt.Println(age3)

	//Tipos de datos basicos
	//string -> cadena de caracteres "hola mundo"
	//int -> numeros enteros
	//float32 -> numeros decimales de 32 bits
	//float64 -> numeros decimales de 64 bits -- ESTE SE USA
	//int32 -> lo mismo pero numeros enteros
	//int64 -> lo mismo pero numeros enteros
	//arrays, slices, etccc pero no entran esta clase
	//bool -> true o false

}

package main

import "fmt"

func main() {
	type Cliente struct {
		Legajo    int
		Nombre    string
		DNI       int
		Telefono  int
		Domicilio string
	}
	cliente1 := Cliente{1, "Juan", 123, 123456, "Calle 1"}
	cliente2 := Cliente{2, "Pedro", 122, 123456, "Calle 2"}
	cliente3 := Cliente{2, "Pedro", 122, 123456, "Calle 2"}
	var clientes = [3]Cliente{cliente1, cliente2, cliente3}

	for i := 0; i < len(clientes); i++ {
		for u := 0; u < len(clientes); u++ {
			fmt.Println(clientes[i], clientes[u])
			if clientes[i] == clientes[u] {

				fmt.Println("repetido")
			}
		}

	}

}

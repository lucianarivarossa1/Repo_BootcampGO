package main

import "fmt"

type Person struct {
	name    string
	age     int
	height  float64
	license bool
}

// podemos agregar un metodo a persona que se llama String donde vamos a devolver una cadena de texto
// muestra lo que queremos que se vea
func (p Person) String() string {

	return fmt.Sprintf("nombre: %s, age %d, height: %f, license %t", p.name, p.age, p.height, p.license)
}

func main() {
	p := Person{
		name:    "Luciana",
		age:     30,
		height:  1.60,
		license: true,
	}

	fmt.Printf("Memory address: %p ", &p)
	//como tengo el metodo String no hace falta que haga esto:
	//fmt.Printf("Person info: nombre: %s, age %d, height: %f, license %t", p.name, p.age, p.height, p.license)

	//hago directamente esto: (no hace falta poner p.String(), lo toma solo)
	fmt.Printf("Person info %s", p)

}

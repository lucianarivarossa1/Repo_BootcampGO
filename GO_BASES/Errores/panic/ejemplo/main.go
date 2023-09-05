package main

import "fmt"

func prueba1() {
	fmt.Println("1")
}
func prueba2() {
	fmt.Println("2")
}
func prueba3() {
	fmt.Println("3")
}
func main() {
	defer prueba3()
	defer prueba1()
	defer prueba2()

}

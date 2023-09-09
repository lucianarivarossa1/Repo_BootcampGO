package main

import (
	"io"
	"os"
)

func main() {
	f, err := os.Open("./files.go")
	if err != nil {
		println(err.Error())
		return
	}
	defer f.Close()

	//leer el archivo

	buf := make([]byte, 2)
	for {
		n, err := f.Read(buf)
		if err != nil {
			switch err.Error() {
			case "EOF":
				println("EOF")
			default:
				println(err.Error())
			}
			break
		}
		println("El archivo dice %s", string(buf[:n]))
	}

	// Vuelve a cero el cursor que va leyendo, sino al hacer el copy el cursor esta al final y no puede mostrar nada
	f.Seek(0, 0)

	//Escribir lo que tiene el archivo
	io.Copy(os.Stdout, f)
}

package main

import (
	"errors"
	"fmt"
)

type ErrInvalidSalary struct {
	MinimunSalary int
	ActualSalary  int
}

func (e ErrInvalidSalary) Error() string {
	return "el salario no alcanza el minimo"
}

var (
	ErrInvalidSalaryThnTenThousand = errors.New("el salario no puede ser menor a 10.000")
)

func main() {
	//Maneras de definir errores:
	//1) Creando con variable global
	//var(ErrInvalidSalary = errors.New("el salario no alcanza el minimo"))
	//2) la forma que esta hecho
	var (
		err    error
		salary = 100_000
	)
	if salary == 150_000 {
		err = ErrInvalidSalary{}
		fmt.Println(err.Error())
		return

	}
	if salary < 10000 {
		//		err = ErrInvalidSalaryThnTenThousand{}
		fmt.Println(err.Error())
		return

	}
	if errors.Is(err, ErrInvalidSalary{}) {
		fmt.Println("el salario no alcanza ")

	}
	//if errors.Is(err,ErrInvalidSalaryThnTenThousand{}){
	//fmt.Println("el salario no alcanza ")

	//}
	fmt.Println("debe pagar impuesto")
}

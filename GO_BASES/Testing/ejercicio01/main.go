package main

const (
	mas50  = 0.17
	mas150 = 0.27
)

func calcularImpuesto(salario float64) float64 {
	var impuesto float64
	if salario > 50000 {
		impuesto = salario*mas50 + salario
	} else if salario > 150000 {
		impuesto = salario*mas150 + salario
	} else {
		impuesto = salario
	}

	return impuesto

}
func main() {

	println(calcularImpuesto(150000))
}

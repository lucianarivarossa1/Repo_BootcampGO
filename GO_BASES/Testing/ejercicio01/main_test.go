package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcularImpuesto(t *testing.T) {

	//Arrange: todo lo que necesita la funcion previo a que sea llamada
	var salario_test = 1000.0

	//Act: el llamado a la funcion
	resultado := calcularImpuesto(salario_test)

	//Assert: la comparacion entre lo que esperabamos y lo que sucedio
	assert.Equal(t, salario_test, resultado)
}

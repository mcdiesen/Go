package main

import (
	"bufio"
	"fmt"
	"os"
)

var numero1, numero2 int
var leyenda string
var resultado int

func main() {

	fmt.Println("Ingrese numero 1")
	fmt.Scanln(&numero1)

	fmt.Println("Ingrese numero 2")
	fmt.Scanln(&numero2)

	fmt.Println("Descripción")

	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		leyenda = scanner.Text()
	}
	resultado = numero1 + numero2

	//scanner := bufio.Writer(os.Stdout)
	fmt.Println("El resultado es ", resultado)

}

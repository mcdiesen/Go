//condicionales
package main

import "fmt"

func main() {

	var num1, num2 int

	println("ingrese el primer valor")
	fmt.Scanf("%d", &num1)

	println("Ingrese el segundo valor ")
	fmt.Scanf("%d", &num2)

	if num1 > num2 {
		fmt.Println("mayor", num1)
	} else {
		fmt.Println("menor", num2)
	}
}

package main

import "fmt"

func main() {
	fmt.Println("Mesma")
	fmt.Print("Linha.")

	x := 7.171313

	fmt.Println("O valor de x é ", x)

	fmt.Printf("Significa Print Formatado, o valor de x é %.2f", x)

	a := 1
	b := 1.9999
	c := false
	d := "opa"

	//Nota-se que alguns número são arredondados automaticamente pelo sistema!
	fmt.Printf("\n%d %f %.2f %t %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %v %v", a, b, c, d)

}

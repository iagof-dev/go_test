package main

import (
	"fmt"
	"os"
	"os/exec"
)

func sum(n1 int, n2 int) int {
	var result int = n1 + n2
	return (result)
}
func sub(n1 int, n2 int) int {
	var result int = n1 - n2
	return (result)
}

func mult(n1 int, n2 int) int {
	var result int = n1 * n2
	return (result)
}

func div(n1 int, n2 int) int {
	var result int = n1 / n2
	return (result)
}

func main() {
	clear := exec.Command("cmd", "/c", "cls")
	clear.Stdout = os.Stdout

	var nome string
	fmt.Println("Digite seu nome:")
	fmt.Scanln(&nome)

	for i := 0; i <= 99999999999; i++ {
		clear.Stdout = os.Stdout
		clear.Run()

		fmt.Println("Olá,", nome, "!")
		fmt.Println("Escolha uma opção:\n[1] Somar\n[2] Subtrair\n[3] Multiplicar\n[4] Dividir")
		var opcao int
		fmt.Scanln(&opcao)

		clear.Stdout = os.Stdout
		clear.Run()

		switch opcao {
		case 1:
			fmt.Print("Digite um numero:")
			var n1 int
			fmt.Scanln(&n1)

			fmt.Print("Digite mais um numero:")
			var n2 int
			fmt.Scanln(&n2)

			fmt.Println("A soma de", n1, "e", n2, "é:", sum(n1, n2))

		case 2:
			fmt.Print("Digite um numero:")
			var n1 int
			fmt.Scanln(&n1)

			fmt.Print("Digite mais um numero:")
			var n2 int
			fmt.Scanln(&n2)

			fmt.Println("A soma de", n1, "e", n2, "é:", sub(n1, n2))
		case 3:
			fmt.Print("Digite um numero:")
			var n1 int
			fmt.Scanln(&n1)

			fmt.Print("Digite mais um numero:")
			var n2 int
			fmt.Scanln(&n2)

			fmt.Println("A soma de", n1, "e", n2, "é:", mult(n1, n2))
		case 4:
			fmt.Print("Digite um numero:")
			var n1 int
			fmt.Scanln(&n1)

			fmt.Print("Digite mais um numero:")
			var n2 int
			fmt.Scanln(&n2)

			fmt.Println("A soma de", n1, "e", n2, "é:", div(n1, n2))
		default:
			fmt.Println("Opção inválida!")

		}

	}
}

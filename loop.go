package main

import "fmt"

func main() {
    fmt.Println("O único loop existente em go é o for")
	fmt.Println("\nDeclaração básica: ")
	var valor int
	var i int
	for i <= 3 {
        valor = valor + i
		i++
    }
	fmt.Println(valor)

	fmt.Println("\nDeclaração clássica: ")
	valor = 0
	for j := 0; j <= 3; j++ {
        valor = valor + j
    }
	fmt.Println(valor)

	fmt.Println("\nÉ também possível utilizar um for sem uma condição de pausa, o que o faria ser executado infinitamente a menos que seja utilizado um break ou return. (praticamente tem o comportamento de um while)")
	valor = 0
	for {
		valor ++
		if valor == 6 {
			break
		}
    }

	nomes := []string{"joão", "maria", "jose", "luders", "carlos", "carla"}

	fmt.Println("\nTemos como imprimir todos os elementos de um vetor utilizando o 'range'")
	for i, nome := range nomes {
		fmt.Println(i, nome)
	}
}


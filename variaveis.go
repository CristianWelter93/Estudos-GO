package main

import (
	"fmt"
	"reflect"
)


func main() {
	fmt.Println("O reflect foi utilizado para pegar o tipo da variável")
	var teste = true
    fmt.Println("Tem que utilizar a palavra var, se um tipo não for fornecido mas esta variável for inicializada, seu tipo será deduzido pelo Go: ")
	fmt.Println(teste)
	fmt.Println(reflect.TypeOf(teste))

	var numero1, numero2 int
	var numero3 int = 3
	fmt.Println("Variáveis numericas sem inicialização terão o valor 0")
	fmt.Println(numero1, numero2, numero3)
	fmt.Println(reflect.TypeOf(numero1))

	fmt.Println("O := é para declaração e atribuição de uma só vez, então não é necessário utilizar o var")
	texto := "String criada sem utilizar o var"
	fmt.Println(texto)

	fmt.Println("Uma constante pode ser declarada em qualquer lugar em que uma var poderia, basta utilizar const")
	fmt.Println(`
	Uma constante numerica não tem um tipo até que seja explicitamente convertida ou que utilize uma função que exija um tipo específico.
	Uma constante aritmética utiliza precisão infinita, ou seja, é limitada apenas pela memoria disponível no host
	`)
	const numero = 900000000000
	fmt.Println(numero)
	fmt.Println(float64(numero))
}


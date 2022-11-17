package main

import "fmt"

type Carro struct {
	Marca string
}

func (c *Carro) MudaMarca (marca string) {
	c.Marca = marca
	fmt.Println(c.Marca)
}

func main() {
	carro := Carro{Marca: "Fiat"}
	carro.MudaMarca("Ford")
	fmt.Println(carro.Marca)
}

func simpleExplanation() {
	/*Ao criar a variável "a" é separado um espaço na memória, guarda o valor 1
	e a variável aponta para o endereço dessa memória*/
	a := 1
	fmt.Println(a)
	/*Ao utilizar o & vai trazer o endereço da memória onde a variável a 
	está apontado, esse lugar na memória tem o valor de um e o "a" aponta
	para esse endereço */
	fmt.Println(&a)
	/*Veja que interessante, ao criar a variável "b" eu estou passando o 
	MESMO endereço de memória do "a" para o "b" e ao atribuir o valor 10
	ao "pointer" do b, ou seja, o "*" é o apontamento e estou apontando
	que o endereço do b é o mesmo do a. Por isso, ao atribuir 10 ao b
	ele afetará o a por ser o mesmo endereço*/
	b := &a
	*b = 10
	fmt.Println(b)
	fmt.Println(a) //novo valor será 10
}
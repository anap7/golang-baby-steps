package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 100; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)//d = i e s = name
		time.Sleep(time.Second)
	}
}
//Thread 01 -> Principal
func main() {
	//-------------------CANAIS-------------------//
	/*Um canal onde guardarei uma string, os canais no go servem para
	que as threads consigam se comunicar e utilizar a mesma variável */
	canal := make(chan string)

	//Thread 02
	go func() {
		canal <- "Veio da Thread 02"
	}()

	//Thread 01
	/*Estou jogando a variável da thread 02 para a thread 01*/
	msg := <-canal
	fmt.Println(msg)

	//-------------------CRIANDO THREADS-------------------//
	/*Ao inserir um "go" na frente da função significa que será 
	criada uma thread e os processos serão feitos em revezamento.
	a task "C" já faz parte de uma thread que é a própria main*/
	go task("A") //Isso é uma go routines -> Green threads
	go task("B")
	task("C")
}

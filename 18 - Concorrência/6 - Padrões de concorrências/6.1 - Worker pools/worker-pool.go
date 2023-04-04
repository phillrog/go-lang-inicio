package main

import "fmt"

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return fibonacci(n-2) + fibonacci(n-1)
}

func worker(tarefas <-chan int, resultados chan<- int) {
	for tarefa := range tarefas {

		resultados <- fibonacci(tarefa)
	}
}

func main() {
	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)

	posicao := int(45)

	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	for i := 0; i < posicao; i++ {
		tarefas <- i
	}
	close(tarefas)

	for i := 0; i < posicao; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}
}

package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	canal := make(chan string)
	go escrever("Olá Mundo", canal)

	// for {
	// 	mensagem, aberto := <-canal
	// 	if !aberto {
	// 		break
	// 	}
	// 	fmt.Println(mensagem)
	// }

	for mensagem := range canal {
		fmt.Println(mensagem)
	}
	fmt.Println("Fim do programa")
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto + " - " + strconv.Itoa(i)
		time.Sleep(time.Second)
	}

	close(canal)
}

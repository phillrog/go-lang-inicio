package main

import "fmt"

func main() {
	fmt.Println("maps")

	u1 := map[string]string{
		"nome":      "teste",
		"sobrenome": "test2",
	}

	fmt.Println(u1["nome"])

	u2 := map[string]map[string]string{
		"curso": {
			"nome":      "teste",
			"professor": "test2",
		},
	}

	fmt.Println(u2)
}

package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Segunda-feira"
	case 2:
		return "Terça-feira"
	case 3:
		return "Quarta-feira"
	case 4:
		return "Quinta-feira"
	case 5:
		return "Sexta-feira"
	case 6:
		return "Sábado"
	case 7:
		return "Domingo"
	default:
		return "Inválido"
	}

}

func diaDaSemana2(numero int) string {
	var diaDaSemana string
	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
		fallthrough
	case numero == 2:
		diaDaSemana = "Segunda-feira"
	case numero == 3:
		diaDaSemana = "Terça-feira"
	case numero == 4:
		diaDaSemana = "Quarta-feira"
	case numero == 5:
		diaDaSemana = "Quinta-feira"
	case numero == 6:
		diaDaSemana = "Sexta-feira"
	case numero == 7:
		diaDaSemana = "Sábado"
	case numero == 8:
		diaDaSemana = "Domingo"
	default:
		diaDaSemana = "Inválido"
	}

	return diaDaSemana
}

func main() {
	fmt.Println("Switch")

	dia := diaDaSemana(1)

	fmt.Println(dia)

	dia2 := diaDaSemana2(7)

	fmt.Println(dia2)

	dia3 := diaDaSemana2(1)

	fmt.Println(dia3)
}

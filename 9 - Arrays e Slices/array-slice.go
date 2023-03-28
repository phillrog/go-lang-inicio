package main

import "fmt"

func main() {
	fmt.Println("Arrays e Slices")

	var a1 [5]int
	a1[0] = 1
	fmt.Print(a1)

	a2 := [5]string{"1", "2", "3", "4"}
	fmt.Print(a2)

	a3 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(a3)

	s1 := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(s1)

	s1 = append(s1, 30)

	fmt.Println(s1)

	s3 := s1[1:3]

	fmt.Println(s3)

	fmt.Println(s3[1])

	s4 := make([]float32, 10, 15)
	fmt.Println(s4)

	fmt.Println(len(s4))
	fmt.Println(cap(s4))
}

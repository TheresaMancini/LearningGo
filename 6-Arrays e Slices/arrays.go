package main

import "fmt"

func main() {
	fmt.Println("Array e Slice")

	//Array é uma lista de valores
	var array1 [5]int

	array1[0] = 1

	fmt.Println(array1)

	// outra forma de declarar

	array2 := [5]string{"P1", "P2", "P3", "P4", "P5"}

	fmt.Println(array2)

	// outra forma de declarar

	array3 := [...]string{"P1", "P2", "P3"} // tamanho baseado no numero de itens passados

	fmt.Println(array3)

	// Slice
	slice := []int{1, 2, 3, 4}

	fmt.Println(slice)

	slice = append(slice, 5)

	fmt.Println(slice)

	slice2 := array2[1:3] // primeiro indice é inclusivo e o segundo exclusivo

	fmt.Println(slice2)

	array2[1] = "PAlretada"

	fmt.Println(slice2)
	fmt.Println(array2)

	// Array interno
	fmt.Println("-----------------------")
	slice3 := make([]int, 3, 4) // funcao make aloca memoria

	fmt.Println(slice3)

	fmt.Println(len(slice3))

	fmt.Println(cap(slice3))

	slice3 = append(slice3, 5)

	fmt.Println(slice3)

	fmt.Println(len(slice3))

	fmt.Println(cap(slice3))

	slice3 = append(slice3, 6)

	fmt.Println(slice3)

	fmt.Println(len(slice3))

	fmt.Println(cap(slice3))

}

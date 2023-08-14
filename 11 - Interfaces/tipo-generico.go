package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	oi := "teste"
	inteiro := 10
	generica(oi)
	generica(inteiro)
}

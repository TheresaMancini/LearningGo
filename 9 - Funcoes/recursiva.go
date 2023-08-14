package main

import (
	"fmt"
	"os"
	"strconv"
)

func fibonacci(n int64) int64 {
	if n <= 1 {
		return n
	} else {
		return fibonacci(n-2) + fibonacci(n-1)
	}
}

func main() {
	fmt.Println("Recursiva")

	//arg := os.Args[1]
	arg := os.Args[1:]

	for _, numero := range arg {
		//fmt.Println(i, numero)

		n, _ := strconv.ParseInt(numero, 10, 0)

		fn := fibonacci(n)

		fmt.Println(fn)
	}

}

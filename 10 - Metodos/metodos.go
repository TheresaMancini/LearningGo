package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Dentro do método Salvando os dados do usaurio: %s ", u.nome)
}

func (u usuario) maiorIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Usuario 1", 20}
	fmt.Println(usuario1)
	maiorIdade := usuario1.maiorIdade()
	fmt.Println(maiorIdade)

	usuario2 := usuario{"Usuario 2", 16}
	fmt.Println(usuario2)
	maiorIdade = usuario2.maiorIdade()
	fmt.Println(maiorIdade)
	usuario2.fazAniversario()
	fmt.Println("Fez Aniversario 1", usuario2.idade)
	maiorIdade = usuario2.maiorIdade()
	fmt.Println(maiorIdade)
	usuario2.fazAniversario()
	fmt.Println("Fez Aniversario 2", usuario2.idade)
	maiorIdade = usuario2.maiorIdade()
	fmt.Println(maiorIdade)

}

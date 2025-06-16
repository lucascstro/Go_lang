package auxiliar

import "fmt"

func Escrever() {
	fmt.Println("Escrevendo do pacote auxiliar")
	escrever2()
}
func escrever2() {
	fmt.Println("Escrevendo localmente do pacote")
}

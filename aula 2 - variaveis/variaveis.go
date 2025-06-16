package main

import (
	"errors"
	"fmt"
)

func main() {
	//formas de declaracao de variaveis:
	var nome string = "Lucas"
	sobrenome := "Castro"

	var (
		variavel1 string = "a"
		variavel2 string = "b"
	)

	variavel3, variavel4, variavel5 := 1, 2, false
	const constante1 string = "constante"
	var(
		numero1 int8 = 123
		numero2 int16 = 12345
		numero3 int32 = 1234567890
		numero4 int64 = 123456789012345678
		numero5 byte = 123 //byte e int8 guardam o mesmo tamanho
		numero6 uint = 12345678901234567834 //pega arch do pc
		numero7 rune = 1234567890 //alias para int32
		numeroFloat32 float32= 11230.01
		numeroFloat64 float64= 1123123123230.01
		booleano bool = false
		erro error = errors.New("Mensagem de erro")
	)
	numeroFloatInf := 12345.01
	booleano1 := false
	char :='B' //retorna int da posicao na tabela asc

	
	fmt.Println(nome + " " + sobrenome)
	fmt.Println(variavel1, variavel2, variavel3, variavel4, variavel5)
	fmt.Println(numero1, numero2, numero3, numero4,numero5,numero6,numero7,numeroFloat32,numeroFloat64,booleano,erro,numeroFloatInf, booleano1, char)

}

package main

import ("fmt")
func main() {
	var variavel int = 10 // tipagem estatica
	fmt.Println(variavel)

	variavel2 := 10 // declaração com inferencia de tipo
	fmt.Println(variavel2)

	var variavel3 int
	variavel3 = 10 // atribuindo valor a variavel3
	fmt.Println(variavel3)

	var a,b,c,d int = 1 , 2, 3, 4 // variaveis em uma linha
	fmt.Println(a, b, c, d)

	var(
		e int = 5
		f float32 = 5.5
		g string = "cinco"
	)
	fmt.Println(e, f, g)

	const EULER float64 = 2.718281828459045235360287471352662497757
	fmt.Println(EULER)
}
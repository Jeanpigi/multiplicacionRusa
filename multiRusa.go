package main

import "fmt"


func main()  {
	var n, m, multiplicando, multiplicador, suma int

	fmt.Println("Escriba el valor de n")
	fmt.Scanf("%d\n", &n)

	fmt.Println("Escriba el valor de m")
	fmt.Scanf("%d\n", &m)

	suma = 0
	multiplicando = n
	multiplicador = m


	for multiplicador >= 1 {
		if multiplicador % 2 == 1 {
			suma = suma + multiplicando
		}

		multiplicando = multiplicando * 2
		multiplicador = multiplicador / 2

		fmt.Printf("el multiplicando es %d\n", multiplicando)
		fmt.Printf("El multiplicador es %d\n", multiplicador)
	}

	fmt.Printf("El valor de la multiplicación rusa de %d por %d es %d\n", n, m, suma)
}
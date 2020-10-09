package main

import "fmt"

func main() {
	const maConstante int = 50

	maConstante = 50

	fmt.Println("ma Constante : ", maConstante)
}

// will print error, of course we couldn't modify a constant !

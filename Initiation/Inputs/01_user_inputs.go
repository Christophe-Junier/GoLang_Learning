package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin) // Scanner creation capturing user input
	fmt.Print("Type something and press Enter : ")
	scanner.Scan()                      // Launch scanner
	entreeUtilisateur := scanner.Text() // Stocker scanner result
	fmt.Println(entreeUtilisateur)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Type a number then press enter : ")
	scanner.Scan()

	nbr, _ := strconv.Atoi(scanner.Text()) // conversion of string in int

	fmt.Printf("res : %d\n", (nbr + 6))
}

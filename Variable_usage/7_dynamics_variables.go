package main

import "fmt"

func main() {
	/*
	   dynamics variables
	*/
	flt := 15.5   //  will be type float
	in := 5       //  will be type int
	st := "hello" //  will be type string
	bol := true   //  will be type boolean

	fmt.Printf("type of variable flt is %T\n", flt)
	fmt.Printf("type of variable in is %T\n", in)
	fmt.Printf("type of variable st is %T\n", st)
	fmt.Printf("type of variable bol is %T\n", bol)
}

/*
Printf() print text and variables but
contrary to fonction Println()
Printf() can show more informations,
just add a special symbol

here are some:

	 %T : type of a value
	 %d : integer
	 %s : string
	 %f : decimal
	 %b : binary representation
*/

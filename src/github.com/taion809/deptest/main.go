package main

import (
	"fmt"
	stringer "github.com/taion809/depteststringer"
	upper "github.com/taion809/deptestupstringer"
)

func printStringUpcased(f stringer.Stringer) {
	fmt.Println(upper.Upper(f))
}

func printString(f stringer.Stringer) {
	fmt.Println(f.Get())
}

func main() {
	fmt.Println("Lets do stringering!")
	f := stringer.NewStringer("I want to print this")
	printString(f)
	printStringUpcased(f)
}

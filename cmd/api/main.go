package main

import (
	read "Pars/internal/htmlRead"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Извлечение...")
	b := read.Read(os.Args[1])
	read.Sort(b)
}

package main

import (
	"fmt"
)

func main() {
	defer func() { fmt.Println("before the panic") }()
	panic("paniced")
}

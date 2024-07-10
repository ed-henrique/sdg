package main

import (
	"fmt"

	"github.com/ed-henrique/sdg/pkg"
)

func main() {
	cpfs := pkg.CPF(100, pkg.WithFormat())

	for _, c := range cpfs {
		fmt.Println(c)
	}
}

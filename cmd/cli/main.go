package main

import (
	"fmt"

	"github.com/ed-henrique/sdg/pkg"
)

func main() {
  cpfs := pkg.CPFWithOptions(100, pkg.Options{ Format: true });

  for _, c := range cpfs {
    fmt.Println(c)
  }
}

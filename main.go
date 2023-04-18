package main

import (
	"fmt"

	"github.com/michaeldouglas/mvcgo/initializers"
)

func init() {
	initializers.LoadEnv()
	initializers.Connection()
}

func main() {
	fmt.Println("Alterado com sucesso!")
}

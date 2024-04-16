package main

import (
	"fmt"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Olá mundo")
	erro := checkmail.ValidateFormat("marcio.kgr@gmail.com")
	fmt.Println(erro)
}

package main

import (
	"app-cli/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Result:")

	aplication := app.Create()
	if erro := aplication.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}

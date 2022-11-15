package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Ricardo-Sales/comand-line-app/app"
)

func main() {
	fmt.Println("ponto de partida")

	aplicacao := app.Gerar()
	if err := aplicacao.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}

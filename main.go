package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gideaopinheiro/compilador/lexico"
	"github.com/gideaopinheiro/compilador/sintatico"
	"github.com/gideaopinheiro/compilador/web"
)

func main() {
	server := web.NewServer()
	server.Start()
	filename := os.Args[1]

	inputBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Erro ao ler arquivo: %v\n", err)
		return
	}

	input := string(inputBytes)
	lxc := lexico.NewLexico()
	symbolTable, err := lxc.CreateSymbolTable(input)

	if err == nil {
		sintatico := sintatico.NewSintatico(symbolTable)
		sintatico.Start()
	} else {
		fmt.Println(err.Error())
	}
}

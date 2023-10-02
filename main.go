package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/gideaopinheiro/compilador/sintatico"
	"github.com/gideaopinheiro/compilador/tipos"
)

func isToken(value string, tokenPatterns []tipos.TokenPattern) (bool, string) {
	for _, tp := range tokenPatterns {
		if regexp.MustCompile(fmt.Sprintf(`(?i)^%s$`, tp.Pattern)).MatchString(value) {
			return true, tp.TokenType
		}
	}
	return false, ""
}

func CreateSymbolTable(input string, tokenPatterns []tipos.TokenPattern) (map[int]tipos.Token, error) {
	symbols := make(map[int]tipos.Token)
	lines := strings.Split(input, "\n")
	index := 1

	for lineIdx, line := range lines {
		commentIndex := strings.Index(line, "// ")
		if commentIndex >= 0 {
			line = line[:commentIndex]
			lines[lineIdx] = line
		}

		var accumulator string
		for colIdx, char := range lines[lineIdx] {
			if char == ' ' || char == ';' {
				if flag, tokenType := isToken(accumulator, tokenPatterns); flag {
					symbols[index] = tipos.Token{
						TokenType:  tokenType,
						TokenValue: accumulator,
					}
					index++
				} else {
					message := fmt.Sprintf("Syntax error at line [%d] col [%d]: %s\n", lineIdx+1, colIdx-1, accumulator)
					return nil, errors.New(message)
				}
				if char == ';' {
					_, tokenType := isToken(string(char), tokenPatterns)
					symbols[index] = tipos.Token{
						TokenType:  tokenType,
						TokenValue: string(char),
					}
					index++
				}
				accumulator = ""
			} else {
				accumulator += string(char)
			}
		}

		if accumulator != "" {
			if flag, tokenType := isToken(accumulator, tokenPatterns); flag {
				symbols[index] = tipos.Token{
					TokenType:  tokenType,
					TokenValue: accumulator,
				}
			}
		}
	}

	return symbols, nil
}

func main() {
	filename := os.Args[1]

	inputBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Erro ao ler arquivo: %v\n", err)
		return
	}

	input := string(inputBytes)

	tokenPatterns := []tipos.TokenPattern{
		{TokenType: "INICIOPROGRAMA", Pattern: "programa_SOL"},
		{TokenType: "LOOP", Pattern: "loop"},
		{TokenType: "TEMPO", Pattern: "20_min|1_hora|1_dia|2_dias|sem_limite|15_min"},
		{TokenType: "VEZES", Pattern: "1|2|3|4|5"},
		{TokenType: "NAVEGADOR", Pattern: "navegador"},
		{TokenType: "DELIMITADOR", Pattern: ";"},
		{TokenType: "LINK_PDF", Pattern: `https://[a-zA-Z0-9._-]+\.com/[a-zA-Z0-9._-]+\.pdf`},
		{TokenType: "LINK_VIDEO", Pattern: "https://youtube.com"},
		{TokenType: "LINK_VIDEOCONFERENCIA", Pattern: "https://meet.google.com"},
		{TokenType: "LINK_WHATSAPP_WEB", Pattern: "https://web.whatsapp.com"},
		{TokenType: "LINK_EMAIL", Pattern: "https://gmail.com"},
	}

	symbolTable, err := CreateSymbolTable(input, tokenPatterns)
	if err == nil {
		sintatico := sintatico.NewSintatico(symbolTable)
		sintatico.Start()
	} else {
		fmt.Println(err.Error())
	}
}

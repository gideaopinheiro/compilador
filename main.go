package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

type Token struct {
	TokenType  string
	TokenValue string
}

type TokenPattern struct {
	TokenType string
	Pattern   string
}

type Symbol struct {
	Index      int
	TokenType  string
	TokenValue string
}

func isToken(value string, tokenPatterns []TokenPattern) (bool, string) {
	for _, tp := range tokenPatterns {
		if regexp.MustCompile(fmt.Sprintf(`(?i)^%s$`, tp.Pattern)).MatchString(value) {
			return true, tp.TokenType
		}
	}
	return false, ""
}

func shouldIgnoreLine(line string) bool {
	return strings.HasPrefix(strings.TrimSpace(line), "//")
}

func CreateSymbolTable(input string, tokenPatterns []TokenPattern) ([]Symbol, error) {
	var symbols []Symbol
	lines := strings.Split(input, "\n")

	for lineIdx, line := range lines {
		if shouldIgnoreLine(line) {
			continue
		}

		var accumulator string
		for colIdx, char := range line {
			if char == ' ' || char == ';' {
				if flag, tokenType := isToken(accumulator, tokenPatterns); flag {
					symbols = append(symbols, Symbol{
						Index:      len(symbols) + 1,
						TokenType:  tokenType,
						TokenValue: accumulator,
					})
				} else {
					message := fmt.Sprintf("Syntax error at line [%d] col [%d]: %s\n", lineIdx+1, colIdx-1, accumulator)
					return nil, errors.New(message)
				}
				if char == ';' {
					_, tokenType := isToken(string(char), tokenPatterns)
					symbols = append(symbols, Symbol{
						Index:      len(symbols) + 1,
						TokenType:  tokenType,
						TokenValue: string(char),
					})
				}
				accumulator = ""
			} else {
				accumulator += string(char)
			}
		}

		if accumulator != "" {
			if flag, tokenType := isToken(accumulator, tokenPatterns); flag {
				symbols = append(symbols, Symbol{
					Index:      len(symbols) + 1,
					TokenType:  tokenType,
					TokenValue: accumulator,
				})
			}
		}
	}

	return symbols, nil
}

func Program() {
	// programa_SOL loop vezes sequência
}

func main() {
	filename := os.Args[1]

	inputBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Erro ao ler arquivo: %v\n", err)
		return
	}

	input := string(inputBytes)

	tokenPatterns := []TokenPattern{
		{"INICIOPROGRAMA", "programa_SOL"},
		{"LOOP", "loop"},
		{"VEZES", "1|2|3|4|5"},
		{"TEMPO", "20_min|1_hora|1_dia|2_dias|sem_limite|15_min"},
		{"NAVEGADOR", "navegador"},
		{"DELIMITADOR", ";"},
		{"LINK_PDF", `https://[a-zA-Z0-9._-]+\.com/[a-zA-Z0-9._-]+\.pdf`},
		{"LINK_VIDEO", "https://youtube.com"},
		{"LINK_VIDEOCONFERENCIA", "https://meet.google.com"},
		{"LINK_WHATSAPP_WEB", "https://web.whatsapp.com"},
		{"LINK_EMAIL", "https://gmail.com"},
	}

	symbolTable, err := CreateSymbolTable(input, tokenPatterns)

	if err == nil {
		for _, symbol := range symbolTable {
			fmt.Println(symbol)
		}
	} else {
		fmt.Println(err.Error())
	}
}

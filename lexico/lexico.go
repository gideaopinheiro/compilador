package lexico

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/gideaopinheiro/compilador/tipos"
)

type Lexico struct {
	tokenPatterns []tipos.TokenPattern
}

func NewLexico() *Lexico {
	return &Lexico{
		tokenPatterns: []tipos.TokenPattern{
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
		},
	}
}

func (l *Lexico) isToken(value string) (bool, string) {
	for _, tp := range l.tokenPatterns {
		if regexp.MustCompile(fmt.Sprintf(`(?i)^%s$`, tp.Pattern)).MatchString(value) {
			return true, tp.TokenType
		}
	}
	return false, ""
}

func (l *Lexico) CreateSymbolTable(input string) (map[int]tipos.Token, error) {
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
		for _, char := range lines[lineIdx] {
			if char == ' ' || char == ';' {
				if flag, tokenType := l.isToken(accumulator); flag {
					symbols[index] = tipos.Token{
						TokenType:  tokenType,
						TokenValue: accumulator,
						Line:       lineIdx + 1,
					}
					index++
				} else {
					message := fmt.Sprintf("[Erro léxico]\nlinha %d: %s\n", lineIdx+1, accumulator)
					return nil, errors.New(message)
				}
				if char == ';' {
					_, tokenType := l.isToken(string(char))
					symbols[index] = tipos.Token{
						TokenType:  tokenType,
						TokenValue: string(char),
						Line:       lineIdx + 1,
					}
					index++
				}
				accumulator = ""
			} else {
				accumulator += string(char)
			}
		}

		if accumulator != "" {
			if flag, tokenType := l.isToken(accumulator); flag {
				symbols[index] = tipos.Token{
					TokenType:  tokenType,
					TokenValue: accumulator,
					Line:       lineIdx + 1,
				}
			}
		}
	}

	return symbols, nil
}

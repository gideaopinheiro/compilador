package sintatico

import (
	"fmt"

	"github.com/gideaopinheiro/compilador/tipos"
)

type Sintatico struct {
	symbolTable  map[int]tipos.Token
	currentIndex int
}

func NewSintatico(sTable map[int]tipos.Token) *Sintatico {
	return &Sintatico{
		symbolTable:  sTable,
		currentIndex: 1,
	}
}

func (s *Sintatico) Start() {
	if s.programa() {
		fmt.Println("Análise sintática bem-sucedida.")
	} else {
		fmt.Println("Erro de análise sintática.")
	}
}

func (s *Sintatico) programa() bool {
	if s.tokenMatches("INICIOPROGRAMA") {
		if s.loop() && s.vezes() && s.sequencia() {
			return true
		}
	}
	return false
}

func (s *Sintatico) loop() bool {
	return s.tokenMatches("LOOP")
}

func (s *Sintatico) vezes() bool {
	return s.tokenMatches("VEZES")
}

func (s *Sintatico) sequencia() bool {
	if s.present() {
		return true
	} else if s.fasesEPIC() {
		return true
	}
	return false
}

func (s *Sintatico) fasesEPIC() bool {
	if s.explore() || s.present() || s.interact() || s.critique() {
		return true
	}
	return false
}

func (s *Sintatico) explore() bool {
	if s.navegar() && s.tempo() && s.tokenMatches("DELIMITADOR") {
		return true
	}
	return false
}

func (s *Sintatico) present() bool {
	if s.visualizarPDF() {
		return true
	} else if s.visualizarVideo() {
		return true
	} else if s.videoconferencia() && s.tempo() && s.tokenMatches("DELIMITADOR") {
		return true
	}
	return false
}

func (s *Sintatico) interact() bool {
	if s.whatsappWeb() {
		return true
	} else if s.email() {
		return true
	} else if s.videoconferencia() && s.tempo() && s.tokenMatches("DELIMITADOR") {
		return true
	}
	return false
}

func (s *Sintatico) critique() bool {
	if s.whatsappWeb() {
		return true
	} else if s.email() {
		return true
	} else if s.videoconferencia() && s.tempo() && s.tokenMatches("DELIMITADOR") {
		return true
	}
	return false
}

func (s *Sintatico) browser() bool {
	return s.tokenMatches("NAVEGADOR")
}

func (s *Sintatico) visualizarPDF() bool {
	if s.browser() && s.linkPDF() {
		return true
	}
	s.currentIndex--
	return false
}

func (s *Sintatico) visualizarVideo() bool {
	if s.browser() && s.linkVideo() {
		return true
	}
	s.currentIndex--
	return false
}

func (s *Sintatico) videoconferencia() bool {
	if s.browser() && s.linkVideoconferencia() {
		return true
	}
	s.currentIndex--
	return false
}

func (s *Sintatico) whatsappWeb() bool {
	if s.browser() && s.linkWhatsappWeb() {
		return true
	}
	s.currentIndex--
	return false
}

func (s *Sintatico) email() bool {
	if s.browser() && s.linkEmail() {
		return true
	}
	s.currentIndex--
	return false
}

func (s *Sintatico) tempo() bool {
	return s.tokenMatches("TEMPO")
}

func (s *Sintatico) navegar() bool {
	return s.tokenMatches("NAVEGADOR")
}

func (s *Sintatico) linkPDF() bool {
	return s.tokenMatches("LINK_PDF")
}

func (s *Sintatico) linkVideo() bool {
	return s.tokenMatches("LINK_VIDEO")
}

func (s *Sintatico) linkVideoconferencia() bool {
	return s.tokenMatches("LINK_VIDEOCONFERENCIA")
}

func (s *Sintatico) linkWhatsappWeb() bool {
	return s.tokenMatches("LINK_WHATSAPP_WEB")
}

func (s *Sintatico) linkEmail() bool {
	return s.tokenMatches("LINK_EMAIL")
}

func (s *Sintatico) tokenMatches(expectedTokenType string) bool {
	if s.currentIndex <= len(s.symbolTable) && s.symbolTable[s.currentIndex].TokenType == expectedTokenType {
		s.currentIndex++
		return true
	}
	return false
}

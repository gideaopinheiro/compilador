package sintatico

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/gideaopinheiro/compilador/tipos"
)

type Sintatico struct {
	symbolTable    map[int]tipos.Token
	currentIndex   int
	stack          []error
	addressesQueue []string
}

func NewSintatico(sTable map[int]tipos.Token) *Sintatico {
	return &Sintatico{
		symbolTable:    sTable,
		currentIndex:   1,
		stack:          nil,
		addressesQueue: nil,
	}
}

func (s *Sintatico) Start() (string, bool) {
	if s.programa() {
		launchBrowser(s.addressesQueue)
		return "Ok.", true
	} else {
		return fmt.Sprintf("[Erro sintático]\n%v\n", s.stack[len(s.stack)-1]), false
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
	if s.fasesEPIC() && s.recurSequencia() {
		return true
	}
	return false
}

func (s *Sintatico) recurSequencia() bool {
	if s.currentIndex >= len(s.symbolTable)-1 {
		return true
	}
	return s.sequencia()
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
	if match := s.tokenMatches("TEMPO"); match {
		if s.symbolTable[s.currentIndex-2].TokenType == "NAVEGADOR" {
			s.addressesQueue = append(s.addressesQueue, "https://google.com")
		}
		return true
	}
	return false
}

func (s *Sintatico) navegar() bool {
	return s.browser()
}

func (s *Sintatico) linkPDF() bool {
	if match := s.tokenMatches("LINK_PDF"); match {
		s.addressesQueue = append(s.addressesQueue, s.symbolTable[s.currentIndex-1].TokenValue)
		return true
	}
	return false
}

func (s *Sintatico) linkVideo() bool {
	if match := s.tokenMatches("LINK_VIDEO"); match {
		s.addressesQueue = append(s.addressesQueue, s.symbolTable[s.currentIndex-1].TokenValue)
		return true
	}
	return false
}

func (s *Sintatico) linkVideoconferencia() bool {
	if match := s.tokenMatches("LINK_VIDEOCONFERENCIA"); match {
		s.addressesQueue = append(s.addressesQueue, s.symbolTable[s.currentIndex-1].TokenValue)
		return true
	}
	return false
}

func (s *Sintatico) linkWhatsappWeb() bool {
	if match := s.tokenMatches("LINK_WHATSAPP_WEB"); match {
		s.addressesQueue = append(s.addressesQueue, s.symbolTable[s.currentIndex-1].TokenValue)
		return true
	}
	return false
}

func (s *Sintatico) linkEmail() bool {
	if match := s.tokenMatches("LINK_EMAIL"); match {
		s.addressesQueue = append(s.addressesQueue, s.symbolTable[s.currentIndex-1].TokenValue)
		return true
	}
	return false
}

func launchBrowser(addresses []string) {
	for _, address := range addresses {
		cmd := exec.Command("xdg-open", address)
		err := cmd.Run()
		if err != nil {
			panic(err)
		}
		time.Sleep(5 * time.Second)

		closeCmd := exec.Command("pkill", "chrome")
		err = closeCmd.Run()
		if err != nil {
			panic(err)
		}
	}
}

func (s *Sintatico) tokenMatches(expectedTokenType string) bool {
	if s.currentIndex <= len(s.symbolTable) && s.symbolTable[s.currentIndex].TokenType == expectedTokenType {
		s.currentIndex++
		return true
	}
	s.stack = append(s.stack, fmt.Errorf("linha %d : \"%s\"  ", s.symbolTable[s.currentIndex].Line, s.symbolTable[s.currentIndex].TokenValue))
	return false
}

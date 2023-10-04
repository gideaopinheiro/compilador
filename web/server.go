package web

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/gideaopinheiro/compilador/lexico"
	"github.com/gideaopinheiro/compilador/sintatico"
)

type Server struct {
	severName string
}

type Program struct {
	code string
}

type ResponseData struct {
	IsCorrect bool
	Response  string
}

func NewServer() *Server {
	return &Server{
		severName: "Interpretador SOL",
	}
}

func interpretadorHandler(w http.ResponseWriter, r *http.Request) {
	lxc := lexico.NewLexico()

	programa := r.PostFormValue("corpo-programa")

	symbolTable, err := lxc.CreateSymbolTable(programa)

	var response string
	var isCorrect bool
	if err == nil {
		sintatico := sintatico.NewSintatico(symbolTable)
		response, isCorrect = sintatico.Start()
	} else {
		response = err.Error()
		isCorrect = false
		fmt.Println(err.Error())
	}

	tmpl := template.Must(template.ParseFiles("./web/response.html"))
	tmpl.Execute(w, ResponseData{IsCorrect: isCorrect, Response: response})
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./web/index.html"))
	tmpl.Execute(w, nil)
}

func (s *Server) Start() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/interpretador", interpretadorHandler)

	log.Fatal(http.ListenAndServe(":3737", nil))
}

package tipos

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

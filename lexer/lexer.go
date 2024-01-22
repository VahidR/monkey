package lexer

type Lexer struct {
	input        string
	position     int  // current char in the input
	readPosition int  // already read char in the input
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}

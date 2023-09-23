package lexer

import "strings"

type Instruction struct {
	symbol OPSym
	params []string
}

type Lexer struct {
	code         string
	instructions []Instruction
}

func NewLexer(code string) *Lexer {
	return &Lexer{
		code: code,
	}
}

func (l *Lexer) Lex() {
	for _, line := range strings.Split(l.code, "\n") {
		if strings.TrimSpace(line) == "" {
			continue
		}
		l.instructions = append(l.instructions, l.parseLine(line))
	}
}

func (l *Lexer) parseLine(line string) Instruction {
	var instruction Instruction
	parts := strings.Split(line, " ")
	instruction.symbol = OPSym(parts[0])
	if len(parts) > 1 {
		parts = strings.Split(parts[1], ",")
	}
	instruction.params = parts
	return instruction
}

func (l *Lexer) String() string {
	var str strings.Builder
	for _, instruction := range l.instructions {
		str.WriteString(string(instruction.symbol))
		for _, param := range instruction.params {
			str.WriteString(" ")
			str.WriteString(param)
		}
		str.WriteString("\n")
	}
	return str.String()
}

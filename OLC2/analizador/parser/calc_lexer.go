// Code generated from CalcLexer.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 23, 122,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 6, 5, 68, 10, 5,
	13, 5, 14, 5, 69, 3, 6, 3, 6, 7, 6, 74, 10, 6, 12, 6, 14, 6, 77, 11, 6,
	3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3,
	11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15,
	3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3,
	21, 3, 21, 3, 22, 6, 22, 114, 10, 22, 13, 22, 14, 22, 115, 3, 22, 3, 22,
	3, 23, 3, 23, 3, 23, 2, 2, 24, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15,
	9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33,
	18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45, 2, 3, 2, 6, 3, 2, 50, 59,
	3, 2, 36, 36, 6, 2, 11, 12, 15, 15, 34, 34, 94, 94, 9, 2, 34, 35, 37, 37,
	45, 45, 47, 48, 60, 60, 66, 66, 93, 95, 2, 123, 2, 3, 3, 2, 2, 2, 2, 5,
	3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13,
	3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2,
	21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2,
	2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2,
	2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2,
	2, 2, 3, 47, 3, 2, 2, 2, 5, 54, 3, 2, 2, 2, 7, 58, 3, 2, 2, 2, 9, 67, 3,
	2, 2, 2, 11, 71, 3, 2, 2, 2, 13, 80, 3, 2, 2, 2, 15, 82, 3, 2, 2, 2, 17,
	84, 3, 2, 2, 2, 19, 86, 3, 2, 2, 2, 21, 89, 3, 2, 2, 2, 23, 92, 3, 2, 2,
	2, 25, 94, 3, 2, 2, 2, 27, 96, 3, 2, 2, 2, 29, 98, 3, 2, 2, 2, 31, 100,
	3, 2, 2, 2, 33, 102, 3, 2, 2, 2, 35, 104, 3, 2, 2, 2, 37, 106, 3, 2, 2,
	2, 39, 108, 3, 2, 2, 2, 41, 110, 3, 2, 2, 2, 43, 113, 3, 2, 2, 2, 45, 119,
	3, 2, 2, 2, 47, 48, 7, 117, 2, 2, 48, 49, 7, 123, 2, 2, 49, 50, 7, 117,
	2, 2, 50, 51, 7, 118, 2, 2, 51, 52, 7, 103, 2, 2, 52, 53, 7, 111, 2, 2,
	53, 4, 3, 2, 2, 2, 54, 55, 7, 113, 2, 2, 55, 56, 7, 119, 2, 2, 56, 57,
	7, 118, 2, 2, 57, 6, 3, 2, 2, 2, 58, 59, 7, 114, 2, 2, 59, 60, 7, 116,
	2, 2, 60, 61, 7, 107, 2, 2, 61, 62, 7, 112, 2, 2, 62, 63, 7, 118, 2, 2,
	63, 64, 7, 110, 2, 2, 64, 65, 7, 112, 2, 2, 65, 8, 3, 2, 2, 2, 66, 68,
	9, 2, 2, 2, 67, 66, 3, 2, 2, 2, 68, 69, 3, 2, 2, 2, 69, 67, 3, 2, 2, 2,
	69, 70, 3, 2, 2, 2, 70, 10, 3, 2, 2, 2, 71, 75, 7, 36, 2, 2, 72, 74, 10,
	3, 2, 2, 73, 72, 3, 2, 2, 2, 74, 77, 3, 2, 2, 2, 75, 73, 3, 2, 2, 2, 75,
	76, 3, 2, 2, 2, 76, 78, 3, 2, 2, 2, 77, 75, 3, 2, 2, 2, 78, 79, 7, 36,
	2, 2, 79, 12, 3, 2, 2, 2, 80, 81, 7, 48, 2, 2, 81, 14, 3, 2, 2, 2, 82,
	83, 7, 61, 2, 2, 83, 16, 3, 2, 2, 2, 84, 85, 7, 35, 2, 2, 85, 18, 3, 2,
	2, 2, 86, 87, 7, 64, 2, 2, 87, 88, 7, 63, 2, 2, 88, 20, 3, 2, 2, 2, 89,
	90, 7, 62, 2, 2, 90, 91, 7, 63, 2, 2, 91, 22, 3, 2, 2, 2, 92, 93, 7, 64,
	2, 2, 93, 24, 3, 2, 2, 2, 94, 95, 7, 62, 2, 2, 95, 26, 3, 2, 2, 2, 96,
	97, 7, 44, 2, 2, 97, 28, 3, 2, 2, 2, 98, 99, 7, 49, 2, 2, 99, 30, 3, 2,
	2, 2, 100, 101, 7, 45, 2, 2, 101, 32, 3, 2, 2, 2, 102, 103, 7, 47, 2, 2,
	103, 34, 3, 2, 2, 2, 104, 105, 7, 42, 2, 2, 105, 36, 3, 2, 2, 2, 106, 107,
	7, 43, 2, 2, 107, 38, 3, 2, 2, 2, 108, 109, 7, 125, 2, 2, 109, 40, 3, 2,
	2, 2, 110, 111, 7, 127, 2, 2, 111, 42, 3, 2, 2, 2, 112, 114, 9, 4, 2, 2,
	113, 112, 3, 2, 2, 2, 114, 115, 3, 2, 2, 2, 115, 113, 3, 2, 2, 2, 115,
	116, 3, 2, 2, 2, 116, 117, 3, 2, 2, 2, 117, 118, 8, 22, 2, 2, 118, 44,
	3, 2, 2, 2, 119, 120, 7, 94, 2, 2, 120, 121, 9, 5, 2, 2, 121, 46, 3, 2,
	2, 2, 6, 2, 69, 75, 115, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'system'", "'out'", "'println'", "", "", "'.'", "';'", "'!'", "'>='",
	"'<='", "'>'", "'<'", "'*'", "'/'", "'+'", "'-'", "'('", "')'", "'{'",
	"'}'",
}

var lexerSymbolicNames = []string{
	"", "SYSTEM", "OUT", "PRINTLN", "NUMBER", "STRING", "PUNTO", "PTCOMA",
	"DIFERENTE", "MAYORIGUAL", "MENORIGUAL", "MAYOR", "MENOR", "MUL", "DIV",
	"ADD", "SUB", "LP", "RP", "L_LLAVE", "R_LLAVE", "WHITESPACE",
}

var lexerRuleNames = []string{
	"SYSTEM", "OUT", "PRINTLN", "NUMBER", "STRING", "PUNTO", "PTCOMA", "DIFERENTE",
	"MAYORIGUAL", "MENORIGUAL", "MAYOR", "MENOR", "MUL", "DIV", "ADD", "SUB",
	"LP", "RP", "L_LLAVE", "R_LLAVE", "WHITESPACE", "ESC_SEQ",
}

type CalcLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewCalcLexer(input antlr.CharStream) *CalcLexer {

	l := new(CalcLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "CalcLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// CalcLexer tokens.
const (
	CalcLexerSYSTEM     = 1
	CalcLexerOUT        = 2
	CalcLexerPRINTLN    = 3
	CalcLexerNUMBER     = 4
	CalcLexerSTRING     = 5
	CalcLexerPUNTO      = 6
	CalcLexerPTCOMA     = 7
	CalcLexerDIFERENTE  = 8
	CalcLexerMAYORIGUAL = 9
	CalcLexerMENORIGUAL = 10
	CalcLexerMAYOR      = 11
	CalcLexerMENOR      = 12
	CalcLexerMUL        = 13
	CalcLexerDIV        = 14
	CalcLexerADD        = 15
	CalcLexerSUB        = 16
	CalcLexerLP         = 17
	CalcLexerRP         = 18
	CalcLexerL_LLAVE    = 19
	CalcLexerR_LLAVE    = 20
	CalcLexerWHITESPACE = 21
)
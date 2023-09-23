package lexer

type OPSym string
type OPCode uint8

// Instruction set
const (
	LDA OPSym = "LDA"
	STA OPSym = "STA"
	BRK OPSym = "BRK"
)

const (
	LDA_IMM  OPCode = 0xA9
	LDA_ZP   OPCode = 0xA5
	LDA_ZPX  OPCode = 0xB5
	LDA_ABS  OPCode = 0xAD
	LDA_ABSX OPCode = 0xBD
	LDA_ABSY OPCode = 0xB9
	LDA_INDX OPCode = 0xA1
	LDA_INDY OPCode = 0xB1
)

const (
	STA_ZP  OPCode = 0x85
	STA_ZPX OPCode = 0x95
	STA_ABS OPCode = 0x8D
	STA_ABX OPCode = 0x9D
	STA_ABY OPCode = 0x99
	STA_INX OPCode = 0x81
	STA_INY OPCode = 0x91
)

const BRK_IMP OPCode = 0x00

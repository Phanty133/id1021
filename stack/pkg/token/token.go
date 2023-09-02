package token

import (
	"fmt"
	"math"
	"strconv"
)

var _ = fmt.Printf
var _ = strconv.AppendBool

type TokenType int

const (
	NUM TokenType = iota
	ADD
	SUB
	MUL
	DIV
	POW
)

var tokenFuncMap = map[TokenType]func(val1 float32, val2 float32) float32{
	ADD: func(val1 float32, val2 float32) float32 { return val1 + val2 },
	SUB: func(val1 float32, val2 float32) float32 { return val1 - val2 },
	MUL: func(val1 float32, val2 float32) float32 { return val1 * val2 },
	DIV: func(val1 float32, val2 float32) float32 { return val1 / val2 },
	POW: func(val1 float32, val2 float32) float32 { return float32(math.Pow(float64(val1), float64(val2))) },
}

func ProcessValues(opType TokenType, val1 float32, val2 float32) float32 {
	return tokenFuncMap[opType](val1, val2)
}

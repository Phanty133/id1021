package solver

import (
	"fmt"
	"regexp"
	"strings"
	"strconv"

	"github.com/phanty133/id1021/stack/pkg/stacks"
	"github.com/phanty133/id1021/stack/pkg/token"
)

var _ = fmt.Append

type Expression[StackType stacks.Stack[float32]] struct {
	numStack StackType
	ops []token.TokenType
}

func (stack *Expression[StackType]) PopulateExpression(expr string) error {
	charMap := map[string]token.TokenType{
		"+": token.ADD,
		"-": token.SUB,
		"*": token.MUL,
		"/": token.DIV,
		"^": token.POW,
	}

	vals := strings.Split(expr, " ")
	numPattern := regexp.MustCompile(`[\d.]+`)
	stack.ops = make([]token.TokenType, 0, len(vals))

	for _, val := range vals {
		numMatch := numPattern.MatchString(val)

		if numMatch {
			num, _ := strconv.ParseFloat(val, 32)
			stack.numStack.Push(float32(num))
		} else {
			tType, ok := charMap[val]

			if !ok {
				return fmt.Errorf("invalid token: %s", val)
			}

			stack.ops = append(stack.ops, tType)
		}
	}

	return nil
}

func (stack *Expression[StackType]) ParseExpression() (float32, error) {
	if stack.numStack.Empty() {
		return 0, fmt.Errorf("empty expression")
	}

	if len(stack.ops) == 0 {
		return 0, fmt.Errorf("invalid expression")
	}

	for _, op := range stack.ops {
		num1, num1err := stack.numStack.Pop()
		num2, num2err := stack.numStack.Pop()

		if num1err != nil || num2err != nil {
			return 0, fmt.Errorf("invalid expression")
		}

		result := token.ProcessValues(op, num1, num2)
		stack.numStack.Push(result)
	}
	
	return stack.numStack.Pop()
}

func Solve[StackType stacks.Stack[float32]](numStack StackType, expr string) (float32, error) {
	// Format string to remove space duplicates
	spacePattern := regexp.MustCompile(`\s{2,}`)
	expr = spacePattern.ReplaceAllString(expr, " ")

	exprStack := Expression[StackType]{
		numStack: numStack,
		// Ops will be set in PopulateExpression
	}

	popErr := exprStack.PopulateExpression(expr)

	if popErr != nil {
		return 0, popErr
	}

	val, parseErr := exprStack.ParseExpression()

	return val, parseErr
}

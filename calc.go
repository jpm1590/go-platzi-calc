package mycalc

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct {
}

func (calc calc) operate(in string, operator string) int {

	cleanIn := strings.Split(in, operator)

	operator1 := parse2Int(cleanIn[0])
	operator2 := parse2Int(cleanIn[1])

	switch operator {
	case "+":
		return operator1 + operator2
	case "-":
		return operator1 - operator2
	case "*":
		return operator1 * operator2
	case "/":
		return operator1 / operator2
	default:
		fmt.Println("OPERACION NO SOPORTADA")
		return 0
	}
}

func readIn() string {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func parse2Int(text string) int {
	value, _ := strconv.Atoi(text)
	return value
}

package geartoothc

import (
	"fmt"
)

func err_syntax(msg string, line int, msgargs ...any) {
	fmt.Printf("Syntax Error (line %d):\n", line)
	fmsg := fmt.Sprintf(msg, msgargs...)
	fmt.Printf("%s\n", fmsg)
}

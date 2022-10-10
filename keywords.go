package main

import (
	"fmt"
	"strings"
)

func kw_program(args []string) {
	if args[0] == "end" {
		fmt.Println("found program end")
		inprog = false
	} else if args[0] == "start" {
		if inprog {
			err_syntax("Defined new program before current program ended", -1)
			return
		}
		fmt.Println("defined program " + args[1])
		currentprog = args[1]
		inprog = true
	} else {
		err_syntax("Unexpected token '%s'", -1, args[0])
	}
}

func kw_native(args []string) {
	native_arg := strings.Join(args, " ")
	pyil_line := fmt.Sprintf("eval(\"\"\"%s\"\"\")", native_arg)
	add_il(pyil_line)
}

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

func kw_define(args []string) {
	var entry = [2]string{args[0], strings.Join(args[1:], " ")}
	fnmap[fnmap_ptr] = entry
	fnmap_ptr++

	fmt.Printf("Defined %s as %s\n", entry[0], entry[1])
}

func kw_call(args []string) {
	il := find_in_fnmap(args[0])
	add_il(il[1])
}

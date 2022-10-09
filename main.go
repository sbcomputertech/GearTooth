package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	switch os.Args[1] {
	case "compile":
		runcompiler()
	case "about":
		showabout()
	}
}

func runcompiler() {
	fileName := os.Args[2]

	fmt.Println("GearTooth compiling " + fileName)

	content, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	ftext := string(content)
	ftext = strings.ReplaceAll(ftext, "\r", "") // because screw you windows i was stuck on this for half an hour

	code_lines := strings.Split(ftext, "\n")

	for i := 0; i < len(code_lines); i++ {
		fmt.Printf("\nTokenizing line %d:\n", i)
		tokenize(code_lines[i])
	}

	cmd := exec.Command("python", "pyil/transform.py", "main.gtil.pre", "main.gtil")
	cmd.Run()

	os.Remove("main.gtil.pre")
}

func tokenize(inLine string) string {
	splitLine := strings.Split(inLine, " ")

	switch splitLine[0] {
	case "program":
		kw_program(splitLine[1:])
	case "native":
		kw_native(splitLine[1:])
	}

	return "todo"
}

func err_syntax(msg string, line int, msgargs ...any) {
	fmt.Printf("Syntax Error (line %d):\n", line)
	fmsg := fmt.Sprintf(msg, msgargs...)
	fmt.Printf("%s\n", fmsg)
}

func kw_program(args []string) {
	if args[0] == "end" {
		fmt.Println("found program end")
	} else if args[0] == "start" {
		fmt.Println("defined program " + args[1])
	} else {
		err_syntax("Unexpected token '%s'", -1, args[0])
	}
}

func kw_native(args []string) {
	native_arg := strings.Join(args, " ")
	pyil_line := fmt.Sprintf("eval(\"\"\"%s\"\"\")", native_arg)
	add_il(pyil_line)
}

func showabout() {
	fmt.Println("\n---=== GearTooth ===---")
	fmt.Println("~ by reddust9")
}

func add_il(il string) {
	os.WriteFile("main.gtil.pre", []byte(il), fs.ModeAppend)
}

package geartoothc

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var currentprog string
var inprog bool

var fnmap [2048][2]string
var fnmap_ptr = 0

var curr_il string

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
	ftext = strings.ReplaceAll(ftext, "\r", "") // because screw you windows i was stuck on this for half an hour // also probably irrelevant now but i'm too afraid to delete it

	code_lines := strings.Split(ftext, "\n")

	for i := 0; i < len(code_lines); i++ {
		fmt.Printf("\nTokenizing line %d:\n", i)
		tokenize(code_lines[i])
	}

	il_filename := currentprog + ".gtil"

	os.WriteFile(il_filename+".pre", []byte(curr_il), 0666)

	cmd := exec.Command("python", "pyil/transform.py", il_filename+".pre", il_filename)
	cmd.Run()

	os.Remove(il_filename + ".pre")
}

func tokenize(inLine string) {
	splitLine := strings.Split(inLine, " ")

	switch splitLine[0] {
	case "program":
		kw_program(splitLine[1:])
	case "native":
		kw_native(splitLine[1:])
	case "define":
		kw_define(splitLine[1:])
	case "call":
		kw_call(splitLine[1:])
	case "exit":
		var code, error = strconv.Atoi(splitLine[1])
		if error != nil {
			err_syntax("Invalid exit code", -1)
		}
		add_il(fmt.Sprintf("exit(%d)", code))
	case "addlib":
		fmt.Println("addlib not implemented yet")
	}
}

func showabout() {
	fmt.Println("\n---=== GearTooth ===---")
	fmt.Println("~ by reddust9")
}

func add_il(il string) {
	curr_il += "\n"
	curr_il += il
	log_il()
}

func log_il() {
	fmt.Println("[log_il] " + curr_il)
	fmt.Println(fnmap[:5])
}

package main

import (
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println("GearTooth preprocessor running")
}

func invokeGTC(filename string) {
	cmd := exec.Command("geartooth", "compile", filename)
	cmd.Run()
}

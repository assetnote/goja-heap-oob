package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/dop251/goja"
)

func main() {
	// force *Cmd.Run to be emitted by the compiler
	_ = exec.Command("x")

	exploit, err := os.ReadFile("exploit.js")
	if err != nil {
		fmt.Printf("Failed to read file: %s", err)
		return
	}

	vm := goja.New()
	v, err := vm.RunString(string(exploit))
	if err != nil {
		fmt.Printf("Err: %v\n", err)
		return
	}

	fmt.Printf("%v\n", v)
}
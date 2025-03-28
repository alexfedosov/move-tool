package main

import (
	"fmt"
	"github.com/alexfedosov/move-tool/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		fmt.Println("Error:", err)
	}
}

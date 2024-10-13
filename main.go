package main

import (
	"fmt"
	"move-tool/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		fmt.Println("Error:", err)
	}
}

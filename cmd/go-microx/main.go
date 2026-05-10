package main

import (
	"fmt"
	"os"

	"github.com/topworklytechnologies/go-microx/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

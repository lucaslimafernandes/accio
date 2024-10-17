package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	_help := flag.Bool("help", false, "Show available commands")

	flag.Parse()

	if *_help {
		flag.PrintDefaults()
	}

	switch os.Args[1] {
	case "run":
		runnn()
	default:
		fmt.Printf("Unknown subcommand: %s\n", os.Args[1])
	}

	if len(os.Args) == 1 {
		fmt.Println("Maybe you forget some commands, use '-help' to see available commands.")
		return
	}
}

func runnn() {
	fmt.Println("Executing...")
}

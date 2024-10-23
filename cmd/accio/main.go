package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	readfiles "github.com/lucaslimafernandes/pkg/read_files"
	sshexec "github.com/lucaslimafernandes/pkg/sshexec"
)

func main() {

	help := flag.Bool("help", false, "Show available commands")
	version := flag.Bool("version", false, "Show about Accion")

	hostsPath := flag.String("hosts", "", "hosts path")
	runfile := flag.String("run", "", "Runfile path")

	flag.Parse()

	if *help {
		flag.PrintDefaults()
		return
	}

	if *version {
		fmt.Printf(`
[project]
name = "Accio"
version = "v1.0.0"
description = "Accio is a tool designed to manage tasks executed on multiple remote servers via SSH."

[repository]
url = "https://github.com/lucaslimafernandes/accio"
		`)
		return
	}

	if *hostsPath != "" && *runfile != "" {

		tasks, err := readfiles.ReadRunfile(runfile)
		if err != nil {
			log.Fatalln(err)
		}

		sshexec.RunCmd(hostsPath, tasks)

	}

	if len(os.Args) == 1 {
		fmt.Println("Maybe you forget some commands, use '-help' to see available commands.")
		return
	}
}

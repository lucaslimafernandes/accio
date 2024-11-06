package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/lucaslimafernandes/pkg/localexec"
	readfiles "github.com/lucaslimafernandes/pkg/read_files"
	sshexec "github.com/lucaslimafernandes/pkg/sshexec"
)

const VERSION = "v1.1.0"

func main() {

	help := flag.Bool("help", false, "Show available commands")
	version := flag.Bool("version", false, "Show Accio version")

	hostsPath := flag.String("hosts", "", "hosts path")
	localrun := flag.Bool("localrun", false, "Execute local")
	runfile := flag.String("run", "", "Runfile path")

	flag.Parse()

	if *help {
		fmt.Printf(`
[project]
name = "Accio"
version = "v1.0.0"
description = "Accio is a tool designed to manage tasks executed on multiple remote servers via SSH."

[repository]
url = "https://github.com/lucaslimafernandes/accio"

`)
		flag.PrintDefaults()
		return
	}

	if *version {
		fmt.Println(VERSION)
		return
	}

	if *localrun && *runfile != "" {

		tasks, err := readfiles.ReadRunfile(runfile)
		if err != nil {
			log.Fatalln(err)
		}

		err = localexec.ExecCmd(tasks)
		if err != nil {
			log.Fatalln(err)
		}

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

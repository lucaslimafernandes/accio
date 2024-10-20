package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	readfiles "github.com/lucaslimafernandes/pkg/read_files"
	sshexec "github.com/lucaslimafernandes/pkg/ssh_exec"
)

func main() {

	help := flag.Bool("help", false, "Show available commands")
	// run := flag.Bool("run", false, "Execute")
	// host := flag.String("h", "", "host (format: ip:port)")
	// user := flag.String("u", "", "user")
	// keyPath := flag.String("k", "", "key path to SSH private key")

	hostsPath := flag.String("hosts", "", "hosts path")
	runfile := flag.String("run", "", "Runfile path")

	flag.Parse()

	if *help {
		flag.PrintDefaults()
	}

	if *hostsPath != "" {

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

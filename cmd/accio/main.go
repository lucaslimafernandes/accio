package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/lucaslimafernandes/pkg/sshconn"
)

func main() {

	help := flag.Bool("help", false, "Show available commands")
	run := flag.Bool("run", false, "Execute")
	host := flag.String("h", "", "host (format: ip:port)")
	user := flag.String("u", "", "user")
	keyPath := flag.String("k", "", "key path to SSH private key")

	flag.Parse()

	if *help {
		flag.PrintDefaults()
	}

	if *run {
		runnn(*host, *user, *keyPath)
	}

	if len(os.Args) == 1 {
		fmt.Println("Maybe you forget some commands, use '-help' to see available commands.")
		return
	}
}

func runnn(h, u, k string) {

	fmt.Println("Executing...")

	err := sshconn.SSHConn(h, u, k)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Closing...")

}

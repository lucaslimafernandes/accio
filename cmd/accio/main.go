package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	readfiles "github.com/lucaslimafernandes/pkg/read_files"
	"github.com/lucaslimafernandes/pkg/sshconn"
)

func main() {

	help := flag.Bool("help", false, "Show available commands")
	run := flag.Bool("run", false, "Execute")
	host := flag.String("h", "", "host (format: ip:port)")
	user := flag.String("u", "", "user")
	keyPath := flag.String("k", "", "key path to SSH private key")

	hostsPath := flag.String("hosts", "", "hosts path")

	flag.Parse()

	if *help {
		flag.PrintDefaults()
	}

	if *hostsPath != "" {

		var hosts *readfiles.Hosts

		hosts, err := readfiles.ReadHostsFile(hostsPath)
		if err != nil {
			log.Fatalln(err)
		}

		runnn(hosts.Nodes[0].Host, hosts.Nodes[0].User, hosts.Nodes[0].PrivateKeyPath)

		runnn(hosts.Nodes[1].Host, hosts.Nodes[1].User, hosts.Nodes[1].PrivateKeyPath)

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

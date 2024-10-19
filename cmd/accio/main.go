package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"sync"

	readfiles "github.com/lucaslimafernandes/pkg/read_files"
	"github.com/lucaslimafernandes/pkg/sshconn"
)

func main() {

	// /usr/bin/hostname
	comm := []string{"whoami", "hostname"}

	help := flag.Bool("help", false, "Show available commands")
	// run := flag.Bool("run", false, "Execute")
	// host := flag.String("h", "", "host (format: ip:port)")
	// user := flag.String("u", "", "user")
	// keyPath := flag.String("k", "", "key path to SSH private key")

	hostsPath := flag.String("hosts", "", "hosts path")

	flag.Parse()

	if *help {
		flag.PrintDefaults()
	}

	if *hostsPath != "" {

		var wg sync.WaitGroup
		var hosts *readfiles.Hosts

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		hosts, err := readfiles.ReadHostsFile(hostsPath)
		if err != nil {
			log.Fatalln(err)
		}

		for _, items := range hosts.Nodes {

			wg.Add(1)

			go func() {
				conn, err := sshconn.InitSSHConn(&items)
				if err != nil {
					log.Printf("Error to connect %v: %v\n", items.Name, err)
				}

				for _, cmd := range comm {
					stdout, stderr, err := sshconn.ExecCmd(ctx, cmd, conn)
					if err != nil {
						fmt.Printf("Error to execute command: %v\n", err)
						fmt.Printf("stderr: %v\n", stderr)
						return
					}

					fmt.Printf("OK: %s\n", stdout)

				}
				wg.Done()
			}()

		}

		wg.Wait()

	}

	if len(os.Args) == 1 {
		fmt.Println("Maybe you forget some commands, use '-help' to see available commands.")
		return
	}
}

// func runnn(h, u, k string) {

// 	fmt.Println("Executing...")

// 	err := sshconn.SSHConn(h, u, k)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	fmt.Println("Closing...")

// }

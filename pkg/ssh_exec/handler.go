package sshexec

import (
	"context"
	"fmt"
	"log"
	"slices"
	"sync"

	readfiles "github.com/lucaslimafernandes/pkg/read_files"
	"github.com/lucaslimafernandes/pkg/sshconn"
)

type TaksLog struct {
	Task   string
	Log    string
	Errors []error
}

type Log struct {
	Node string
	Task []TaksLog
}

func RunCmd(hostsPath *string, tasks *readfiles.Runfile) {

	var errors []error
	// var logs Log
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
				// logs = Log{
				// 	Node: items.Name,
				// 	Task: []TaksLog{
				// 		Task
				// 	},
				// }
			}

			for _, exec := range tasks.Tasks {

				if slices.Contains(exec.Node, items.Name) {
					stdout, stderr, err := sshconn.ExecCmd(ctx, exec.Command, conn)
					if err != nil {
						fmt.Printf("Error to execute command: %v\n", err)
						fmt.Printf("stderr: %v\n", stderr)

						errors = append(errors, fmt.Errorf("%v: %v", err, stderr))
						continue
					}

					fmt.Printf("OK: %s\n", stdout)
				}

			}

			wg.Done()
		}()

	}

	wg.Wait()

	fmt.Println("Errors occured")
	for _, e := range errors {
		fmt.Println(e)
	}

}

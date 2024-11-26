package sshexec

import (
	"context"
	"fmt"
	"log"
	"slices"
	"sync"

	readfiles "github.com/lucaslimafernandes/pkg/read_files"
	"github.com/lucaslimafernandes/pkg/sshconn"
	"github.com/lucaslimafernandes/pkg/utilities"
)

var mutex sync.Mutex

func RunCmd(hostsPath *string, tasks *readfiles.Runfile) {

	// var errors []error
	var nodesLogs []utilities.Log
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

			var logs utilities.Log

			conn, err := sshconn.InitSSHConn(&items)
			if err != nil {
				log.Printf("Error to connect %v: %v\n", items.Name, err)

				mutex.Lock()
				logs = utilities.Log{
					Node: items.Name,
					Task: append(logs.Task, utilities.TaskLog{
						Task:   "SSH Connect",
						Ok:     false,
						Errors: []error{fmt.Errorf("failed to establish connection %v: %v", items.Name, err)},
					}),
				}

				mutex.Unlock()
				return
			}

			mutex.Lock()
			logs = utilities.Log{
				Node: items.Name,
				Task: append(logs.Task, utilities.TaskLog{
					Task: "SSH Connect",
					Ok:   true,
				}),
			}
			mutex.Unlock()

			for _, exec := range tasks.Tasks {

				if slices.Contains(exec.Node, items.Name) {
					stdout, stderr, err := sshconn.ExecCmd(ctx, exec.Command, conn, tasks)
					if err != nil {
						utilities.ErrPrint(&items.Name, &exec.Name, &stderr)

						mutex.Lock()
						logs = utilities.Log{
							Node: items.Name,
							Task: append(logs.Task, utilities.TaskLog{
								Task:   exec.Name,
								Ok:     false,
								Errors: []error{fmt.Errorf("%v: %v", err, stderr)},
							}),
						}
						mutex.Unlock()

						// errors = append(errors, fmt.Errorf("%v: %v", err, stderr))
						continue
					}

					mutex.Lock()
					logs = utilities.Log{
						Node: items.Name,
						Task: append(logs.Task, utilities.TaskLog{
							Task: exec.Name,
							Ok:   true,
						}),
					}
					mutex.Unlock()
					utilities.OKPrint(&items.Name, &exec.Name, &stdout)
				}

			}

			nodesLogs = append(nodesLogs, logs)
			wg.Done()
		}()

	}

	wg.Wait()

	utilities.Finish(&nodesLogs)

}

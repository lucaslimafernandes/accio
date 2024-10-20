package utilities

import "log"

func APrint(node, task, stdout *string) {

	log.Printf("Node: %s\tTask: %s\t stdout: %s", *node, *task, *stdout)

}

package utilities

import (
	"fmt"
	"log"
	"os"

	"github.com/muesli/termenv"
)

type TaskLog struct {
	Task   string
	Ok     bool
	Errors []error
}

type Log struct {
	Node string
	Task []TaskLog
}

func styleAndLog(node, task, stdout *string, std, bgColor string) {
	output := termenv.NewOutput(os.Stdout, termenv.WithProfile(termenv.TrueColor))
	s := output.String("Task:", *task)
	styledString := s.Foreground(output.Color("#000000")).Background(output.Color(bgColor))
	log.Printf("Node: %s\t%s\t %s: %s", *node, std, styledString, *stdout)
}

// OKPrint print log with task background green
func OKPrint(node, task, stdout *string) {
	styleAndLog(node, task, stdout, "stdout", "#00ff00")
}

// WarnPrint print log with task background yellow
func WarnPrint(node, task, stdout *string) {
	styleAndLog(node, task, stdout, "stdout", "#ffff00")
}

// ErrPrint print log with task background red
func ErrPrint(node, task, stdout *string) {
	styleAndLog(node, task, stdout, "stderr", "#ff0000")
}

func Finish(nodesLogs *[]Log) {

	var thereErrs bool

	fmt.Println("\nThe execution is done")
	fmt.Printf("\n")

	for _, itemNode := range *nodesLogs {
		for _, value := range itemNode.Task {

			if !value.Ok {
				thereErrs = true
				break
			}

		}

		if thereErrs {
			fmt.Printf("Have errors occurred in: %v\n", itemNode.Node)
			for _, value := range itemNode.Task {
				if len(value.Errors) > 0 {
					fmt.Printf("%s: %s\n", Colorize(&value.Task, "red"), value.Errors)
					// fmt.Println(value.Errors)
				}
			}
		}
	}

}

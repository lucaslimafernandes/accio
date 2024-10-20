package utilities

import (
	"log"
	"os"

	"github.com/muesli/termenv"
)

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

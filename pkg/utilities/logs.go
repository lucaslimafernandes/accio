package utilities

import (
	"log"
	"os"

	"github.com/muesli/termenv"
)

func OKPrint(node, task, stdout *string) {

	output := termenv.NewOutput(os.Stdout, termenv.WithProfile(termenv.TrueColor))

	s := output.String(*task)

	styledString := s.Foreground(output.Color("#000000")).Background(output.Color("#00ff00"))

	log.Printf("Node: %s\tTask: %s\t stdout: %s", *node, styledString, *stdout)

}

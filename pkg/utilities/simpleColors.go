package utilities

import (
	"os"

	"github.com/muesli/termenv"
)

func Colorize(stdout *string, color string) termenv.Style {

	var colorChoice string

	switch color {
	case "red":
		colorChoice = "#ff0000"
	case "yellow":
		colorChoice = "#ffff00"
	default:
		colorChoice = "#00ff00" // green
	}

	output := termenv.NewOutput(os.Stdout, termenv.WithProfile(termenv.TrueColor))

	s := output.String(*stdout)
	styledString := s.Foreground(output.Color("#000000")).Background(output.Color(colorChoice))

	return styledString

}

package utilities

import (
	"log"
	"os"
	"strings"
)

func SaveOutputAsFile(filepath, text *string) {

	newFile, err := os.Create(*filepath)
	if err != nil {
		log.Fatalf("failed to create file: %v\n", err)
	}
	defer newFile.Close()

	_, err = newFile.WriteString(*text)
	if err != nil {
		log.Fatalf("failed to write in file: %v\n", err)
	}

}

func ExtractVars(text string) string {

	start := strings.Index(text, "{{") + 2
	end := strings.Index(text, "}}")

	if start < end {
		extracted := text[start:end]

		return strings.TrimSpace(extracted)
	}

	return ""
}

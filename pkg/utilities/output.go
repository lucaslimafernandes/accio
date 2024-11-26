package utilities

import (
	"log"
	"os"
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

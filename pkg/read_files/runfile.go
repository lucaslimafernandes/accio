package readfiles

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Task struct {
	Name    string   `yaml:"name"`
	Node    []string `yaml:"node"`
	Command string   `yaml:"command"`
}

type Runfile struct {
	Name  string `yaml:"name"`
	Tasks []Task `yaml:"tasks"`
}

// RootUser bool   `yaml:"as_root"` // Not implemented
// Envs // Not implemented

func ReadRunfile(path *string) (*Runfile, error) {

	var runfile Runfile

	f, err := os.ReadFile(*path)
	if err != nil {
		return &runfile, fmt.Errorf("unable to read runfile: %v", err)
	}

	err = yaml.Unmarshal(f, &runfile)
	if err != nil {
		return &runfile, fmt.Errorf("unable to parse runfile, please verify file: %v", err)
	}

	return &runfile, nil

}

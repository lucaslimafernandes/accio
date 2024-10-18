package readfiles

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type Node struct {
	Name           string `toml:"name"`
	Host           string `toml:"host"`
	User           string `toml:"user"`
	PrivateKeyPath string `toml:"private_key_path"`
}

type Hosts struct {
	Nodes []Node `toml:"nodes"`
}

func ReadHostsFile(path *string) (*Hosts, error) {

	var hosts Hosts

	f, err := os.ReadFile(*path)
	if err != nil {
		return nil, fmt.Errorf("unable to read hosts file: %v", err)
	}

	_, err = toml.Decode(string(f), &hosts)
	if err != nil {
		return nil, fmt.Errorf("unable to decode toml file: %v\nplease, check the file with the example", err)
	}

	return &hosts, nil

}

package readfiles

type Node struct {
	Name           string `toml:"name"`
	Host           string `toml:"host"`
	User           string `toml:"user"`
	PrivateKeyPath string `toml:"private_key_path"`
}

type All struct {
	Nodes []Node `toml:"nodes"`
}

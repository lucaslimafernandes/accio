package readfiles

type Task struct {
	Name    string   `yaml:"name"`
	Node    []string `yaml:"node"`
	Command string   `yaml:"command"`
}

type Runfile struct {
	Name       string `yaml:"name"`
	RemoteUser string `yaml:"remote_user"`
	Tasks      []Task `yaml:"tasks"`
}

# Welcome to Accio

From Latin, accio, meaning action

So, do something on multiple servers, that's the idea.

## Commands

* `accio -h` - Show available commands.
* `accio -hosts` - hosts path.
* `accio -run` - Runfile path.

## Project layout

```plaintext
.
├── cmd
│   └── accio
│       └── main.go
├── docs
│   ├── example.md
│   └── index.md
├── examples
│   └── simple.yaml
├── go.mod
├── go.sum
├── hosts.example.toml
├── IDEA.md
├── LICENSE
├── mkdocs.yml
├── pkg
│   ├── read_files
│   │   ├── hosts.go
│   │   ├── hosts_test.go
│   │   └── runfile.go
│   ├── sshconn
│   │   ├── commands.go
│   │   └── sshconn.go
│   ├── ssh_exec
│   │   └── handler.go
│   └── utilities
│       └── logs.go
└── README.md

```

# Welcome to Accio

From Latin, accio, meaning action

So, do something on multiple servers, that's the idea.

Accio is a tool designed to manage tasks executed on multiple remote servers via SSH. It allows you to configure and automate command execution on a variety of nodes using YAML configuration files, making remote server management more efficient and scalable.

## Features

- Define and manage SSH hosts in a simple configuration file.

- Execute tasks remotely via SSH using runfiles.

- Automate task execution on multiple servers.

- Support for environment variables and multi-node commands.

## Commands

* `accio -help` - Show available commands.
* `accio -version` - Show about Accio.
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

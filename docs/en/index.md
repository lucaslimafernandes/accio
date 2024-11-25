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
* `accio -version` - Show Accio version.
* `accio -hosts` - hosts path.
* `accio -run` - Runfile path.
* `accio -localrun` - Localhost execution.

## Project layout

```plaintext
.
├── build.sh
├── CHANGELOG.md
├── cmd
│   └── accio
│       └── main.go
├── docs
│   ├── en
│   │   ├── changelogs.md
│   │   ├── contributing.md
│   │   ├── errors.md
│   │   ├── hosts.md
│   │   ├── index.md
│   │   ├── installation.md
│   │   └── runfiles.md
│   ├── index.md
│   └── pt-BR
│       ├── changelogs.md
│       ├── contributing.md
│       ├── errors.md
│       ├── hosts.md
│       ├── index.md
│       ├── installation.md
│       └── runfiles.md
├── examples
│   ├── first_runfile.yaml
│   ├── install_docker.yaml
│   ├── multi_cmds.yaml
│   └── simple.yaml
├── go.mod
├── go.sum
├── hosts.example.toml
├── hosts.toml
├── IDEA.md
├── LICENSE
├── mkdocs.yml
├── pkg
│   ├── localexec
│   │   └── handler.go
│   ├── read_files
│   │   ├── hosts.go
│   │   ├── hosts_test.go
│   │   └── runfile.go
│   ├── sshconn
│   │   ├── commands.go
│   │   └── sshconn.go
│   ├── sshexec
│   │   └── handler.go
│   └── utilities
│       ├── logs.go
│       └── simpleColors.go
├── project.toml
└── README.md

```

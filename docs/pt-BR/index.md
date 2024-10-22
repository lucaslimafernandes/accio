# Welcome to Accio

Do Latin, accio, significa ação.

Então, faça algo em vários servidores, essa é a ideia.

## Commands

* `accio -h` - Exibe os comandos disponíveis.
* `accio -hosts` - Caminho para o arquivo de hosts.
* `accio -run` - Caminho para o arquivo de execução (Runfile).

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

# Welcome to Accio

Do Latin, accio, significa ação.

Então, faça algo em vários servidores, essa é a ideia.

Accio é uma ferramenta projetada para gerenciar tarefas executadas em vários servidores remotos via SSH. Ela permite que você configure e automatize a execução de comandos em diversos nós utilizando arquivos de configuração YAML, tornando o gerenciamento de servidores remotos mais eficiente e escalável.

## Funcionalidades

- Defina e gerencie hosts SSH em um arquivo de configuração simples.

- Execute tarefas remotamente via SSH usando runfiles.

- Automatize a execução de tarefas em vários servidores.

- Suporte a variáveis de ambiente e comandos em múltiplos nós.

## Commands

* `accio -help` - Exibe os comandos disponíveis.
* `accio -version` - Exibe sobre o Accio.
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

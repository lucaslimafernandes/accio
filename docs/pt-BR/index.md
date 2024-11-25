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
* `accio -version` - Exibe a versão do Accio.
* `accio -hosts` - Caminho para o arquivo de hosts.
* `accio -run` - Caminho para o arquivo de execução (Runfile).
* `accio -localrun` - Executa em modo local.

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

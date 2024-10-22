# Runfiles

Os arquivos de execução (runfiles) são configurações YAML usadas para definir uma série de tarefas a serem executadas em nós remotos via SSH. Esses arquivos são estruturados para incluir configurações de ambiente, nomes de tarefas, nós associados e os comandos específicos a serem executados. Abaixo, fornecemos uma explicação detalhada de como criar e estruturar esses arquivos de execução.

## Estrutura do Arquivo

Um arquivo de execução típico inclui várias seções principais:

- `runfile`: Um título ou descrição do propósito do arquivo de execução.
- `envs`: (Opcional) Variáveis de ambiente globais aplicáveis a todas as tarefas.
- `tasks`: Uma lista de tarefas, onde cada tarefa especifica um comando a ser executado em um ou mais nós.

## Examplo de Runfile

Aqui está um exemplo de um arquivo de execução simples chamado `first_runfile.yaml`:

```yaml
runfile: first_runfile

envs:
  - key: MYNAME
    value: Lucas

tasks:
  - name: hello
    node: 
      - oci1
    command: 
      echo "Hello, $MYNAME!"

```

A saída:

```plaintext
2024/10/21 22:41:02 Node: oci1	stdout	 Task: hello: Hello, Lucas!

The execution is done
```

## Detalhamento das Seções

### runfile

Este campo é usado para fins de documentação, descrevendo o arquivo de execução.

### envs

Define as variáveis de ambiente que são configuradas antes da execução de qualquer comando nas tarefas. Cada variável deve ter uma chave e um valor.

```yaml
envs:
  - key: TEST1
    value: VALUE_1

  - key: TEST2
    value: VALUE_2
```

### tasks

- `name`: Nome descritivo da tarefa para identificação.
- `node`: Lista de nós onde o comando será executado. Cada nó deve corresponder a um host configurado.
- `command`: O comando a ser executado. Pode incluir variáveis de ambiente inline, operadores de shell, etc.

```yaml
- name: hostname
  node: 
    - oci1
  command: hostname
```

## Escrevendo comandos

Os comandos devem ser válidos como comandos de shell. Eles podem incluir o uso de variáveis de ambiente, pipe, e outras funcionalidades do shell. Se um comando envolver várias instruções, separe-as com ponto e vírgula ou use um formato de várias linhas.

## Uso Avançado

Para cenários mais complexos, como instalar softwares ou executar serviços, considere dividir as tarefas em vários arquivos de execução ou usar dependências dentro das tarefas para garantir a ordem de execução.

- Comandos em várias linhas:

```yaml
command: 
  whoami ;
  hostname -I ;
  echo "$TEST1" ;
  echo "$TEST2" ;
  echo "It's working!"
```

- Múltiplas tarefas:

```yaml
tasks:
  - name: who_am_i
    node: 
      - oci1
      - oci2
    command: whoami

  - name: check_env1
    node: 
      - oci1
      - oci2
    command: echo "$TEST1"
```

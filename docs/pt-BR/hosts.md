# Arquivo de Configuração de Hosts

O arquivo de hosts é utilizado para configurar o acesso SSH a múltiplos servidores (nós). Cada bloco `[[nodes]]` no arquivo define uma conexão com um servidor específico, incluindo detalhes como o nome do nó, endereço IP, usuário e o caminho para a chave privada SSH. Abaixo estão os detalhes de cada campo e como devem ser preenchidos.

## Formato do Arquivo

O arquivo é organizado em blocos, cada um representando um nó. Cada bloco é delimitado por `[[nodes]]` e contém os seguintes campos:

- `name`: Um nome descritivo para o nó. Este identificador é usado para se referir ao servidor.
- `host`: O endereço IP ou nome do host do servidor, seguido pela porta do serviço SSH (geralmente a porta 22).
- `user`: A conta de usuário utilizada para a autenticação SSH.
- `private_key_path`: O caminho local para a chave privada SSH usada para autenticar no servidor.

## Exemplo de Configuração

Aqui está um exemplo de como dois servidores podem ser configurados no arquivo de hosts:

```toml
Copiar código
[[nodes]]
name = "oci1"
host = "10.128.0.1:22"
user = "ubuntu"
private_key_path = "/home/lucas/.ssh/id_rsa"

[[nodes]]
name = "oci2"
host = "10.128.0.2:22"
user = "ubuntu"
private_key_path = "/home/lucas/.ssh/id_rsa"
```

## Uso

Este arquivo é normalmente utilizado por scripts ou aplicações que automatizam conexões SSH a múltiplos servidores. Ele permite ao usuário configurar o acesso a todos os servidores necessários uma vez e reutilizar essa configuração de forma eficiente e segura.


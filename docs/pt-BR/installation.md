# Instalação

Você pode instalar o **Accio** de várias maneiras, dependendo da sua preferência. Abaixo estão as três principais opções:

## Instalar via Git Clone e Go Install

Este método é ideal se você deseja compilar o projeto a partir do código-fonte usando Go:

```bash
# Clone o repositório
git clone https://github.com/lucaslimafernandes/accio.git

# Altere para o diretório do projeto
cd accio

# Compile e instale usando Go
go install
```

Isso irá compilar o projeto e colocar o binário no diretório `$GOPATH/bin`, tornando-o disponível globalmente se `$GOPATH/bin` fizer parte do PATH do sistema.

## Baixar Binário Pré-Compilado

Se você preferir não compilar a partir do código-fonte, pode baixar um binário pré-compilado da página de releases do repositório. Após o download, mova o binário para `/usr/bin` (ou qualquer diretório incluído no seu PATH) para torná-lo disponível em todo o sistema:

```bash
# Baixe o binário (substitua o link pelo correto na página de releases)
wget https://github.com/lucaslimafernandes/accio/releases/download/vX.Y.Z/accio-linux-amd64

# Torne o binário executável
chmod +x accio-linux-amd64

# Mova o binário para /usr/bin para torná-lo acessível globalmente
sudo mv accio-linux-amd64 /usr/bin/accio
```
Agora, você pode executar `accio` de qualquer diretório no seu sistema.

## Instalar via Go (Direto do Repositório)

Se você já tiver o Go instalado, pode instalar o binário diretamente usando o gerenciador de pacotes do Go:

```bash
go install github.com/lucaslimafernandes/accio@latest
```

Esse comando buscará a versão mais recente do Accio, irá compilá-lo e colocará o binário no diretório `$GOPATH/bin`. Certifique-se de que `$GOPATH/bin` esteja no PATH do seu sistema.


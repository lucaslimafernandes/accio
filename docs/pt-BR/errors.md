# Erros

Este arquivo documenta as mensagens de erro comuns que podem ocorrer durante a execução dos comandos em diferentes nós remotos. Cada seção descreve o erro, suas possíveis causas e como solucioná-lo.

## Mensagens de Erro Comuns

*Erro:* `Process exited with status 127: bash: line 1: wh: command not found`

*Descrição:*

Este erro ocorre quando um comando específico não é encontrado no servidor remoto. O código de saída `127` é retornado pelo shell do Linux para indicar que o comando não está disponível no sistema.

*Exemplo de Erro:*

```plaintext
Have errors occurred in: oci2
raise_error: [failed to execute command: Process exited with status 127: bash: line 1: wh: command not found]
```

```plaintext
Have errors occurred in: oci1
raise_error: [failed to execute command: Process exited with status 127: bash: line 1: wh: command not found]
```

*Causa:*

O erro está relacionado ao uso de um comando que não está instalado ou disponível no ambiente do servidor remoto. No exemplo acima, o comando `wh` não foi encontrado, o que indica que foi digitado incorretamente ou que o binário esperado não está presente.

*Solução:*

1. *Verifique o Comando:* Certifique-se de que o comando digitado está correto. No caso acima, o comando correto poderia ser `whoami`, mas foi digitado como `wh`.

2. *Instale o Comando Faltante:* Se o comando correto estiver ausente no servidor, você pode instalá-lo. Por exemplo, para instalar `whoami`, use:

```bash
sudo apt-get install coreutils
```

3. *Corrija o Script:* Se o erro estiver em um script ou runfile, edite o script para usar o comando correto ou verifique se todos os comandos necessários estão instalados no servidor.


## Outros Erros

Se você encrontrar algum bug ou possuir sugestões de melhoreias, por favor, abra uma issue. Para contribuir, siga o[guia de contribuição](contributing.md)

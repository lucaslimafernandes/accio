# Errors

This file documents common error messages that may occur during the execution of commands on different remote nodes. Each section describes the error, its possible causes, and how to resolve it.

## Common Error Messages

*Error:* `Process exited with status 127: bash: line 1: wh: command not found`

*Description:*

This error occurs when a specific command is not found on the remote server. The exit code `127` is returned by the Linux shell to indicate that the command is not available on the system.

*Error Example:*

```plaintext
Have errors occurred in: oci2
raise_error: [failed to execute command: Process exited with status 127: bash: line 1: wh: command not found]
```

```plaintext
Have errors occurred in: oci1
raise_error: [failed to execute command: Process exited with status 127: bash: line 1: wh: command not found]
```

*Cause:*

The error is related to the use of a command that is either not installed or unavailable in the remote server environment. In the example above, the command `wh` was not found, indicating that it was either mistyped or the expected binary is missing.

*Solution:*

1. *Check the Command*: Ensure the command entered is correct. In the case above, the correct command might be `whoami`, but it was mistyped as `wh`.

2. *Install the Missing Command:* If the correct command is missing on the server, you can install it. For example, to install `whoami`, use:

```bash
sudo apt-get install coreutils
```

3. *Fix the Script:* If the error is in a script or runfile, edit the script to use the correct command or check that all necessary commands are installed on the server.


## Outros Erros

If you encounter any bugs or have feature requests, please open an issue. To contribute code, follow the [contributing guide](contributing.md)

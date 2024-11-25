# Accio

Accio is a tool designed to manage tasks executed on multiple remote servers via SSH. It allows you to configure and automate command execution on a variety of nodes using YAML configuration files, making remote server management more efficient and scalable.

## Features

- Define and manage SSH hosts in a simple configuration file.

- Execute tasks remotely via SSH using runfiles.

- Automate task execution on multiple servers.

- Support for environment variables and multi-node commands.

## Installation

You can install **Accio** in several ways, depending on your preference. Below are the three main options:

### Install via Git Clone and Go Install

This method is ideal if you want to build the project from source using Go:

```bash
# Clone the repository
git clone https://github.com/lucaslimafernandes/accio.git

# Change to the project directory
cd accio

# Build and install using Go
go install
```

This will compile the project and place the binary in your `$GOPATH/bin` directory, making it available globally if `$GOPATH/bin` is part of your system's PATH.


### Download Prebuilt Binary

If you prefer to skip building from source, you can download a prebuilt binary from the releases page of the repository. Once downloaded, move the binary to `/usr/bin` (or any directory included in your PATH) to make it available system-wide:

```bash
# Download the binary (replace the link with the actual one from the releases page)
wget https://github.com/lucaslimafernandes/accio/releases/download/vX.Y.Z/accio-linux-amd64

# Make the binary executable
chmod +x accio-linux-amd64

# Move the binary to /usr/bin to make it globally accessible
sudo mv accio-linux-amd64 /usr/bin/accio
```

Now, you can run `accio` from any directory on your system.

### Install via Go (Direct from Repository)

If you have Go installed, you can install the binary directly using Go's built-in package management:

```bash
go install github.com/lucaslimafernandes/accio@latest
```

This command will fetch the latest version of Accio, build it, and place the binary in your `$GOPATH/bin` directory. Ensure that `$GOPATH/bin` is in your system's PATH.

## Usage

Accio supports a set of commands to interact with the remote servers and manage task execution. Below are some of the key commands:

- `accio -help` - Show available commands.
- `accio -version` - Show Accio version.
- `accio -hosts` - Specify the path to the hosts configuration file.
- `accio -run` - Specify the path to the runfile for executing tasks.
- `accio -localrun` - Execute in localhost.

*Example usage:*

```bash
accio -hosts /path/to/hosts -run /path/to/runfile.yaml

accio -localrun -run /path/to/runfile.yaml
```

## Configuration

### Hosts File

The hosts file is used to define the SSH access to multiple servers. Each server is represented as a node with the following fields:

- `name:` A descriptive name for the node.

- `host:` The IP address and SSH port of the server.

- `user:` The username used for authentication.

- `private_key_path:` Path to the SSH private key.

*Example configuration:*

```toml
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

For more details, refer to the [Hosts Configuration Documentation](/docs/en/hosts.md).

### Runfile

Runfiles are used to define a sequence of tasks to be executed on the nodes. These files are written in YAML format and can include environment variables and commands for multiple nodes.

*Example runfile:*

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
For more information, check the [Runfile Documentation](/docs/en/runfiles.md).

Error Handling
Accio documents common error messages that may occur during execution. Below is an example of a common error:

```plaitext
Process exited with status 127: bash: line 1: wh: command not found
```

This error indicates that a command is not available on the remote node. To resolve it, verify the command or install the missing dependencies. More details on error handling can be found in the [Errors Documentation](/docs/en/errors.md).

## Contributing

Your contributions are welcome! If you encounter any bugs or have feature requests, please open an issue. To contribute code, follow these steps:

1. Fork the repository.

2. Clone your forked repository to your local machine.

3. Create a new branch for your feature or bugfix (git checkout -b feature-name).

4. Make your changes and commit them (git commit -m "Description of changes").

5. Push your branch (git push origin feature-name).

6. Open a pull request with a clear description of your changes.

For more details, check the [Contributing Guide](/docs/en/contributing.md).

License
This project is licensed under the MIT License. See the [LICENSE](/LICENSE) file for more information.




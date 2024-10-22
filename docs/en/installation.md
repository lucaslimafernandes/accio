# Installation

You can install **Accio** in several ways, depending on your preference. Below are the three main options:

## Install via Git Clone and Go Install

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


## Download Prebuilt Binary

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

## Install via Go (Direct from Repository)

If you have Go installed, you can install the binary directly using Go's built-in package management:

```bash
go install github.com/lucaslimafernandes/accio@latest
```

This command will fetch the latest version of Accio, build it, and place the binary in your `$GOPATH/bin` directory. Ensure that `$GOPATH/bin` is in your system's PATH.


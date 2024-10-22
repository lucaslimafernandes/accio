# Hosts Configuration File

The hosts file is used to set up SSH access to multiple servers (nodes). Each `[[nodes]]` block in the file defines a connection to a specific server, including details such as the node name, IP address, user, and the path to the SSH private key. Below are details about each field and how they should be filled out.

## File Format


The file is organized into blocks, each representing one node. Each block is delineated by `[[nodes]]` and contains the following fields:

- `name`: A descriptive name for the node. This identifier is used to reference the server.
- `host`: The IP address or hostname of the server, followed by the SSH service port (usually port 22).
- `user`: The user account used for SSH authentication.
- `private_key_path`: The local path to the SSH private key used for authenticating to the server.

## Example Configuration

Here is an example of how two servers might be configured in the hosts file:

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

## Usage

This file is typically used by scripts or applications that automate SSH connections to multiple servers. It allows the user to configure access to all necessary servers once and reuse this configuration efficiently and securely.


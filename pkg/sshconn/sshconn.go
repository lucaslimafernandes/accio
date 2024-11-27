package sshconn

import (
	"fmt"
	"os"
	"time"

	readfiles "github.com/lucaslimafernandes/pkg/read_files"
	"golang.org/x/crypto/ssh"
)

func InitSSHConn(node *readfiles.Node) (*ssh.Client, error) {

	key, err := os.ReadFile(node.PrivateKeyPath)
	if err != nil {
		return nil, fmt.Errorf("unable to read private key: %v", err)
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return nil, fmt.Errorf("unable to parse private key: %v", err)
	}

	config := &ssh.ClientConfig{
		User: node.User,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         6 * time.Hour,
	}

	return ssh.Dial("tcp", node.Host, config)

}

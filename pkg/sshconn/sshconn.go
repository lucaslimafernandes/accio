package sshconn

import (
	"fmt"
	"os"
	"time"

	readfiles "github.com/lucaslimafernandes/pkg/read_files"
	"golang.org/x/crypto/ssh"
)

// func SSHConn(host, username, keyPath string) error {

// 	key, err := os.ReadFile(keyPath)
// 	if err != nil {
// 		return fmt.Errorf("unable to read private key: %v", err)
// 	}

// 	signer, err := ssh.ParsePrivateKey(key)
// 	if err != nil {
// 		return fmt.Errorf("unable to parse private key: %v", err)
// 	}

// 	config := &ssh.ClientConfig{
// 		User: username,
// 		Auth: []ssh.AuthMethod{
// 			ssh.PublicKeys(signer),
// 		},
// 		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
// 		Timeout:         6 * time.Hour,
// 	}

// 	client, err := ssh.Dial("tcp", host, config)
// 	if err != nil {
// 		return fmt.Errorf("unable to connect: %v", err)
// 	}
// 	defer client.Close()

// 	session, err := client.NewSession()
// 	if err != nil {
// 		return fmt.Errorf("unable to create session: %v", err)
// 	}
// 	defer session.Close()

// 	// single command
// 	var b bytes.Buffer

// 	session.Stdout = &b

// 	// err = session.Run("/usr/bin/whoami")
// 	// if err != nil {
// 	// 	return fmt.Errorf("failed to run: %v", err)
// 	// }

// 	// fmt.Println(b.String())

// 	err = session.Run("/usr/bin/hostname")
// 	if err != nil {
// 		return fmt.Errorf("failed to run: %v", err)
// 	}

// 	fmt.Println(b.String())

// 	return nil

// }

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

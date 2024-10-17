package sshconn

import (
	"bytes"
	"fmt"
	"os"

	"golang.org/x/crypto/ssh"
)

func SSHConn(host, username, keyPath string) error {

	key, err := os.ReadFile(keyPath)
	if err != nil {
		return fmt.Errorf("unable to read private key: %v", err)
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return fmt.Errorf("unable to parse private key: %v", err)
	}

	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", host, config)
	if err != nil {
		return fmt.Errorf("unable to connect: %v", err)
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		return fmt.Errorf("unable to create session: %v", err)
	}
	defer session.Close()

	// single command
	var b bytes.Buffer

	session.Stdout = &b

	// err = session.Run("/usr/bin/whoami")
	// if err != nil {
	// 	return fmt.Errorf("failed to run: %v", err)
	// }

	// fmt.Println(b.String())

	err = session.Run("/usr/bin/hostname")
	if err != nil {
		return fmt.Errorf("failed to run: %v", err)
	}

	fmt.Println(b.String())

	return nil

}

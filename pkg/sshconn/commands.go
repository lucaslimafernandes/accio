package sshconn

import (
	"bytes"
	"context"
	"fmt"

	"golang.org/x/crypto/ssh"
)

func ExecCmd(ctx context.Context, cmd string, client *ssh.Client) (string, string, error) {

	session, err := client.NewSession()
	if err != nil {
		return "", "", fmt.Errorf("failed to create session: %v", err)
	}
	defer session.Close()

	var stdoutBuffer bytes.Buffer
	var stderrBuffer bytes.Buffer

	session.Stdout = &stdoutBuffer
	session.Stderr = &stderrBuffer

	err = session.Start(cmd)
	if err != nil {
		return "", "", fmt.Errorf("failed to start command: %v", err)
	}

	// channel to wait for the command to complete.
	doneChan := make(chan error)
	go func() {
		doneChan <- session.Wait()
	}()

	select {
	case <-ctx.Done(): // If the context is cancelled, try close session before exit
		session.Signal(ssh.SIGINT)
		return "", "", ctx.Err()
	case err := <-doneChan:
		if err != nil {
			return stdoutBuffer.String(), stderrBuffer.String(), fmt.Errorf("failed to execute command: %v", err)
		}
		return stdoutBuffer.String(), stderrBuffer.String(), nil
	}

}

func NewExecCmd(ctx *context.Context, client *ssh.Client, commands []string) (string, string, error) {

	var stdoutBuffer bytes.Buffer
	var stderrBuffer bytes.Buffer

	// commands := []string{"comando_1", "comando_2", "comando_3"}

	session, err := client.NewSession()
	if err != nil {
		return "", "", fmt.Errorf("failed to create session: %v", err)
	}
	defer session.Close()

	session.Stdout = &stdoutBuffer
	session.Stderr = &stderrBuffer

	// stdin pipe
	stdin, err := session.StdinPipe()
	if err != nil {
		return "", "", fmt.Errorf("failed to create sdin pipe: %v", err)
	}

	go func() {
		for _, cmd := range commands {
			fmt.Fprintln(stdin, cmd)
		}
	}()

	// Shell
	err = session.Shell()
	if err != nil {
		return "", "", fmt.Errorf("failed to create shell: %v", err)
	}

	// Wait session

	err = session.Wait()
	if err != nil {
		return "", "", fmt.Errorf("failed to wait session: %v", err)
	}

	return "", "", nil
}

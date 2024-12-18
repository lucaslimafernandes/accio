package sshconn

import (
	"bytes"
	"context"
	"fmt"

	readfiles "github.com/lucaslimafernandes/pkg/read_files"
	"golang.org/x/crypto/ssh"
)

func ExecCmd(ctx context.Context, cmd string, client *ssh.Client, task *readfiles.Runfile) (string, string, error) {

	session, err := client.NewSession()
	if err != nil {
		return "", "", fmt.Errorf("failed to create session: %v", err)
	}
	defer session.Close()

	envCmd := ""
	for _, _env := range task.Envs {
		envCmd += fmt.Sprintf("export %s='%s' ;", _env.Key, _env.Value) // Safely quote the value
	}
	cmd = envCmd + cmd

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

package localexec

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"

	readfiles "github.com/lucaslimafernandes/pkg/read_files"
)

func ExecCmd(task *readfiles.Runfile) (string, string, error) {

	var cmdStrings []string

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for _, comm := range task.Tasks {
		cmdStrings = append(cmdStrings, comm.Command)
	}

	fullCommand := strings.Join(cmdStrings, "")

	cmd := exec.CommandContext(ctx, "bash", "-c", fullCommand)

	for _, env := range task.Envs {
		cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", env.Key, env.Value))
	}

	// Incluír VarEnvs atual
	cmd.Env = append(cmd.Env, "PATH="+os.Getenv("PATH"))

	// buffer stdout e stderr
	var stdoutBuffer bytes.Buffer
	var stderrBuffer bytes.Buffer

	cmd.Stdout = &stdoutBuffer
	cmd.Stderr = &stderrBuffer

	err := cmd.Run()
	if err != nil {
		return "", "", fmt.Errorf("failed to execute command: %v", err)
	}

	return stdoutBuffer.String(), stderrBuffer.String(), nil

}

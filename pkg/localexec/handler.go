package localexec

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"

	readfiles "github.com/lucaslimafernandes/pkg/read_files"
)

func ExecCmd(cmdString string, task *readfiles.Runfile) (string, string, error) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cmd := exec.CommandContext(ctx, "bash", "-c", cmdString)

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

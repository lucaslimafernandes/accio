package localexec

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"

	readfiles "github.com/lucaslimafernandes/pkg/read_files"
	"github.com/lucaslimafernandes/pkg/utilities"
)

func ExecCmd(task *readfiles.Runfile) error {

	var logs utilities.Log
	var nodesLogs []utilities.Log

	runnerName := "localrun"

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for _, exec := range task.Tasks {
		stdout, stderr, err := execute(ctx, exec.Command, task)
		if err != nil {
			err_out := fmt.Sprintf("%v", err)
			utilities.ErrPrint(&runnerName, &exec.Name, &err_out)

			logs = utilities.Log{
				Node: runnerName,
				Task: append(logs.Task, utilities.TaskLog{
					Task:   exec.Name,
					Ok:     false,
					Errors: []error{fmt.Errorf("%v: %v", err, stderr)},
				}),
			}
			continue
		}
		utilities.OKPrint(&runnerName, &exec.Name, &stdout)
		if exec.SaveOutputAsFile != nil {
			utilities.SaveOutputAsFile(exec.SaveOutputAsFile, &stdout)
		}
	}

	nodesLogs = append(nodesLogs, logs)

	utilities.Finish(&nodesLogs)

	return nil

}

func execute(ctx context.Context, fullCommand string, task *readfiles.Runfile) (string, string, error) {

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

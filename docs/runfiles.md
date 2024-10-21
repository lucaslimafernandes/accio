# Runfiles

Runfiles are YAML configurations used to define a series of tasks to be executed on remote nodes via SSH. These files are structured to include environmental settings, task names, associated nodes, and the specific commands to be executed. Below, we provide a detailed explanation of how to create and structure these runfiles.

## File Structure
A typical runfile includes several key sections:

- `runfile`: A title or description of the runfile's purpose.
- `envs`: (Optional) Global environment variables applicable to all tasks.
- `tasks`: A list of tasks, where each task specifies a command to be executed on one or more nodes.

## Example Runfile

Here's an example of a simple runfile named first_runfile.yaml:

```yaml
runfile: first_runfile

envs:
  - key: MYNAME
    value: Lucas

tasks:
  - name: hello
    node: 
      - oci1
    command: 
      echo "Hello, $MYNAME!"

```

The output:

```plaintext
2024/10/21 22:41:02 Node: oci1	stdout	 Task: hello: Hello, Lucas!

The execution is done
```

## Sections Explained

### runfile

This field is for documentation purposes to describe the runfile.

### envs

Defines environment variables that are set before any commands are executed in the tasks. Each variable must have a key and a value.

### tasks

- `name`: Descriptive name of the task for identification.
- `node`: List of nodes where the command will be executed. Each node should correspond to a configured host.
- `command`: The command to be executed. Can include inline environment variables, shell operators, etc.

## Writing Commands

Commands should be valid shell commands. They can include environment variable usage, piping, and other shell functionalities. If a command involves multiple statements, separate them with semicolons or use a multi-line format.

## Advanced Usage

For more complex scenarios, such as installing software or running services, consider breaking down the tasks into multiple runfiles or using dependencies within tasks to ensure order of execution.



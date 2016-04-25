package commandrunner

import "os/exec"

func Run_cmd(args ...string) exec.Cmd {
	cmd := exec.Command(args[0])
	cmd.Args = args
	return *cmd
}

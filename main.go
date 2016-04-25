package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"testProject/CommandRunner"
)

func main() {
	var cmd (exec.Cmd)
	//cmd := exec.Command("ping", "-c 3", "google.com")
	arg1 := ""
	arg2 := ""
	if len(os.Args) > 0 {
		if os.Args[1] != "" {

			arg1 = os.Args[1]

			if os.Args[2] != "" {
				// nested otherwise out of range

				arg2 = os.Args[2]

			}

		}
	}


	if arg1 != "" && arg2 != "" {
		cmd = commandrunner.Run_cmd(arg1, "-c 3", arg2)
	} else {
		cmd = commandrunner.Run_cmd("ping", "-c 3", "google.com")
	}
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("in all caps: %q\n", out.String())
}

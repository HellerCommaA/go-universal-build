package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	//"strings"
)

func main() {
	//cmd := exec.Command("ping", "-c 3", "google.com")
	cmd := run_cmd("ping", "-c 3", "google.com");
	//cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("in all caps: %q\n", out.String())
}

func run_cmd (cmd_input string, ... string) *exec.Cmd {
	cmd := exec.Command(cmd_input);
	for {
	}
	return cmd;
}

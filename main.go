package main

import (
	//"bytes"
	"fmt"
	//"log"
	"os"
	//"os/exec"
	//"testProject/CommandRunner"
	"testProject/JsonProcessor"
)

// Length of REQURED program arguments
// Else program will fail
const ARG_LEN int = 2

func main() {
	//var cmd (exec.Cmd)
	//cmd := exec.Command("ping", "-c 3", "google.com")

	if len(os.Args) != ARG_LEN {
		fmt.Printf("Usage: %s jsonfile\n", os.Args[0]);
		os.Exit(1);
	}

	arg1 := os.Args[1];
	jsonprocessor.HandleFile(arg1);

	//cmd = commandrunner.Run_cmd(arg1);
	//var out bytes.Buffer;
	//cmd.Stdout = &out;
	//err := cmd.Run();
	//if err != nil {
	//	log.Fatal(err);
	//}
	//fmt.Printf(out.String());
	//fmt.Printf("\n");
}

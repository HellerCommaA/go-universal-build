package jsonprocessor

import (
	"io/ioutil"
	"fmt"
	"os"
	"encoding/json"
)

func HandleFile(inp string)  {
	file, e := ioutil.ReadFile(inp);
	if e != nil {
		fmt.Printf("File Error at %s\n", inp);
		os.Exit(1);
	}
	var jsontype jsonobject;
	json.Unmarshal(file, &jsontype);
	//fmt.Printf("%s\n", string(file));

	fmt.Printf("\n results: %v \n", jsontype);
	fmt.Printf("\n %s\n", jsontype.Build);

}

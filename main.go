package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func main() {
	flag.Parse()
	transportFileName := flag.Args()[0]

	currentDir, _ := os.Getwd()
	bytes, err := ioutil.ReadFile(currentDir + "/.direction")
	if err != nil {
		panic(err)
	}

	hostNameLine := strings.Split(string(bytes), "\n")[0]
	pathLine := strings.Split(string(bytes), "\n")[1]
	hostName := strings.Split(hostNameLine, "=")[1]
	path := strings.Split(pathLine, "=")[1]

	direction := hostName + ":" + path
	out, err := exec.Command("scp", transportFileName, direction).Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
}

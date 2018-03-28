package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Print("golamb v0.1\n\n")

	var handlerName = flag.String("handler", "main", "Handler name")
	flag.Parse()

	didBuildLambda := buildLambda(*handlerName)
	if didBuildLambda {
		createZip(*handlerName)
	}
}

func buildLambda(handlerName string) bool {

	var handlerGoFile = fmt.Sprintf("%s.go", handlerName)

	fmt.Printf("Building lambda from %s... ", handlerGoFile)

	//GOOS=linux GOARCH=amd64 go build -o main main.go
	cmd := exec.Command("go", "build", "-o", handlerName, handlerGoFile)
	env := os.Environ()
	env = append(env, "GOOS=linux")
	env = append(env, "GOARCH=amd64")
	cmd.Env = env

	return execCmdWithOutput(cmd)
}

func createZip(handlerName string) {
	zipFileName := fmt.Sprintf("%s.zip", handlerName)

	fmt.Printf("Creating zip file %s... ", zipFileName)

	cmd := exec.Command("zip", zipFileName, handlerName)

	execCmdWithOutput(cmd)
}

func execCmdWithOutput(cmd *exec.Cmd) bool {
	stdOutStderr, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println("error")
	} else {
		fmt.Println("success")
	}

	fmt.Printf("%s\n", stdOutStderr)

	return err == nil
}

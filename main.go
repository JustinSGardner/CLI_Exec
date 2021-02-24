package main

import (
	"fmt"
	"os"
	"os/exec"
	repo "github.com/JustinSGardner/CLI_Exec/utils"
	"syscall"
)

func main() {
	args := os.Args[1:]
	repo := repo.GetUrl(args)
	fmt.Println(repo)

	
    binary, lookErr := exec.LookPath("git")
    if lookErr != nil {
        panic(lookErr)
				
    }

    execArgs := []string{"git", "clone", repo}

    env := os.Environ()

    execErr := syscall.Exec(binary, execArgs, env)
    if execErr != nil {
        panic(execErr)
    }
}


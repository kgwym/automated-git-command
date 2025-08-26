package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 3 || os.Args[1] != "-m" {
		fmt.Println("Usage: yohan -m \"commit message\"")
		return
	}

	commitMessage := os.Args[2]

	// git add .
	cmdAdd := exec.Command("git", "add", ".")
	cmdAdd.Stdout = os.Stdout
	cmdAdd.Stderr = os.Stderr
	if err := cmdAdd.Run(); err != nil {
		fmt.Println("Error running git add:", err)
		return
	}

	// git commit -m "message"
	cmdCommit := exec.Command("git", "commit", "-m", commitMessage)
	cmdCommit.Stdout = os.Stdout
	cmdCommit.Stderr = os.Stderr
	if err := cmdCommit.Run(); err != nil {
		fmt.Println("Error running git commit:", err)
		return
	}

	// git push
	cmdPush := exec.Command("git", "push")
	cmdPush.Stdout = os.Stdout
	cmdPush.Stderr = os.Stderr
	if err := cmdPush.Run(); err != nil {
		fmt.Println("Error running git push:", err)
		return
	}

	fmt.Println("Changes pushed successfully!")
}

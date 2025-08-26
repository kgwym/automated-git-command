package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// Step 1: git add .
	cmdAdd := exec.Command("git", "add", ".")
	cmdAdd.Stdout = os.Stdout
	cmdAdd.Stderr = os.Stderr
	if err := cmdAdd.Run(); err != nil {
		fmt.Println("Error running git add:", err)
		return
	}

	// Step 2: get staged changes
	cmdDiff := exec.Command("git", "diff", "--staged", "--name-status")
	diffOut, err := cmdDiff.Output()
	if err != nil {
		fmt.Println("Error getting git diff:", err)
		return
	}
	diffLines := strings.Split(string(diffOut), "\n")

	if len(diffLines) == 0 || (len(diffLines) == 1 && diffLines[0] == "") {
		fmt.Println("No changes staged for commit.")
		return
	}

	// Step 3: auto-generate commit message
	addedFiles := []string{}
	modifiedFiles := []string{}
	deletedFiles := []string{}

	for _, line := range diffLines {
		parts := strings.Fields(line)
		if len(parts) < 2 {
			continue
		}
		status, file := parts[0], parts[1]
		switch status {
		case "A":
			addedFiles = append(addedFiles, file)
		case "M":
			modifiedFiles = append(modifiedFiles, file)
		case "D":
			deletedFiles = append(deletedFiles, file)
		}
	}

	msgParts := []string{}
	if len(addedFiles) > 0 {
		msgParts = append(msgParts, fmt.Sprintf("Add %s", strings.Join(addedFiles, ", ")))
	}
	if len(modifiedFiles) > 0 {
		msgParts = append(msgParts, fmt.Sprintf("Update %s", strings.Join(modifiedFiles, ", ")))
	}
	if len(deletedFiles) > 0 {
		msgParts = append(msgParts, fmt.Sprintf("Remove %s", strings.Join(deletedFiles, ", ")))
	}

	commitMessage := strings.Join(msgParts, "; ")
	if commitMessage == "" {
		commitMessage = "Update project"
	}

	fmt.Println("Auto Commit Message:", commitMessage)

	// Step 4: let user edit or confirm
	fmt.Println("Edit commit message or press Enter to keep:")
	reader := bufio.NewReader(os.Stdin)
	userInput, _ := reader.ReadString('\n')
	userInput = strings.TrimSpace(userInput)
	if userInput != "" {
		commitMessage = userInput
	}

	// Step 5: git commit
	cmdCommit := exec.Command("git", "commit", "-m", commitMessage)
	cmdCommit.Stdout = os.Stdout
	cmdCommit.Stderr = os.Stderr
	if err := cmdCommit.Run(); err != nil {
		fmt.Println("Error running git commit:", err)
		return
	}

	// Step 6: git push
	cmdPush := exec.Command("git", "push")
	cmdPush.Stdout = os.Stdout
	cmdPush.Stderr = os.Stderr
	if err := cmdPush.Run(); err != nil {
		fmt.Println("Error running git push:", err)
		return
	}

	fmt.Println("Changes pushed successfully!")
}

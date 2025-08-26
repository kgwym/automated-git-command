package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// ANSI color codes
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Cyan   = "\033[36m"
	Blue   = "\033[34m"
	Bold   = "\033[1m"
)

func main() {
	fmt.Println(Cyan + Bold + "\n=== YOHAN: Git Automation Tool ===" + Reset)

	// Step 1: git add .
	fmt.Println(Blue + "Staging all changes..." + Reset)
	cmdAdd := exec.Command("git", "add", ".")
	cmdAdd.Stdout = os.Stdout
	cmdAdd.Stderr = os.Stderr
	if err := cmdAdd.Run(); err != nil {
		fmt.Println(Red+"Error running git add:", err, Reset)
		return
	}
	fmt.Println(Green + "Changes staged successfully!" + Reset)

	// Step 2: get staged changes
	cmdDiff := exec.Command("git", "diff", "--staged", "--name-status")
	diffOut, err := cmdDiff.Output()
	if err != nil {
		fmt.Println(Red+"Error getting git diff:", err, Reset)
		return
	}
	diffLines := strings.Split(string(diffOut), "\n")

	if len(diffLines) == 0 || (len(diffLines) == 1 && diffLines[0] == "") {
		fmt.Println(Yellow + "No changes staged for commit." + Reset)
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

	fmt.Println(Cyan + Bold + "\nSuggested Commit Message:" + Reset)
	fmt.Println(Yellow + commitMessage + Reset)

	// Step 4: let user edit or confirm
	fmt.Println(Cyan + "\nEdit commit message or press Enter to keep:" + Reset)
	reader := bufio.NewReader(os.Stdin)
	userInput, _ := reader.ReadString('\n')
	userInput = strings.TrimSpace(userInput)
	if userInput != "" {
		commitMessage = userInput
	}

	// Step 5: git commit
	fmt.Println(Blue + "\nCommitting changes..." + Reset)
	cmdCommit := exec.Command("git", "commit", "-m", commitMessage)
	cmdCommit.Stdout = os.Stdout
	cmdCommit.Stderr = os.Stderr
	if err := cmdCommit.Run(); err != nil {
		fmt.Println(Red+"Error running git commit:", err, Reset)
		return
	}
	fmt.Println(Green + "Commit successful!" + Reset)

	// Step 6: git push
	fmt.Println(Blue + "\nPushing changes to remote..." + Reset)
	cmdPush := exec.Command("git", "push")
	cmdPush.Stdout = os.Stdout
	cmdPush.Stderr = os.Stderr
	if err := cmdPush.Run(); err != nil {
		fmt.Println(Red+"Error running git push:", err, Reset)
		return
	}

	fmt.Println(Green + Bold + "\n🎉 Changes pushed successfully!" + Reset)
}

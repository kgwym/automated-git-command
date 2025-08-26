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
	Magenta = "\033[35m"
	Bold   = "\033[1m"
)

func main() {
	if len(os.Args) > 1 {
		runCommandLine(os.Args[1:], bufio.NewReader(os.Stdin))
	} else {
		runInteractiveMenu()
	}
}

// Command-line mode
func runCommandLine(args []string, reader *bufio.Reader) {
	switch args[0] {
	case "commit":
		autoCommitAndPush(reader)
	case "status":
		runGitCommand("status", "📝")
	case "log":
		runGitCommand("log", "--oneline", "-5", "📜")
	case "branch":
		runGitCommand("branch", "🌿")
	case "checkout":
		if len(args) < 2 {
			fmt.Println(Red + "⚠️ Usage: yohan checkout <branch>" + Reset)
			return
		}
		runGitCommand("checkout", args[1], "🔀")
	case "pull":
		runGitCommand("pull", "⬇️")
	case "push":
		runGitCommand("push", "⬆️")
	case "merge":
		if len(args) < 2 {
			fmt.Println(Red + "⚠️ Usage: yohan merge <branch>" + Reset)
			return
		}
		runGitCommand("merge", args[1], "🔗")
	case "stash":
		runGitCommand("stash", "📦")
	case "stash-pop":
		runGitCommand("stash", "📤")
	case "tag":
		if len(args) < 2 {
			fmt.Println(Red + "⚠️ Usage: yohan tag <name>" + Reset)
			return
		}
		runGitCommand("tag", args[1], "🏷️")
	case "diff":
		runGitCommand("diff", "🔍")
	case "remote":
		runGitCommand("remote", "-v", "🌐")
	case "clone":
		if len(args) < 2 {
			fmt.Println(Red + "⚠️ Usage: yohan clone <url>" + Reset)
			return
		}
		runGitCommand("clone", args[1], "📥")
	case "reset":
		if len(args) < 2 || args[1] != "--hard" {
			fmt.Println(Red + "⚠️ Usage: yohan reset --hard" + Reset)
			return
		}
		fmt.Print(Yellow + "⚠️ Are you sure you want to reset --hard? (y/n): " + Reset)
		confirm, _ := reader.ReadString('\n')
		if strings.TrimSpace(confirm) == "y" {
			runGitCommand("reset", "--hard", "💥")
		} else {
			fmt.Println(Green + "✅ Cancelled reset." + Reset)
		}
	default:
		fmt.Println(Red + "❌ Unknown command:", args[0] + Reset)
	}
}

// Interactive menu mode
func runInteractiveMenu() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println(Cyan + Bold + "\n✨ YOHAN: All-in-One Git CLI ✨" + Reset)
		fmt.Println(Yellow + "Select an option:" + Reset)
		fmt.Println("1️⃣ Auto commit & push changes")
		fmt.Println("2️⃣ Git status")
		fmt.Println("3️⃣ Git log (last 5 commits)")
		fmt.Println("4️⃣ Git branch")
		fmt.Println("5️⃣ Git checkout <branch>")
		fmt.Println("6️⃣ Git pull")
		fmt.Println("7️⃣ Git push")
		fmt.Println("8️⃣ Git merge <branch>")
		fmt.Println("9️⃣ Git stash")
		fmt.Println("🔟 Git stash pop")
		fmt.Println("1️⃣1️⃣ Git tag <name>")
		fmt.Println("1️⃣2️⃣ Git diff")
		fmt.Println("1️⃣3️⃣ Git remote -v")
		fmt.Println("1️⃣4️⃣ Git clone <url>")
		fmt.Println("1️⃣5️⃣ Git reset --hard")
		fmt.Println("0️⃣ Exit")

		fmt.Print(Cyan + "\nEnter option: " + Reset)
		option, _ := reader.ReadString('\n')
		option = strings.TrimSpace(option)

		switch option {
		case "1":
			autoCommitAndPush(reader)
		case "2":
			runGitCommand("status", "📝")
		case "3":
			runGitCommand("log", "--oneline", "-5", "📜")
		case "4":
			runGitCommand("branch", "🌿")
		case "5":
			fmt.Print(Cyan + "Enter branch name: " + Reset)
			branch, _ := reader.ReadString('\n')
			runGitCommand("checkout", strings.TrimSpace(branch), "🔀")
		case "6":
			runGitCommand("pull", "⬇️")
		case "7":
			runGitCommand("push", "⬆️")
		case "8":
			fmt.Print(Cyan + "Enter branch to merge: " + Reset)
			branch, _ := reader.ReadString('\n')
			runGitCommand("merge", strings.TrimSpace(branch), "🔗")
		case "9":
			runGitCommand("stash", "📦")
		case "10":
			runGitCommand("stash", "📤")
		case "11":
			fmt.Print(Cyan + "Enter tag name: " + Reset)
			tag, _ := reader.ReadString('\n')
			runGitCommand("tag", strings.TrimSpace(tag), "🏷️")
		case "12":
			runGitCommand("diff", "🔍")
		case "13":
			runGitCommand("remote", "-v", "🌐")
		case "14":
			fmt.Print(Cyan + "Enter repository URL: " + Reset)
			url, _ := reader.ReadString('\n')
			runGitCommand("clone", strings.TrimSpace(url), "📥")
		case "15":
			fmt.Print(Yellow + "⚠️ Are you sure you want to reset --hard? (y/n): " + Reset)
			confirm, _ := reader.ReadString('\n')
			if strings.TrimSpace(confirm) == "y" {
				runGitCommand("reset", "--hard", "💥")
			} else {
				fmt.Println(Green + "✅ Cancelled reset." + Reset)
			}
		case "0":
			fmt.Println(Green + "👋 Bye!" + Reset)
			return
		default:
			fmt.Println(Red + "❌ Invalid option." + Reset)
		}
	}
}

// Run git commands with icon
func runGitCommand(args ...string) {
	icon := ""
	if len(args) > 0 {
		icon = args[len(args)-1] // Last argument can be an icon
		args = args[:len(args)-1]
	}

	fmt.Println(Magenta + "⚡ Running git", strings.Join(args, " "), icon + "..." + Reset)
	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println(Red+"❌ Error:", err, Reset)
	} else {
		fmt.Println(Green + "✅ Done!" + Reset)
	}
}

// Auto commit & push changes
func autoCommitAndPush(reader *bufio.Reader) {
	fmt.Println(Blue + "📥 Staging all changes..." + Reset)
	cmdAdd := exec.Command("git", "add", ".")
	cmdAdd.Stdout = os.Stdout
	cmdAdd.Stderr = os.Stderr
	if err := cmdAdd.Run(); err != nil {
		fmt.Println(Red+"❌ Error running git add:", err, Reset)
		return
	}
	fmt.Println(Green + "✅ Changes staged successfully!" + Reset)

	cmdDiff := exec.Command("git", "diff", "--staged", "--name-status")
	diffOut, err := cmdDiff.Output()
	if err != nil {
		fmt.Println(Red+"❌ Error getting git diff:", err, Reset)
		return
	}
	diffLines := strings.Split(string(diffOut), "\n")
	if len(diffLines) == 0 || (len(diffLines) == 1 && diffLines[0] == "") {
		fmt.Println(Yellow + "⚠️ No changes staged for commit." + Reset)
		return
	}

	added, modified, deleted := []string{}, []string{}, []string{}
	for _, line := range diffLines {
		parts := strings.Fields(line)
		if len(parts) < 2 {
			continue
		}
		status, file := parts[0], parts[1]
		switch status {
		case "A":
			added = append(added, file)
		case "M":
			modified = append(modified, file)
		case "D":
			deleted = append(deleted, file)
		}
	}

	msgParts := []string{}
	if len(added) > 0 {
		msgParts = append(msgParts, fmt.Sprintf("➕ Add %s", strings.Join(added, ", ")))
	}
	if len(modified) > 0 {
		msgParts = append(msgParts, fmt.Sprintf("📝 Update %s", strings.Join(modified, ", ")))
	}
	if len(deleted) > 0 {
		msgParts = append(msgParts, fmt.Sprintf("❌ Remove %s", strings.Join(deleted, ", ")))
	}
	commitMsg := strings.Join(msgParts, "; ")
	if commitMsg == "" {
		commitMsg = "🔄 Update project"
	}

	fmt.Println(Cyan + Bold + "\nSuggested Commit Message:" + Reset)
	fmt.Println(Yellow + commitMsg + Reset)

	fmt.Println(Cyan + "\nEdit commit message or press Enter to keep:" + Reset)
	userInput, _ := reader.ReadString('\n')
	userInput = strings.TrimSpace(userInput)
	if userInput != "" {
		commitMsg = userInput
	}

	fmt.Println(Blue + "\n💾 Committing changes..." + Reset)
	cmdCommit := exec.Command("git", "commit", "-m", commitMsg)
	cmdCommit.Stdout = os.Stdout
	cmdCommit.Stderr = os.Stderr
	if err := cmdCommit.Run(); err != nil {
		fmt.Println(Red+"❌ Error running git commit:", err, Reset)
		return
	}
	fmt.Println(Green + "✅ Commit successful!" + Reset)

	fmt.Println(Blue + "\n🚀 Pushing changes to remote..." + Reset)
	cmdPush := exec.Command("git", "push")
	cmdPush.Stdout = os.Stdout
	cmdPush.Stderr = os.Stderr
	if err := cmdPush.Run(); err != nil {
		fmt.Println(Red+"❌ Error running git push:", err, Reset)
		return
	}
	fmt.Println(Green + Bold + "\n🎉 Changes pushed successfully!" + Reset)
}

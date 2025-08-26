YOHAN - All-in-One Git CLI

✨ YOHAN is a colorful, emoji-friendly, interactive Git CLI built in Go. It simplifies Git operations like commit, push, pull, branch management, stashing, tagging, and more.

Features

Auto-generate commit messages based on staged changes

Editable commit messages before committing

Interactive menu mode with colorful output and icons

Full support for all common Git operations

Works on Mac, Linux, Windows terminals

Setup
1. Prerequisites

Git
 installed and configured

Go
 installed (version 1.20+)

2. Installation

Clone the repository:

git clone <repo-url>
cd yohan


Build the executable:

go build -o yohan main.go


(Optional) Move to a directory in your PATH for global usage:

# Mac / Linux
sudo mv yohan /usr/local/bin/

# Windows (PowerShell)
Move-Item .\yohan.exe C:\Windows\System32\

Usage

YOHAN can be used in two modes:

1. Interactive Menu Mode

Run the tool without arguments:

yohan


You’ll see a colorful menu with all Git commands. Choose an option and follow prompts.

2. Direct Command-Line Mode

Run specific commands directly:

yohan <command> [options]

Supported Commands
Basic Operations
Command	Description	Example
status	Show repository status	yohan status
add <files>	Stage files (. for all)	yohan add .
commit	Auto-generate commit message	yohan commit
commit -m "message"	Commit with your own message	yohan commit -m "Fix bug"
push	Push current branch to remote	yohan push
pull	Pull latest changes	yohan pull
log [n]	Show last n commits (default 5)	yohan log 10
diff	Show unstaged changes	yohan diff
reset --hard	Hard reset current branch	yohan reset --hard
Branching & Merging
Command	Description	Example
branch	List all branches	yohan branch
branch <name>	Create a new branch	yohan branch feature/login
checkout <branch>	Switch to a branch	yohan checkout develop
merge <branch>	Merge branch into current	yohan merge feature/login
Remote Operations
Command	Description	Example
remote	List remotes	yohan remote
remote add <name> <url>	Add remote	yohan remote add origin https://...
clone <url>	Clone repository	yohan clone https://github.com/user/repo.git
Stashing
Command	Description	Example
stash	Stash changes	yohan stash
stash-pop	Apply last stash	yohan stash-pop
stash list	Show all stashes	yohan stash list
Tagging
Command	Description	Example
tag	List tags	yohan tag
tag <name>	Create a new tag	yohan tag v1.0
Other Useful Commands
Command	Description	Example
clean	Remove untracked files	yohan clean
show <file>	Show file content in last commit	yohan show main.go
Auto Commit & Push

Stages all changes automatically

Generates a commit message (Add / Update / Remove)

Shows commit message for editing before commit

Pushes changes to remote

Example:

yohan commit

Color & Icon Legend

✅ Success

❌ Error

⚠️ Warning

📥 Staging changes

💾 Committing

🚀 Pushing

➕ Add files

📝 Modified files

❌ Removed files

Supported Platforms

MacOS Terminal

Linux Terminal

Windows PowerShell / CMD
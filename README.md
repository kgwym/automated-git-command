content := []byte(`# Yohan Git Automation Tool

This is a simple Go-based command-line tool to automate Git operations.

## Features
- Automatically creates a README.md if it doesn't exist
- Runs \`git add .\`
- Commits with your provided message
- Pushes changes to the remote repository

## Usage
\`\`\`bash
yohan -m "Your commit message"
\`\`\`

This will:
1. Create README.md if it doesn't exist
2. Stage all changes
3. Commit with the message you provide
4. Push to the current branch

## Setup Instructions

### Prerequisites
- Git installed
- Go installed

### Mac
1. Open Terminal.
2. Install Go:  
\`\`\`bash
brew install go
\`\`\`
3. Clone your repo and place `yohan.go` inside it.
4. Build the tool:  
\`\`\`bash
go build -o yohan yohan.go
\`\`\`
5. Move it to your PATH (optional):  
\`\`\`bash
mv yohan /usr/local/bin/
\`\`\`

### Linux
1. Open Terminal.
2. Install Go:  
\`\`\`bash
sudo apt update
sudo apt install golang-go git -y
\`\`\`
3. Clone your repo and place `yohan.go` inside it.
4. Build the tool:  
\`\`\`bash
go build -o yohan yohan.go
\`\`\`
5. Move it to your PATH (optional):  
\`\`\`bash
sudo mv yohan /usr/local/bin/
\`\`\`

### Windows
1. Install Git for Windows: [https://git-scm.com/download/win](https://git-scm.com/download/win)  
2. Install Go for Windows: [https://go.dev/dl/](https://go.dev/dl/)  
3. Open Command Prompt or PowerShell.
4. Navigate to the folder with `yohan.go`.
5. Build the tool:  
\`\`\`powershell
go build -o yohan.exe yohan.go
\`\`\`
6. Add the folder with `yohan.exe` to your PATH environment variable.

## Author
Your Name
`)

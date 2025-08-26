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

## Author
Your Name
`)

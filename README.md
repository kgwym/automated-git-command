Absolutely! Here’s a **full, descriptive, professional README** that covers **installation, setup, and usage** for **macOS, Linux, and Windows**, along with detailed explanations for all features and workflows:

---

# ✨ YOHAN: All-in-One Git CLI ✨

**YOHAN** is a comprehensive Git CLI wrapper written in **Go**.
It simplifies Git workflows, reduces repetitive commands, and provides a colorful, interactive, and intuitive interface for developers across **macOS, Linux, and Windows**.

YOHAN supports both **interactive menu-driven operations** and **direct command-line execution**, complete with **emoji-enhanced output** and **smart commit message suggestions**.

---

## 🛠 Features

1. **Interactive Menu Mode**

   * Launch a user-friendly menu with numbered options for common Git commands.
   * Ideal for beginners and users who prefer visual prompts over typing commands.

2. **Command-Line Mode**

   * Execute commands directly: `yohan <command>`
   * Perfect for automation, scripts, or fast workflow.

3. **Auto Commit & Push**

   * Stages all changes (`git add .`) automatically.
   * Suggests commit messages based on added, modified, and deleted files.
   * Allows editing the suggested message before committing.
   * Pushes commits to the remote repository.

4. **Safety Features**

   * Confirms before running destructive commands such as `git reset --hard`.
   * Prevents empty commits if no changes are staged.

5. **Emoji & Color Output**

   * Uses ANSI colors for clarity.
   * Adds emojis to indicate actions, making terminal output more readable.

6. **Supports All Common Git Commands**

   * `status`, `log`, `branch`, `checkout`, `pull`, `push`, `merge`, `stash`, `tag`, `diff`, `remote`, `clone`, and `reset`.

---

## 💻 Cross-Platform Setup

YOHAN can be installed on **macOS, Linux, and Windows**.

### Prerequisites

* **Go (>=1.20)** installed: [https://go.dev/dl/](https://go.dev/dl/)
* **Git** installed and available in PATH

  ```bash
  git --version
  ```

---

### 1️⃣ macOS Setup

1. Install Go (if not already installed):

   ```bash
   brew install go
   ```

2. Clone the repository:

   ```bash
   git clone <your-repo-url>
   cd yohan-cli
   ```

3. Build the binary:

   ```bash
   go build -o yohan
   ```

4. Move binary to system PATH:

   ```bash
   sudo mv yohan /usr/local/bin/
   ```

5. Verify installation:

   ```bash
   yohan
   ```

---

### 2️⃣ Linux Setup (Ubuntu/Debian)

1. Install Go and Git:

   ```bash
   sudo apt update
   sudo apt install -y golang-go git
   ```

2. Clone the repository:

   ```bash
   git clone <your-repo-url>
   cd yohan-cli
   ```

3. Build the binary:

   ```bash
   go build -o yohan
   ```

4. Move binary to PATH:

   ```bash
   sudo mv yohan /usr/local/bin/
   ```

5. Verify installation:

   ```bash
   yohan
   ```

---

### 3️⃣ Windows Setup

1. Install Go: [https://go.dev/dl/](https://go.dev/dl/)
   Ensure Go bin directory is in your PATH (e.g., `C:\Go\bin`)

2. Install Git: [https://git-scm.com/download/win](https://git-scm.com/download/win)

3. Open **Git Bash** and clone the repository:

   ```bash
   git clone <your-repo-url>
   cd yohan-cli
   ```

4. Build the binary:

   ```bash
   go build -o yohan.exe
   ```

5. Add binary to PATH (optional): Move `yohan.exe` to a folder in PATH (e.g., `C:\Users\<YourName>\bin`)

6. Verify:

   ```bash
   yohan
   ```

> **Tip:** Use Git Bash or Windows Terminal for full color and emoji support.

---

## 📖 Usage

### Interactive Menu Mode

Run:

```bash
yohan
```

You’ll see a menu like this:

```
✨ YOHAN: All-in-One Git CLI ✨
Select an option:
1️⃣ Auto commit & push changes
2️⃣ Git status
3️⃣ Git log (last 5 commits)
4️⃣ Git branch
5️⃣ Git checkout <branch>
...
0️⃣ Exit
```

* Enter the number of the desired command
* Follow prompts to enter branch names, tags, or confirm destructive actions

---

### Command-Line Mode

Direct commands:

```bash
yohan commit           # Auto commit & push changes
yohan status           # Show Git status
yohan log              # Last 5 commits
yohan branch           # List branches
yohan checkout dev     # Checkout branch 'dev'
yohan pull             # Pull changes from remote
yohan push             # Push changes
yohan merge main       # Merge branch
yohan stash            # Stash changes
yohan stash-pop        # Apply stash
yohan tag v1.0.0       # Create a tag
yohan diff             # Show diff
yohan remote           # Show remotes
yohan clone <url>      # Clone repo
yohan reset --hard     # Hard reset (confirmation required)
```

---

## 📝 Auto Commit & Push Workflow

1. Stages all changes: `git add .`
2. Detects Added (A), Modified (M), and Deleted (D) files
3. Generates a suggested commit message:

   ```
   ➕ Add main.go, utils.go; 📝 Update README.md; ❌ Remove old_test.go
   ```
4. Allows editing before committing
5. Executes `git commit -m "..."` and pushes automatically

---

## ⚠️ Safety Features

* `reset --hard` always asks for confirmation.
* Prevents committing if no changes are staged.

---

## 🎨 Example Terminal Output

```
📥 Staging all changes...
✅ Changes staged successfully!
Suggested Commit Message:
➕ Add main.go, utils.go; 📝 Update README.md
Edit commit message or press Enter to keep:
💾 Committing changes...
✅ Commit successful!
🚀 Pushing changes to remote...
🎉 Changes pushed successfully!
```

---

## 🤝 Contributing

* Fork the repository
* Make improvements or add new features
* Submit pull requests

Ideas:

* Add custom shortcut mode (`yohan -m "commit message"`)
* Integrate GitHub/GitLab PR creation
* Add configuration for default branch push

---

## 📜 License

MIT License © 2025 – Built with ❤️ by **Yohan Manohara**

---

If you want, I can **also create a visual “All-in-One Setup Table” for macOS, Linux, and Windows** so users can copy-paste commands directly without reading the long instructions.

Do you want me to make that table too?

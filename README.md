Here's a full **README.md** file for your Go-based Git CLI wrapper project:

---

# ✨ YOHAN: All-in-One Git CLI ✨

A colorful, interactive, and user-friendly Git CLI wrapper written in **Go**.
It simplifies common Git workflows with emojis, clear prompts, and an interactive menu system.

---

## 🚀 Features

* ✅ **Interactive Menu Mode** – Run `yohan` with no arguments to use the menu-based CLI.
* ✅ **Command-Line Mode** – Run `yohan <command>` for direct Git commands.
* ✅ **Smart Auto Commit & Push** – Suggests commit messages automatically based on changes.
* ✅ **Safety Checks** – Confirms before running dangerous operations like `git reset --hard`.
* ✅ **Emoji-enhanced Output** – Friendly, colorful command execution with clear status icons.

---

## 🛠 Installation

1. **Clone this repository**:

   ```bash
   git clone <your-repo-url>
   cd yohan-cli
   ```

2. **Build the binary**:

   ```bash
   go build -o yohan
   ```

3. **Move binary to PATH (optional)**:

   ```bash
   sudo mv yohan /usr/local/bin/
   ```

Now you can run `yohan` globally from anywhere! 🎉

---

## 📖 Usage

You can use **two modes**:

### 1️⃣ Interactive Menu Mode

Just run:

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
```

### 2️⃣ Command-Line Mode

You can run specific commands directly:

```bash
yohan commit           # Auto commit & push
yohan status           # Git status
yohan log              # Last 5 commits
yohan branch           # List branches
yohan checkout dev     # Checkout branch
yohan pull             # Git pull
yohan push             # Git push
yohan merge main       # Merge branch
yohan stash            # Stash changes
yohan stash-pop        # Apply stash
yohan tag v1.0.0       # Create a tag
yohan diff             # Show diffs
yohan remote           # Show remotes
yohan clone <url>      # Clone repo
yohan reset --hard     # Hard reset (asks for confirmation)
```

---

## 🔑 Example Workflow

1. Stage, commit, and push with **automatic commit messages**:

   ```bash
   yohan commit
   ```

   → Suggests commit messages like:

   ```
   ➕ Add file1.go; 📝 Update main.go; ❌ Remove old.go
   ```

   You can edit before confirming.

2. Quickly check status:

   ```bash
   yohan status
   ```

3. Checkout a new branch:

   ```bash
   yohan checkout feature-login
   ```

---

## ⚠️ Safety Features

* `yohan reset --hard` → Always asks for confirmation before running.
* Auto commit → Prevents empty commits if no changes are staged.

---

## 📦 Dependencies

* **Go** (>= 1.20)
* **Git** installed and available in PATH

---

## 🎨 Demo Screenshot (Terminal)

```
✨ YOHAN: All-in-One Git CLI ✨
1️⃣ Auto commit & push changes
2️⃣ Git status
3️⃣ Git log (last 5 commits)
...
```

---

## 🤝 Contributing

Feel free to fork, improve, and submit pull requests! 🚀

---

## 📜 License

MIT License © 2025 – Built with ❤️ by Yohan

---

Do you want me to also add a **"Custom Alias Setup"** section in the README so that you can run `yohan -m "msg"` like `git commit -m "msg"` (shortcut style)?

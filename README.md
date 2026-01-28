<h1 align="center">Ingin Menjadi Programmer Handal Namun Enggan Ngoding</h1>

> "Consistency is key. Even if it's automated." 🤫

**Smart Lazy Git** is a tool designed to keep your GitHub contribution graph lush and green. Whether you need to fill in the gaps from the past or ensure your future activity looks disciplined while you sleep, this tool has you covered.

It generates **empty commits** (`--allow-empty`) that count towards your GitHub contribution statistics without altering your actual codebase.

## 🚀 Features

- **Time Travel Commits:** Manually trigger commits for specific dates in the past (backdating) using the Go script.
- **Real-time Mode:** Quickly generate an activity point for the current timestamp.
- **Automated Consistency:** Includes a GitHub Actions workflow to commit automatically on a schedule (Morning & Afternoon).
- **Stealth Mode:** Uses empty commits, so your file history remains clean.

## 🛠️ How It Works

This project uses two main components:

1.  **The Time Machine (`main.go`)**: A Golang script that manipulates Git environment variables (`GIT_AUTHOR_DATE`, `GIT_COMMITTER_DATE`) to create commits at precise timestamps.
2.  **The Bot (`.github/workflows/auto-commit.yml`)**: A scheduled workflow that runs on GitHub servers to push updates while you are away.

## ⚠️ Disclaimer

This tool is for **educational and entertainment purposes only.** While a green contribution graph looks cool, actual coding skills matter more. Don't use this to deceive employers or clients. Use it to decorate your profile, not to lie about your work history.
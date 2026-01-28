package main

import (
	"fmt"
	"imphnen/data"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

func main() {
	// Select a random commit message
	selected := data.CommitMessages[rand.Intn(len(data.CommitMessages))]

	// Create an empty git commit with the current timestamp
	iso := time.Now().Format(time.RFC3339)
	cmd := exec.Command("git", "commit", "--allow-empty", "-m", selected, "--date", iso)

	// Set the environment variables for author and committer date
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, fmt.Sprintf("GIT_AUTHOR_DATE=%s", iso))
	cmd.Env = append(cmd.Env, fmt.Sprintf("GIT_COMMITTER_DATE=%s", iso))

	// Redirect output to standard output and error
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the command
	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating commit: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Auto commit created with date %s and message: \"%s\"\n", iso, selected)
}
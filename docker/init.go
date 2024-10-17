package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Create symlink if it doesn't exist
	if _, err := os.Stat("/usr/local/bin/rustdesk-api-server-pro"); os.IsNotExist(err) {
		err := os.Symlink("/app/rustdesk-api-server-pro", "/usr/local/bin/rustdesk-api-server-pro")
		if err != nil {
			fmt.Println("Error creating symlink:", err)
			os.Exit(1)
		}
	}

	// Run sync if server.db doesn't exist
	if _, err := os.Stat("/app/server.db"); os.IsNotExist(err) {
		cmd := exec.Command("/app/rustdesk-api-server-pro", "sync")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Println("Error running sync:", err)
			os.Exit(1)
		}
	}

	// Add admin user if .init.lock doesn't exist and env vars are set
	if _, err := os.Stat("/app/.init.lock"); os.IsNotExist(err) {
		adminUser := os.Getenv("ADMIN_USER")
		adminPass := os.Getenv("ADMIN_PASS")
		if adminUser != "" && adminPass != "" {
			cmd := exec.Command("/app/rustdesk-api-server-pro", "user", "add", adminUser, adminPass, "--admin")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				fmt.Println("Error adding admin user:", err)
				os.Exit(1)
			}
			if err := os.WriteFile("/app/.init.lock", []byte{}, 0644); err != nil {
				fmt.Println("Error creating .init.lock:", err)
				os.Exit(1)
			}
		}
	}

	// Start the server
	cmd := exec.Command("/app/rustdesk-api-server-pro", "start")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
}

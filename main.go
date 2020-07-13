package main

import (
	"flag"
)

// scan given a path crawls it and its subfolders
// sarching for Git repositories
func scan(path string) {
	print("scan")
}

// stats generates a niche graph of git contributions
func stats(email string) {
	print("stats")
}

func main() {
	var folder string
	var email string
	flag.StringVar(&folder, "add", "", "add a new folder to scan for Git repositories.")
	flag.StringVar(&email, "email", "your@email.com", "the email to scan.")
	flag.Parse()

	if folder != "" {
		scan(folder)
		return
	}

	stats(email)
}

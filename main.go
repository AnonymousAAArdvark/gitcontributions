package main

import "flag"

func main() {
	var folder string
	var email string
	flag.StringVar(&folder, "add", "", "add new folder to scan for git repositories")
	flag.StringVar(&email, "email", "your@email.com", "the email to scan")
	flag.Parse()

	if folder != "" {
		scan(folder)
		return
	}

	stats(email)
}
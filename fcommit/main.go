package main

import (
	"log"
	"os/exec"
)

func main() {
	c := exec.Command("git", "add", "--all")
	out, err := c.Output()
	if err != nil {
		log.Println(err)
	}
	log.Print(string(out))
	c = exec.Command("git", "commit", "-m", "init")
	out, err = c.Output()
	if err != nil {
		log.Println(err)
	}
	log.Print(string(out))
	c = exec.Command("git", "push", "origin master")
	out, err = c.Output()
	if err != nil {
		log.Println(err)
	}
	log.Print(string(out))
}

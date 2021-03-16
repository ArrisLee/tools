package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {
	flag.Parse()
	args := flag.Args()
	commitMessage := strings.Join(flag.Args()[:len(args)-1], " ")
	branch := args[len(args)-1]
	c := exec.Command("bash", "-c", "git add --all")
	log.Println("adding...")
	c.Run()
	c = exec.Command("bash", "-c", fmt.Sprintf("git commit -m '%s'", commitMessage))
	log.Println("committing...")
	c.Run()
	c = exec.Command("bash", "-c", fmt.Sprintf("git push origin %s", branch))
	log.Println("pushing to " + branch)
	c.Run()
	log.Println("done!")
}

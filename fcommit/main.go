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
	commitMsg := strings.Join(flag.Args()[:len(args)-1], " ")
	branch := args[len(args)-1]
	c := exec.Command("bash", "-c", "git add --all")
	log.Println("adding...")
	c.Run()
	c = exec.Command("bash", "-c", fmt.Sprintf("git commit -m '%s'", commitMsg))
	log.Println("committing...")
	if out, err := c.Output(); err == nil {
		log.Print(string(out))
	} else {
		log.Fatal(err)
	}
	c = exec.Command("bash", "-c", fmt.Sprintf("git push origin %s", branch))
	log.Println("pushing to " + branch)
	c.Output()
	log.Println("done!")

}

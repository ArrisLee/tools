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
	c.Run()
	c = exec.Command("bash", "-c", fmt.Sprintf("git commit -m '%s'", commitMsg))
	c.Run()
	c = exec.Command("bash", "-c", fmt.Sprintf("git push origin %s", branch))
	out, err := c.Output()
	if err != nil {
		log.Fatal(err)
	}
	log.Print(string(out))
}

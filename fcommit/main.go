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
	// fmt.Println(commitMsg)
	// fmt.Println(branch)
	// fmt.Println(fmt.Sprintf("git push origin %s", branch))
	c := exec.Command("bash", "-c", "git add --all")
	log.Println("adding...")
	c.Run()
	c = exec.Command("bash", "-c", fmt.Sprintf("git commit -m '%s'", commitMsg))
	log.Println("committing...")
	c.Run()
	c = exec.Command("bash", "-c", fmt.Sprintf("git push origin %s", branch))
	log.Println("pushing...")
	out, err := c.Output()
	if err != nil {
		log.Fatal(err)
	}
	log.Print(string(out))
}

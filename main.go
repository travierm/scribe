package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hpcloud/tail"
	. "github.com/logrusorgru/aurora"
)

func main() {
	path, failed := getPathArg()
	if failed {
		log.Fatal("No path argument given")
		os.Exit(0)
	}

	t, err := tail.TailFile(path, tail.Config{Follow: true, ReOpen: true})
	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}

	for line := range t.Lines {
		outputLog(line)
	}
}

func getPathArg() (string, bool) {
	path := os.Args[1]

	if path == "" {
		return "", true
	}

	return path, false
}

func outputLog(line *tail.Line) {
	fmt.Println(Blue(line.Text))
}

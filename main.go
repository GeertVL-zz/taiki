package main

import (
	"fmt"
	"github.com/geertvl/taiki/repl"
	"os"
)

func main() {
	fmt.Printf("Taiki Language at your service!\n")
	fmt.Printf("Enter your command\n")
	repl.Start(os.Stdin, os.Stdout)
}

package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	add    = flag.Bool("add", false, "")
	del    = flag.String("del", "", "")
	modify = flag.Int("modify", 0, "")
	list   = flag.Int("list", 999, "")
	show   = flag.String("show", "today", "")
	clean  = flag.Bool("clean", false, "")
)

func main() {
	flag.Parse()

	today := NewToday()

	// show latest file content if no parameter is passed in
	if len(os.Args) == 1 {
		today.Show()
		return
	}

	// take in one parameter at a time
	if len(os.Args) > 2 {
		fmt.Println("Please input one parameter at a time!")
		os.Exit(1)
	}

	if *add {
		today.AddPoints()
		return
	}

	fmt.Println(*clean)
	return

}

package main

import (
	"flag"
	"fmt"
	"os"
	"today/util"
)

var (
	_     = flag.Bool("add", false, "")      // done
	check = flag.String("check", "", "")     // done
	del   = flag.String("del", "", "")       // done
	mod   = flag.Int("mod", 0, "")           // todo
	_     = flag.Bool("ls", false, "")       // done
	_     = flag.Bool("clr", false, "")      // done
	ll    = flag.Int("ll", -1, "")           // done
	show  = flag.String("show", "today", "") // todo
)

func main() {
	flag.Parse()

	today := util.NewToday()
	today.LoadFile("") // load latest file

	// show the latest file content if no parameter is passed in
	if len(os.Args) == 1 {
		today.Show()
		return
	}

	// take in one parameter at a time
	if len(os.Args) > 2 {
		fmt.Println("Please input one parameter at a time...")
		os.Exit(1)
	}

	// add points to today
	if util.IsFlagPassedIn("add") {
		today.AddPoints()
		return
	}

	// check points
	if util.IsFlagPassedIn("check") {
		today.CheckPoints(*check)
		return
	}

	// delete points from today
	if util.IsFlagPassedIn("del") {
		today.DelPoints(*del)
		return
	}

	// modify one point
	if util.IsFlagPassedIn("mod") {
		today.ModifyPoint(*mod)
		return
	}

	if util.IsFlagPassedIn("clr") {
		today.Clear()
		return
	}

	if util.IsFlagPassedIn("ls") {
		today.ListFiles(-1)
		return
	}

	if util.IsFlagPassedIn("ll") {
		today.ListFiles(*ll)
		return
	}

	if util.IsFlagPassedIn("show") {
		today.ShowFile(*show)
		return
	}
}

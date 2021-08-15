package main

import (
	"flag"
	"fmt"
	"os"
	"today/util"
)

var (
	add     = flag.Bool("add", false, "")
	check   = flag.String("check", "", "")
	del     = flag.String("del", "", "")
	modify  = flag.Int("modify", 0, "")
	ls      = flag.Bool("ls", false, "")
	lsLimit = flag.Int("lsLimit", 999, "")
	show    = flag.String("show", "today", "")
	clean   = flag.Bool("clean", false, "")
)

func main() {
	flag.Parse()

	today := util.NewToday()
	today.LoadLatest()

	// show latest file content if no parameter is passed in
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
		deletePoints := *del
		fmt.Println(deletePoints)
		today.DelPoints(deletePoints)
	}

	// modify one point
	if util.IsFlagPassedIn("modify") {
		fmt.Println(*modify)
		today.ModifyPoint(*modify)
	}

	if util.IsFlagPassedIn("clean") {
		today.Clean()
	}

}

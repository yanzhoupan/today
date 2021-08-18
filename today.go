package main

import (
	"flag"
	"fmt"
	"os"
	"today/util"
)

var (
	add    = flag.Bool("add", false, "")
	check  = flag.String("check", "", "")
	del    = flag.String("del", "", "")
	modify = flag.Int("mod", 0, "")
	ls     = flag.Bool("ls", false, "")
	ll     = flag.Int("ll", -1, "")
	show   = flag.String("show", "today", "")
	clear  = flag.Bool("clr", false, "")
)

func main() {
	flag.Parse()

	today := util.NewToday()
	today.LoadLatest()

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
		today.ModifyPoint(*modify)
		return
	}

	if util.IsFlagPassedIn("clr") {
		today.Clear()
		return
	}

	if util.IsFlagPassedIn("ls") {
		util.ListFiles(-1)
		return
	}

	if util.IsFlagPassedIn("ll") {
		util.ListFiles(*ll)
		return
	}
}

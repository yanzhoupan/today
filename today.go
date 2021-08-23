package main

import (
	"flag"
	"fmt"
	"os"
	"time"
	"today/util"
)

var (
	_       = flag.Bool("add", false, "")   // done
	check   = flag.String("check", "", "")  // done
	del     = flag.String("delete", "", "") // done
	mod     = flag.Int("modify", 0, "")     // done
	history = flag.Int("history", 0,
		"List some hisroty today's, the input is the number of histories to list") // done
	_    = flag.Bool("clear", false, "")    // done
	show = flag.String("show", "today", "") // todo
)

func main() {
	flag.Parse()

	today := util.NewToday()
	today.LoadFile(time.Now().Format("2006-01-02")) // load latest file

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
	if util.IsFlagPassedIn("delete") {
		today.DelPoints(*del)
		return
	}

	// modify one point
	if util.IsFlagPassedIn("modify") {
		today.ModifyPoint(*mod)
		return
	}

	if util.IsFlagPassedIn("clear") {
		today.Clear()
		return
	}

	if util.IsFlagPassedIn("history") {
		today.ListFiles(*history)
		return
	}

	if util.IsFlagPassedIn("show") {
		today.ShowFile(*show)
		return
	}
}

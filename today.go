package main

import (
	"flag"
	"fmt"
	"os"
	"time"
	"today/util"
)

var (
	_       = flag.Bool("add", false, "Add todo points to today")         // done
	check   = flag.String("check", "", "Check points for today")          // done
	remove  = flag.String("remove", "", "Remove given points from today") // done
	mod     = flag.Int("modify", 0, "Modify a given point in today")      // done
	history = flag.Int("history", 0,
		"List some history dates, the input is the number of histories to list") // done
	_    = flag.Bool("clear", false, "clear today") // done
	show = flag.String("show", "",
		"Show the checklist for a given date, for example 'today -show=2006-01-02'") // done
	export = flag.String("export", "2021-08-01|2021-08-23",
		"Export contents in a given time range to a .txt file") // todo
	analysis = flag.String("analysis", "word_cloud|2021-08-01|2021-08-23", "Do analysis") // todo
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
		today.DelPoints(*remove)
		return
	}

	// modify one point
	if util.IsFlagPassedIn("modify") {
		today.ModifyPoint(*mod)
		return
	}

	// clear today's checklist
	if util.IsFlagPassedIn("clear") {
		today.Clear()
		return
	}

	// list history days that have checklist
	if util.IsFlagPassedIn("history") {
		today.ListFiles(*history)
		return
	}

	// show the content of a specific date
	if util.IsFlagPassedIn("show") {
		today.ShowFile(*show)
		return
	}

	// export contents in a given time range
	if util.IsFlagPassedIn("export") {
		return
	}

	// do analysis
	if util.IsFlagPassedIn("analysis") {
		return
	}
}

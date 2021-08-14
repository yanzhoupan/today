package util

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

const (
	FILE_PATH = "/home/app/github.com.yanzhoupan.today/"
)

type today struct {
	name  string
	lines []string
}

func NewToday() *today {
	return &today{
		name:  "tmp",
		lines: []string{},
	}
}

func (t *today) LoadLatest() {
	return
}

func (t *today) AddPoints() {
	fmt.Println("input your today's tasks:")
	var tasks string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		tasks = scanner.Text()
	}
	fmt.Println(tasks)
}

func (t *today) DelPoints(points string) {
	return
}

func (t *today) ModifyPoint(pointIdx int) {
	return
}

func (t *today) ListFiles(limit int) {
	return
}

// Show a given file
func Show(name string) {
	return
}

// Show current file
func (t *today) Show() {
	return
}

func Clean(name string) {
	return
}

func (t *today) Clean() {

}

func IsFlagPassedIn(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

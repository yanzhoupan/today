package util

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	// FILE_PATH = "/home/app/github.com.yanzhoupan.today/"
	FOLDER = "./app/"
	TODO   = "todo"
	DONE   = "done"
)

type today struct {
	name  string
	lines []string
}

func NewToday() *today {
	return &today{
		name:  "today",
		lines: []string{},
	}
}

func (t *today) LoadLatest() {
	if !Exists(FOLDER) {
		os.Mkdir(FOLDER, os.ModePerm)
		fmt.Println("Folder created.")
		return
	}

	// read files latest file
	fileNames := FileNames(FOLDER)
	if len(fileNames) == 0 {
		fmt.Println("Cannot load latest file, no files found...")
		os.Exit(0)
	}

	t.name = fileNames[len(fileNames)-1]
	file, err := os.Open(FOLDER + t.name)
	if err != nil {
		fmt.Printf("Open file error: %s\n", err)
		return
	}
	defer file.Close()
	br := bufio.NewReader(file)
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		lineBr := string(line) + "\n"
		t.lines = append(t.lines, lineBr)
	}
	return
}

// Add points to today
func (t *today) AddPoints() {
	fmt.Println("Input today's tasks (separate with '|'):")

	var tasks string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		tasks = scanner.Text()
	}

	t.name = time.Now().Format("2006-01-02")
	t.lines = strings.Split(tasks, "|")
	for idx, line := range t.lines {
		t.lines[idx] = fmt.Sprintf("%d) ", idx+1) + line + " |" + TODO + "\n"
	}
	t.ToFile()
}

// Check points
func (t *today) CheckPoints(s string) {
	points := strings.Split(s, ",")
	for _, i := range points {
		if idx, err := strconv.Atoi(i); err == nil {
			lineIdx := idx - 1
			if lineIdx < 0 || lineIdx > len(t.lines) {
				fmt.Println("Index out of range, point not exists.")
				continue
			}
			t.lines[lineIdx] = strings.TrimSpace(strings.Split(t.lines[lineIdx], "|")[0]) + " |" + DONE + "\n"
		} else {
			fmt.Println("Input point can not eonvert to int: ", i)
			os.Exit(1)
		}
	}
	t.ToFile()
}

// Delete points from today
func (t *today) DelPoints(points string) {
	return
}

// Modify single point
func (t *today) ModifyPoint(pointIdx int) {
	return
}

// List files
func (t *today) ListFiles(limit int) {
	fileNames := FileNames(FOLDER)
	for _, fn := range fileNames {
		fmt.Println(fn)
	}
}

// Show current file
func (t *today) Show() {

	// fmt.Println("Date: ", t.name)
	fmt.Println(fmt.Sprintf("\033[1;36m%s: %s\033[0m", "Date", t.name))
	for _, line := range t.lines {
		contents := strings.Split(line, "|")
		var task string
		var status string
		if len(contents) == 0 {
			fmt.Println("No content.")
		} else if len(contents) == 1 {
			task = contents[0]
			status = ""
		} else {
			task = contents[0]
			status = strings.TrimSpace(contents[1])
		}

		var printLine string
		if status == TODO {
			// print red
			printLine = task + fmt.Sprintf("\033[1;31m[%s]\033[0m", status)
		} else if status == DONE {
			// print green
			printLine = task + fmt.Sprintf("\033[1;32m[%s]\033[0m", status)
		} else {
			printLine = task
		}
		fmt.Println(printLine)
	}
}

// Clean all the contents of today
func (t *today) Clean() {
	t.lines = []string{""}
	t.ToFile()
}

// Write today structure to coressponding file
func (t *today) ToFile() {
	var file *os.File
	var fileErr error

	filePath := FOLDER + t.name

	file, fileErr = os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if fileErr != nil {
		fmt.Println("Error when creating file: ", fileErr)
	}
	defer file.Close()

	write := bufio.NewWriter(file)
	for idx, _ := range t.lines {
		write.WriteString(t.lines[idx])
	}
	write.Flush()
	fmt.Println("Today's tasks saved.")
	t.Show()
}

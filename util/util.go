package util

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	// FILE_PATH = "/home/app/github.com.yanzhoupan.today/"
	FOLDER = "~/.github.yanzhoupan.today/"
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

// AddPoints Add points to today
func (t *today) AddPoints() {
	fmt.Println("Input today's tasks (separate with '|'):")

	var tasks string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		tasks = scanner.Text()
	}

	t.name = time.Now().Format("2006-01-02")
	currTasksCnt := len(t.lines)
	lines := strings.Split(tasks, "|")
	for idx, line := range lines {
		t.lines = append(t.lines, fmt.Sprintf("%d) ", currTasksCnt+idx+1)+line+" |"+TODO+"\n")
	}
	t.ToFile()
}

// CheckPoints Check points
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
			fmt.Println("Input can not convert to int: ", i)
			os.Exit(1)
		}
	}
	t.ToFile()
}

// DelPoints Delete points from today
func (t *today) DelPoints(points string) {
	pts := strings.Split(points, ",")
	sort.Strings(pts)
	for _, p := range pts {
		if idx, err := strconv.Atoi(p); err == nil {
			lineIdx := idx - 1
			if lineIdx < 0 || lineIdx >= len(t.lines) {
				fmt.Println("Index out of range: ", idx)
				continue
			}
			t.lines[lineIdx] = ""
		} else {
			fmt.Println("Input can not convert to int: ", p)
			os.Exit(1)
		}
	}

	decCnt := 0
	var newLines []string
	for _, line := range t.lines {
		if line == "" {
			decCnt += 1
		} else {
			newLines = append(newLines, DescLineIndex(line, decCnt))
		}
	}

	t.lines = newLines
	if len(t.lines) == 0 {
		t.lines = []string{""}
	}

	t.ToFile()
}

// ModifyPoint Modify single point
func (t *today) ModifyPoint(pointIdx int) {
	return
}

// Show current file
func (t *today) Show() {
	fmt.Println(fmt.Sprintf("\033[1;36m%s: %s\033[0m", "Date", t.name))

	fmtLen := 0
	for _, line := range t.lines {
		if len(line) > fmtLen {
			fmtLen = len(line)
		}
	}
	fmtLen -= 4 // minus the length of status

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
			printLine = fmt.Sprintf("%-"+fmt.Sprintf("%d", fmtLen)+"s", task) + fmt.Sprintf("\033[1;31m[%s]\033[0m", status)
		} else if status == DONE {
			// print green
			printLine = fmt.Sprintf("%-"+fmt.Sprintf("%d", fmtLen)+"s", task) + fmt.Sprintf("\033[1;32m[%s]\033[0m", status)
		} else {
			printLine = fmt.Sprintf("%-20s", task)
		}
		fmt.Println(printLine)
	}
}

// Clear Clean all the contents of today
func (t *today) Clear() {
	t.lines = []string{""}
	t.ToFile()
}

// ToFile Write today structure to corresponding file
func (t *today) ToFile() {
	var file *os.File
	var fileErr error

	filePath := FOLDER + t.name

	file, fileErr = os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if fileErr != nil {
		fmt.Println("Error when creating file: ", fileErr)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	write := bufio.NewWriter(file)
	for idx := range t.lines {
		_, err := write.WriteString(t.lines[idx])
		if err != nil {
			fmt.Println("Failed to write string.")
			return
		}
	}
	err := write.Flush()
	if err != nil {
		fmt.Println("Failed to write buffer in to file.")
		os.Exit(1)
	}
	fmt.Println("Today's tasks saved.")
	t.Show()
}

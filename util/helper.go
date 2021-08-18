package util

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// Show a given file
func Show(name string) {
	return
}

// Exists Check whether a dir exists
func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// IsFlagPassedIn Check if a flag is passed in
func IsFlagPassedIn(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

// FileNames Get all file names in the given foler
func FileNames(folder string) []string {
	fileInfoList, err := ioutil.ReadDir(FOLDER)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	var fileNames []string
	for _, finfo := range fileInfoList {
		fileNames = append(fileNames, finfo.Name())
	}

	return fileNames
}

// ListFiles List limit number of files
func ListFiles(limit int) {
	fileNames := FileNames(FOLDER)
	if limit == -1 {
		limit = len(fileNames)
	}
	if len(fileNames) == 0 {
		fmt.Println("Nothing to show, please add your first note.")
		return
	}
	for idx := 0; idx < limit; idx += 1 {
		fmt.Println(fmt.Sprintf("\033[1;36m%s\033[0m", fileNames[idx]))
	}
}

func DescLineIndex(line string, descCnt int) string {
	contents := strings.SplitN(line, ")", 2)
	if idx, err := strconv.Atoi(contents[0]); err == nil {
		newIdx := idx - descCnt
		rest := contents[1]
		return fmt.Sprintf("%d", newIdx) + ") " + rest
	}
	fmt.Println("Error when convert string to int: ", contents[0])
	os.Exit(1)
	return "error"
}

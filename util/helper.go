package util

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

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

// FileNames Get all file names in the given folder
func FileNames(folder string) []string {
	fileInfoList, err := ioutil.ReadDir(folder)
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

// DescLineIndex decrease the index of a line by descCnt, it's used when points are deleted
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

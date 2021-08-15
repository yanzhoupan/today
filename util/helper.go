package util

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

// Show a given file
func Show(name string) {
	return
}

// Checker whether a dir exists
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

// Check if a flag is passed in
func IsFlagPassedIn(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

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

package main

import (
	"fmt"
	"bufio"
	"os"
)

type checklist struct {
	name string
	date string
	lines []string
}

func main(){
	fmt.Println("input your today's tasks:")
	var tasks string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan(){
		tasks = scanner.Text()
	}
	fmt.Println(tasks)
	return

}
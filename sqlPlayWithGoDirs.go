package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func makeDirs(a []string) {
	var mode uint = 0755
	for _, j := range a {
		err := os.MkdirAll(j, os.FileMode(mode))

		if err != nil {
			log.Println(err)
		}
	}

	fmt.Println("Finish creating required dirs")
}

func removeDirs(a []string) {
	for _, j := range a {
		err := os.RemoveAll(j)
		if err != nil {
			log.Println(err)
		}
	}

	fmt.Println("Finish deleting required dirs")
}

func main() {

	// Array of Required Dirs
	var requiredPaths = []string{
		"./mysql2/",
		"./mysql2/dbinit",
		"./mysql2/dbinit",
		"./mssql2/",
		"./postgres2/",
		"./postgres2/dbinit",
		"./postgres2/dbdata",
	}

	var Action = flag.String("Action", "", "(create|delete) Used to created or remove asscaited lab dir")
	flag.Parse()

	switch *Action {
	case "create":
		makeDirs(requiredPaths)
	case "delete":
		removeDirs(requiredPaths)
	default:
		fmt.Println("Invalid option. Use 'Create' or 'Delete'")
	}
}

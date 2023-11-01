package main

import (
	"bufio"
	"fmt"
	"geektrust/course"
	"geektrust/types"
	"log"
	"os"
	"strings"
)

func main() {

	defer (func() {
		recover()
	})()

	cliArgs := os.Args[1:]

	if len(cliArgs) == 0 {
		fmt.Println("Please provide the input file path")

		return
	}

	filePath := cliArgs[0]
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Error opening the input file")

		return
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	var courses []types.Course
	courseEmployeeRegMap := make(map[string]types.CourseData)
	courseRegIdMap := make(map[string]string)

	for scanner.Scan() {
		args := scanner.Text()
		argList := strings.Fields(args)
		course.ExecuteCommandsFactory(argList, &courses, courseEmployeeRegMap, courseRegIdMap)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

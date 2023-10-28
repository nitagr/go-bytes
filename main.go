package main

import (
	"bufio"
	"fmt"
	"geektrust/constants"
	"geektrust/types"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func executeCommands(commandText []string, courses *[]types.Course) {
	currentCommand, parameters := commandText[0], commandText[1:]

	switch currentCommand {
	case constants.ADD_COURSE_OFFERING:

		name, instructor, minCapacity, maxCapacity := parameters[0], parameters[1], parameters[2], parameters[3]
		minCap, _ := strconv.Atoi(minCapacity)
		maxCap, _ := strconv.Atoi(maxCapacity)
		courseId := constants.OFFERING + "-" + name + "-" + instructor
		course := types.Course{
			Id:         courseId,
			Name:       name,
			Instructor: instructor,
			Date:       time.Now(),
			MinimumCap: int32(minCap),
			MaximumCap: int32(maxCap),
		}

		*courses = append(*courses, course)
	case constants.REGISTER:
		// email, courseName := parameters[0], parameters[1]

	}
}

func main() {
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

	courses := []types.Course{}

	for scanner.Scan() {

		args := scanner.Text()
		argList := strings.Fields(args)

		executeCommands(argList, &courses)

	}
	fmt.Println("total courses", len(courses))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

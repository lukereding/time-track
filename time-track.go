package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// define struct for the projects
	type Project struct {
		Name string
	}

	// define the log file
	// get $HOME
	Home := os.Getenv("HOME")
	logFilePath := Home + "/.time-track.csv"

	// define the config file
	configFilePath := Home + "/.time-track"

	// create a log file if it doesn't exist
	_, err := os.Stat(logFilePath)

	// if it doesn't exist, create the file
	if os.IsNotExist(err) {
		var _, err = os.Create(logFilePath)
		if err != nil {
			panic("can't create the file 😥")
		}
		fmt.Println("created ", logFilePath)
	}

	// create config file if it doesn't exist
	_, err = os.Stat(configFilePath)
	if os.IsNotExist(err) {
		var _, err = os.Create(configFilePath)
		if err != nil {
			panic("can't create the file 😥")
		}
		fmt.Println("created ", configFilePath)
	}

	// parse arguments
	addProjectPtr := flag.String("add-project", "personal", "a string")
	rmProjectPtr := flag.String("rm-project", "", "a string")
	flag.Parse()

	if *addProjectPtr != "personal" {

		f, err := os.OpenFile(configFilePath, os.O_APPEND|os.O_WRONLY, 0600)

		if err != nil {
			panic(err)
		}

		defer f.Close()

		if _, err = f.WriteString(*addProjectPtr); err != nil {
			panic(err)
		}

		os.Exit(0)
	}

	if *rmProjectPtr != "" {

		input, err := ioutil.ReadFile(configFilePath)
		check(err)
		// get the lines from the file
		lines := strings.Split(string(input), "\n")

		// loop through the line, getting rid of the project to remove
		for i, line := range lines {
			if strings.Contains(line, *rmProjectPtr) {
				fmt.Printf("removing project %v", *rmProjectPtr)
				lines[i] = ""
			}
		}
		// write results back to the config file
		output := strings.Join(lines, "\n")
		err = ioutil.WriteFile(configFilePath, []byte(output), 0644)
		check(err)

		os.Exit(0)

	}

	// read in the csv file
	log, err := os.Open(logFilePath)

	if err != nil {
		panic("could not open the log file 😮")
	}
	defer log.Close()

	// read in the csv file
	reader := csv.NewReader(log)

	// create the slice of slices to hold things
	var stuff [][]string
	var index, date, project string

	// for loop for each row
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic("can't read record in the csv file")
		}

		// if it's the second entry, it's the dates
		// if it's the third, it's the activity
		for i := 0; i < len(record); i++ {
			if i == 0 {
				index = record[i]
			} else if i == 1 {
				date = record[i]
			} else if i == 2 {
				project = record[i]
			}
		}
		thisRow := []string{index, date, project}
		stuff = append(stuff, thisRow)
	}

	fmt.Println("stuff: ", stuff)
	// fmt.Println("project: ", project)

	// records, err := reader.ReadAll()
	// if err != nil {
	// 	panic("could not read in the csv file 😮")
	// }

	// fmt.Println(records)

	readerStdin := bufio.NewReader(os.Stdin)
	fmt.Printf("What are you working on? ¶ ")
	text, err := readerStdin.ReadString('\n')
	// trim whitespace
	trimmedText := strings.TrimSpace(text)

	if err != nil {
		panic("could not read from stdin")
	}

	// get the time
	currentTime := time.Now().Local().Format("2 Jan 2006 15:04")

	// append to slice
	newRow := []string{strconv.Itoa(len(stuff)), currentTime, trimmedText}
	stuff = append(stuff, newRow)
	fmt.Println(stuff)

	// write the new csv

	// working here
	// try closing 'log' then reopening it with write permissions with OpenFile
	log.Close()

	// reopen the log file
	f, err := os.OpenFile(logFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)

	// write the stuff to the csv file
	w := csv.NewWriter(f)
	w.WriteAll(stuff)

	if err := w.Error(); err != nil {
		fmt.Println(err)
	}

}

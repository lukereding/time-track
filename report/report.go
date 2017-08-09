package report

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"
)

func Report() {

	type Entry struct {
		date    string
		project string
	}

	// open the file
	Home := os.Getenv("HOME")
	logFilePath := Home + "/.time-track.csv"
	file, err := os.Open(logFilePath)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// create new csv reader
	reader := csv.NewReader(file)

	var results []Entry

	// subset to the last week
	rightNow := time.Now()
	then := rightNow.AddDate(0, 0, -7)

	fmt.Printf("using entries since %v\n\n", then)

	for {
		// read in each line
		line, err := reader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		// parse the date
		date, _ := strconv.ParseInt(line[1], 10, 64)
		tm := time.Unix(date, 0)
		formatedTime := tm.Format("Jan 2 2006")

		if tm.After(then) {
			s := Entry{date: formatedTime, project: line[2]}
			// add to results
			results = append(results, s)
		}

	}

	// find the unique projects
	var unique []string
	m := map[string]bool{}

	for _, v := range results {
		if !m[v.project] {
			m[v.project] = true
			unique = append(unique, v.project)
		}
	}

	// create map to hold unique projects and numbers of instances
	countsMap := make(map[string]int)

	// for each project
	for _, proj := range unique {
		// loop through the results
		counts := 0
		for j := 0; j < len(results); j++ {
			p := results[j]
			if p.project == proj {
				counts += 1
			}
		}
		countsMap[proj] = counts
	}

	const padding = 3
	w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', tabwriter.AlignRight|tabwriter.Debug)
	defer w.Flush()

	for key, value := range countsMap {
		fmt.Fprintln(w, key+"\t"+strings.Repeat("#", value))
	}

	// make histogram

}

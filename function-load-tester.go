package main

import (
	"encoding/csv"
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func ping(endpoint string, n int, timeout int) {
	for i := 0; i < n; i++ {
		client := http.Client{
			Timeout: time.Duration(timeout) * time.Second,
		}
		client.Get(endpoint)
	}
}

func main() {
	// Parse CSV file, -dataPath flag
	var dataPath string
	flag.StringVar(&dataPath, "dataPath", "invocations_per_function_md.anon.d01.csv", "dataPath")

	// Parse functionsCount, -functionsCount flag
	var functionsCount int
	flag.IntVar(&functionsCount, "functionsCount", 1, "functionsCount")

	// Parse timeInterval, -timeInterval flag (in seconds)
	var timeInterval int
	flag.IntVar(&timeInterval, "timeInterval", 1, "timeInverval in seconds")

	// Parse endpoint, -endpoint flag
	var endpoint string
	flag.StringVar(&endpoint, "endpoint", "http://localhost:8080/ping", "endpoint")

	// Parse timeout, -timeout flag (in seconds)
	var timeout int
	flag.IntVar(&timeout, "timeout", 10, "timeout in seconds")

	flag.Parse()

	var functionsExecuted int = 0
	var timeElapsed int = 0

	resultsFile, err := os.Create("results.csv")

	if err != nil {
		log.Fatalln("failed to open file", err)
	}

	w := csv.NewWriter(resultsFile)

	file, err := os.Open(dataPath)
	if err != nil {
		log.Fatal("Unable to read file "+dataPath, err)
	}
	defer file.Close()

	var data = make([][]string, functionsCount)
	var index int = 0

	csvReader := csv.NewReader(file)

	csvReader.Read() // Read first line

	for {
		// Read row
		record, err := csvReader.Read()

		// Stop when reached end of file
		if err == io.EOF || index == functionsCount-1 {
			break
		}

		data[index] = record

		index++
	}

	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+dataPath, err)
	}

	var column int = 4 // Start at column 4 to ignore hashes

	for {
		for i := 1; i < functionsCount-1; i++ {
			var hits = data[i][column]
			numericHits, err := strconv.Atoi(hits)
			if err != nil {
				log.Fatal("Error converting string to int ", err)
			}

			go ping(endpoint, numericHits, timeout)
			functionsExecuted += numericHits
		}

		if column > 4 {
			resultsFile.Truncate(0) // comment or uncomment
			resultsFile.Seek(0, 0)

			results := [][]string{
				{"functions_executed", strconv.Itoa(functionsExecuted)},
				{"time_elapsed", strconv.Itoa(timeElapsed)},
				{"average_functions_per_second", strconv.Itoa(functionsExecuted / timeElapsed)},
			}

			err = w.WriteAll(results)

			if err != nil {
				log.Fatal(err)
			}
		}

		time.Sleep(time.Duration(timeInterval) * time.Second)
		timeElapsed += timeInterval
		column++

		if column >= 1444 {
			break
		}
	}

	defer resultsFile.Close()
}

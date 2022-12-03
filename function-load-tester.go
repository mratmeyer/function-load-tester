package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

func ping(n int) {
	for i := 0; i < n; i++ {
		fmt.Print("ping")
		// http.Get("http://[::]:8000/")
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

	flag.Parse()

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

			go ping(numericHits)
		}

		time.Sleep(time.Duration(timeInterval) * time.Second)
		column++

		if column >= 1444 {
			break
		}
	}
}

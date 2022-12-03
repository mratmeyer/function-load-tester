package main

import (
	"encoding/csv"
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
	var filePath = "invocations_per_function_md.anon.d01.csv"

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer file.Close()

	var data [100][]string
	var index int = 0

	csvReader := csv.NewReader(file)

	csvReader.Read() // Read first line

	for {
		// Read row
		record, err := csvReader.Read()

		data[index] = record

		index++

		// Stop when reached end of file
		if err == io.EOF || index == 99 {
			break
		}
	}
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	// fmt.Println(data)

	var column int = 4

	for {
		for i := 1; i < len(data)-1; i++ {
			// fmt.Println(data[0][column])
			// fmt.Println(i, column)
			var hits = data[i][column]
			numericHits, err := strconv.Atoi(hits)
			if err != nil {
				log.Fatal("Error converting string to int ", err)
			}

			go ping(numericHits)

			// fmt.Print(hits, ",")
		}

		// fmt.Println(len(data[0]))

		time.Sleep(1 * time.Second)
		fmt.Println(column)
		column++

		if column >= 1444 {
			break
		}
	}

	// var n int = 10

	// go say("world", n)
	// say("hello", n)
}

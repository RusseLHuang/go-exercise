package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Read csv failed")
	}

	return records
}

func main() {

	limit := flag.Int("limit", 0, "limit time to answer question")
	file := flag.String("file", "problems.csv", "file of the problems")

	flag.Parse()

	reader := bufio.NewReader(os.Stdin)
	records := readCsvFile(*file)
	correctAns := 0

	fmt.Println("Press any key to start")
	reader.ReadString('\n')

	var timeout time.Duration
	timeout = time.Duration(*limit)
	timer := time.NewTimer(timeout * time.Second)

loopBreak:
	for _, question := range records {
		fmt.Println(question[0])

		ansChannel := make(chan string)

		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)

			ansChannel <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println("Total questions ", len(records))
			fmt.Println("Correctly answered ", correctAns)
			break loopBreak
		case answer := <-ansChannel:
			if answer == question[1] {
				correctAns++
			}
		}

	}

	fmt.Println("Total questions ", len(records))
	fmt.Println("Correctly answered ", correctAns)

}

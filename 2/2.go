package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isNum(line string, index int) string {
	stringNums := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	slice := line[index:]

	var toReturn string

	for key, value := range stringNums {
		if strings.HasPrefix(slice, key) {
			toReturn = value
			break
		}
	}

	return toReturn
}

func main() {
	fmt.Println("Challenge 2")
	fmt.Println("Reading inputs from file")

	file, err := os.Open("./inputs.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		firstNum := ""
		lastestNum := ""

		for index, ch := range line {
			if _, err := strconv.Atoi(fmt.Sprintf("%c", ch)); err == nil {
				if firstNum == "" {
					firstNum = fmt.Sprintf("%c", ch)
				}

				lastestNum = fmt.Sprintf("%c", ch)
			} else {
				stringNum := isNum(line, index)
				if stringNum != "" {
					if firstNum == "" {
						firstNum = stringNum
					}

					lastestNum = stringNum
				}
			}
		}

		combined, err := strconv.Atoi(firstNum + lastestNum)

		if err != nil {
			log.Fatal(err)
		}

		sum = sum + combined
	}

	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

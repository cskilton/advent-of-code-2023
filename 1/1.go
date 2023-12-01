package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Challenge 1")
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

		for _, ch := range line {
			if _, err := strconv.Atoi(fmt.Sprintf("%c", ch)); err == nil {
				if firstNum == "" {
					firstNum = fmt.Sprintf("%c", ch)
				}

				lastestNum = fmt.Sprintf("%c", ch)
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

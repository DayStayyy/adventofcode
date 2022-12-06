package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./day6/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	marker := 0
	scanner := bufio.NewScanner(file)
Exit:
	for scanner.Scan() {
		phrase := scanner.Text()
		for index, _ := range phrase {
			for i := 0; i < 4; i++ {
				if strings.Count(phrase[index:index+4], string(phrase[index+i])) != 1 {
					break
				} else if i == 3 {
					marker = index + 3
					continue Exit
				}
			}
		}
	}

	fmt.Println("Part 1 : \nThe first marker is ", marker+1)
	// fmt.Println("Part 2 : \nThe sum of the priorities is", priorities_sum_mart_two)

}

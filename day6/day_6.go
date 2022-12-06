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

	marker_one := 0
	marker_two := 0
	marker_one_is_good := true
	marker_two_is_good := true
	scanner := bufio.NewScanner(file)
Exit:
	for scanner.Scan() {
		phrase := scanner.Text()
		for index, _ := range phrase {
			marker_one_is_good = true
			marker_two_is_good = true

			if index+14 > len(phrase) {
				break
			}
			if index+3 >= len(phrase) {
				break
			}

			for i := 0; i < 14; i++ {
				if strings.Count(phrase[index:index+4], string(phrase[index+i])) != 1 {
					marker_one_is_good = false
				} else if i == 3 && marker_one == 0 && marker_one_is_good {
					marker_one = index + 3
				}

				if strings.Count(phrase[index:index+14], string(phrase[index+i])) != 1 {
					marker_two_is_good = false
					if !marker_two_is_good && !marker_one_is_good {
						break
					}
				} else if i == 13 && marker_two_is_good {
					marker_two = index + 13
					continue Exit
				}
			}
		}
	}

	fmt.Println("Part 1 : \nThe first marker is ", marker_one+1)
	fmt.Println("Part 1 : \nThe Second marker is ", marker_two+1)
}

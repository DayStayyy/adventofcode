package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./day2/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	letter_points := map[byte]int{
		'X': 1,
		'Y': 2,
		'Z': 3,
	}
	rock_paper_scissors := map[byte]map[byte]int{
		'A': {'X': 3, 'Y': 6, 'Z': 0},
		'B': {'X': 0, 'Y': 3, 'Z': 6},
		'C': {'X': 6, 'Y': 0, 'Z': 3},
	}

	points := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		points += letter_points[scanner.Text()[2]]
		points += rock_paper_scissors[scanner.Text()[0]][scanner.Text()[2]]
	}
	fmt.Println("Part 1 : \nYour number of points is ", points)
}

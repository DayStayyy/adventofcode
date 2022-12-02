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

	choices_seconde_part := map[byte]map[byte]byte{
		'A': {'X': 'Z', 'Y': 'X', 'Z': 'Y'},
		'B': {'X': 'X', 'Y': 'Y', 'Z': 'Z'},
		'C': {'X': 'Y', 'Y': 'Z', 'Z': 'X'},
	}

	points := 0
	points_second_part := 0
	choice := byte('A')
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		points += letter_points[scanner.Text()[2]]
		points += rock_paper_scissors[scanner.Text()[0]][scanner.Text()[2]]
		choice = choices_seconde_part[scanner.Text()[0]][scanner.Text()[2]]
		points_second_part += letter_points[choice]
		points_second_part += rock_paper_scissors[scanner.Text()[0]][choice]
	}
	fmt.Println("Part 1 : \nYour number of points is ", points)
	fmt.Println("Part 2 : \nYour number of points is ", points_second_part)
}

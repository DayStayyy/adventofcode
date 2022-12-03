package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./day3/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	priorities_sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		priorities_sum += calcul_priorities_sum_of_rucksack(scanner.Text()[0:len(scanner.Text())/2], scanner.Text()[len(scanner.Text())/2:])
	}

	fmt.Println("Part 1 : \nThe sum of the priorities is", priorities_sum)
}

func calcul_priorities_sum_of_rucksack(stringA, stringB string) int {
	list_of_find_letter := ""
	priorities_sum := 0
	for _, letter := range stringA {
		if strings.Contains(stringB, string(letter)) && !strings.Contains(list_of_find_letter, string(letter)) {
			priorities_sum += to_number(letter)
			list_of_find_letter += string(letter)
		}
	}
	return priorities_sum
}

func to_number(letter rune) int {
	if rune(letter) > 96 {
		return int(letter) - 96
	} else {
		return int(letter) - 38
	}
}

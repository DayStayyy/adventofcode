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
	priorities_sum_mart_two := 0
	var group []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		priorities_sum += calcul_priorities_sum_of_rucksack(scanner.Text()[0:len(scanner.Text())/2], scanner.Text()[len(scanner.Text())/2:])
		group = append(group, scanner.Text())
		if len(group) == 3 {
			priorities_sum_mart_two += common_letters(group)
			group = []string{}
		}
	}

	fmt.Println("Part 1 : \nThe sum of the priorities is", priorities_sum)
	fmt.Println("Part 2 : \nThe sum of the priorities is", priorities_sum_mart_two)

	// var list [6]string
	// list[0] = "vJrwpWtwJgWrhcsFMMfFFhFp"
	// list[1] = "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL"
	// list[2] = "PmmdzqPrVvPwwTWBwg"
	// list[3] = "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn"
	// list[4] = "ttgJtRGJQctTZtZT"
	// list[5] = "CrZsJsPPZsGzwwsLwLmpwMDw"

	// for _, word := range list {
	// 	priorities_sum += calcul_priorities_sum_of_rucksack(word[0:len(word)/2], word[len(word)/2:])
	// }

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

func common_letters(group []string) int {
	for _, letter := range group[0] {
		if strings.Contains(group[1], string(letter)) && strings.Contains(group[2], string(letter)) {
			return to_number(letter)
		}
	}
	return 0
}

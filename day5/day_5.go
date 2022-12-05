package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./day4/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	var crate [...][...]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		scanner.Text()
	}
	fmt.Println("Part 1 : \nThe number of pairs is", nb_of_pairs)
	fmt.Println("Part 1 : \nThe number of overlap is", nb_of_overlap)

}

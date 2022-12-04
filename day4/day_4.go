package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

	nb_of_pairs := 0
	nb_of_overlap := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		assignements := strings.FieldsFunc(scanner.Text(), Split)
		a1, _ := strconv.ParseInt(assignements[0], 10, 64)
		b1, _ := strconv.ParseInt(assignements[1], 10, 64)
		a2, _ := strconv.ParseInt(assignements[2], 10, 64)
		b2, _ := strconv.ParseInt(assignements[3], 10, 64)

		if (a1 <= a2 && b1 >= b2) || (a2 <= a1 && b2 >= b1) {
			nb_of_pairs++
		}
		if (a1 <= a2 && b1 >= a2) || (a2 <= a1 && b2 >= a1) {
			nb_of_overlap++
		}

	}
	fmt.Println("Part 1 : \nThe number of pairs is", nb_of_pairs)
	fmt.Println("Part 1 : \nThe number of overlap is", nb_of_overlap)

}

func Split(r rune) bool {
	return r == ',' || r == '-'
}

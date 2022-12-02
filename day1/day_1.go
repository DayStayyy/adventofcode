package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	biggest := 0
	actual := 0
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			if actual > biggest {
				biggest = actual
			}
			actual = 0
		} else {
			nb, err := strconv.ParseInt(scanner.Text(), 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			actual += int(nb)
		}
	}
	fmt.Println(biggest)
}

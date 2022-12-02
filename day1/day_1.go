package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./day1/data.txt")
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
	biggest2 := 0
	biggest3 := 0
	actual := 0
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			if actual > biggest {
				biggest, biggest2, biggest3 = actual, biggest, biggest2
			} else if actual > biggest2 {
				biggest2, biggest3 = actual, biggest2
			} else if actual > biggest3 {
				biggest3 = actual
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
	fmt.Println("First Part : \nBiggest Elf have ", biggest, " calories\nSecond Part : \nThree biggest elves have ", biggest+biggest2+biggest3, " calories")
}

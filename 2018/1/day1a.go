package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("2018/1/input1a")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	value := 0
	for scanner.Scan() {
		line := string(scanner.Text())
		i, err := strconv.Atoi(line[1:])
		if line[0] == '+' {
			if err != nil {
				log.Fatal(err)
			}
			value += i
		} else {
			value -= i
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(value)
}


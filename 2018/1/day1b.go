package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	value := 0
	dict := make(map[int]string)
	dict[int(value)] = ""
	for true {
		file, err := os.Open("2018/1/input1a")
		if err != nil {
			log.Fatal(err)
		}


		scanner := bufio.NewScanner(file)
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
			if _, exist := dict[int(value)]; exist {
				fmt.Println(value)
				fmt.Println(dict)
				fmt.Println(len(dict))
				os.Exit(0)
			}
			dict[int(value)] = ""
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		err = file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}
}


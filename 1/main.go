package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type targetValues struct {
	value1 int64
	value2 int64
}

func getTargets(target int64) (targetValues, error) {
	data, err := ioutil.ReadFile("./data.txt")
	check(err)

	items := strings.Split(string(data), "\n")
	targetNumbers := targetValues{}

	for n := 0; n < len(items); n++ {
		current := n
		if targetNumbers.value1 != 0 {
			break
		}

		for i := 0; i < len(items); i++ {
			if items[i] != items[current] {
				first, err := strconv.ParseInt(items[current], 0, 64)
				check(err)

				second, err := strconv.ParseInt(items[i], 0, 64)
				check(err)

				total := first + second

				if total == target {
					targetNumbers.value1 = first
					targetNumbers.value2 = second
					break
				}
			}
		}
	}

	if targetNumbers.value1 > 0 {
		return targetNumbers, nil
	}

	return targetNumbers, errors.New("failed")
}

func main() {
	targetNumbers, err := getTargets(2020)
	check(err)

	fmt.Println(targetNumbers.value1 * targetNumbers.value2)
}

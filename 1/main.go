package main

import (
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

func sum(slice []int64) int64 {
	total := int64(0)
	for i := 0; i < len(slice); i++ {
		total += slice[i]
	}

	return total
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func getData() []int {
	data, err := ioutil.ReadFile("./data.txt")
	check(err)

	items := strings.Split(string(data), "\n")

	intData := make([]int, len(items))

	for i := 0; i < len(items); i++ {
		num, err := strconv.ParseInt(items[i], 0, 0)
		check(err)

		intData = append(intData, int(num))
	}

	return intData
}

func part1(data []int) {
	checkedValues := make([]int, len(data))

	for i := 0; i < len(data); i++ {
		diff := 2020 - data[i]

		if contains(checkedValues, diff) {
			fmt.Println(diff * data[i])
		}

		checkedValues = append(checkedValues, data[i])
	}
}

func part2(data []int) {

}

func main() {
	data := getData()
	part1(data)
}

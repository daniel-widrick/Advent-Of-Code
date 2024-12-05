package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	fileName := "input.txt"
	file, err := os.Open(fileName)
	if err != nil {
		log.Panicf("Error opening file: %s, %v", fileName, err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	safeCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(line, " ")
		var numbers []int

		for _, str := range values {
			num, err := strconv.Atoi(str)
			if err != nil {
				log.Panicf("Unable to parse int: %s :: %s, %v",line,str,err)
			}
			numbers = append(numbers, num)
		}
		if isTolerable(numbers) {
			safeCount++
		}
		
	}
	log.Printf("Found %d safe reports",safeCount)

}


func isTolerable(values []int) bool {
	if isSafe(values) {
		return true
	}
	for i := 0; i < len(values); i++ {
		var newValues []int
		for j := 0; j < len(values); j++ {
			if i != j {
				newValues = append(newValues, values[j])
			}
		}
		if isSafe(newValues) {
			return true
		}
	}
	return false
}
func isSafe(values []int) bool {
	increments := 0
	decrements := 0

	for i := 1; i < len(values); i++ {
		diff := values[i-1] - values[i]
		abs := math.Abs(float64(diff))
		if !(abs >=1 && abs <=3) {
			return false
		}
		if diff < 0 {
			increments++
		} else if diff > 0 {
			decrements++
		}
	}
	if increments > 0 && decrements > 0 {
		return false
	}
	return true
}

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
		increments := 0
		decrements := 0
		safe := true
		for i := 1; i < len(values); i++ {
			v1, err := strconv.Atoi(values[i-1])
			if err != nil {
				log.Panicf("Error parsing int: %s, %v", line, err)
			}
			v2, err := strconv.Atoi(values[i])
			if err != nil {
				log.Panicf("Error parsing int: %s, %v", line, err)
			}
			diff := v1 - v2
			abs := math.Abs(float64(diff))
			if abs != 1 && abs != 2 && abs != 3 {
				log.Printf("%-30s\tdeviation too far\tNOT SAFE", line)
				safe = false
				break
			}
			if diff < 0 {
				decrements++
			} else if diff > 0 {
				increments++
			}
		}
		if increments > 0 && decrements > 0 {
			log.Printf("%-30s\tIncrease and Decrease\tNOT SAFE", line)
			continue
		} else if safe {
			log.Printf("%-30s\t\t\t\tSAFE!!!", line)
			safeCount++
		}
	}
	log.Printf("Found %d safe reports!", safeCount)
}

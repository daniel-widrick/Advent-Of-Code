package main

import (
	"bufio"
	"log"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func main(){
	fileName := "input.txt"
	file, err := os.Open(fileName)
	if err != nil {
		log.Panicf("error opening %s, %v",fileName, err)
	}

	ruleMap := make(map[string][]string)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break; //end of rules
		}
		parts := strings.Split(line,"|")
		ruleMap[parts[0]] = append(ruleMap[parts[0]], parts[1])
	}
	
	middleSumOrdered := 0
	middleSumFixed := 0

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		sorted := make([]string, len(parts))
		copy(sorted, parts)
		log.Printf("line: %v", parts)
		localRuleGraph := make(map[string][]string)

		for _, part1 := range parts {
			for _, part2 := range parts {
				_, exists := ruleMap[part1]
				if exists && slices.Contains(ruleMap[part1], part2) {
					localRuleGraph[part1] = append(localRuleGraph[part1], part2)
				}
			}
		}
		sort.Slice(sorted, func(i, j int) bool {
			return len(localRuleGraph[sorted[i]]) > len(localRuleGraph[sorted[j]])
		})
		log.Printf("Sort: %v", sorted)
		log.Printf("%v", localRuleGraph)
		middle, err := strconv.Atoi(sorted[len(sorted)/2])
		if err != nil {
			log.Printf("Error parsing int: %v, %v", sorted, err)
		}
		if slices.Equal(parts,sorted) {
			middleSumOrdered += middle
		} else {
			middleSumFixed += middle
		}
	}
	log.Printf("Total of middles part 1: %d", middleSumOrdered)
	log.Printf("Total of middles part 2: %d", middleSumFixed)
}

package main

import (
	"bufio"
	"log"
	"os"
)

type mapObject struct {
	Passable bool
	Visited bool
}

type Guard struct {
	x int
	y int
	xVel int
	yVel int
}

func main() {
	fileName := "input.txt"
	file, err := os.Open(fileName)
	if err != nil {
		log.Panicf("Unable to open file: %s, %v",fileName,err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	
	visitCount := 0
	x := 0
	y := 0
	guard := Guard{
		x: 0,
		y: 0,
		yVel: -1,
		xVel: 0,
	}

	lineCount := 1
	lineLen := 0
	if scanner.Scan() {
		lineLen = len(scanner.Text())
		for scanner.Scan() {
			lineCount++
		}
		file.Seek(0,0)
	}
	guardMap := make([][]mapObject,lineLen)
	for i := range guardMap {
		guardMap[i] = make([]mapObject,lineCount)
	}
	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for x = 0; x < len(line); x++ {
			char := line[x]
			switch char {
				case '.':
					guardMap[x][y] = mapObject{
						Passable: true,
						Visited: false,
					}
				case '#':
					guardMap[x][y] = mapObject{
						Passable: false,
						Visited: false,
					}
				case '^':
					guardMap[x][y] = mapObject{
						Passable: true,
						Visited: true,
					}
					guard.x = x
					guard.y = y
					log.Printf("Found guard at: %d, %d",guard.x, guard.y)
					visitCount += 1
				default:
					log.Panicf("Found unknown tile %d,%d,%s",x,y,string(char))
			}
		}
		y += 1
	}
	log.Printf("map:\n%v",guardMap)
	log.Printf("Map Bounds: %d, %d", len(guardMap), lineCount)

	for true {
		guard.x += guard.xVel
		guard.y += guard.yVel
		if guard.x < 0 || guard.y < 0 || guard.x >= lineLen || guard.y >= lineCount {
			log.Printf("Exiting map at: %d,%d", guard.x, guard.y)
			log.Printf("Bounds %d, %d", len(guardMap), lineLen)
			break; //Exited map
		}
		if !guardMap[guard.x][guard.y].Passable {
			//Path is blocked roll back... rotate and repeat
			guard.x -= guard.xVel
			guard.y -= guard.yVel
			if guard.xVel == 1 {
				//turn from right to down
				//log.Printf("Turn Down %d, %d",guard.x, guard.y)
				guard.xVel = 0
				guard.yVel = 1
			} else if guard.yVel == 1 {
				//log.Printf("Turn Left %d, %d",guard.x, guard.y)
				//trun from down to left
				guard.xVel = -1
				guard.yVel = 0
			} else if guard.xVel == -1 {
				//log.Printf("Turn Up %d, %d",guard.x, guard.y)
				//turn from left to up
				guard.xVel = 0;
				guard.yVel = -1
			} else if guard.yVel == -1 {
				//log.Printf("Turn Right%d, %d",guard.x, guard.y)
				//turn from up to right
				guard.xVel = 1
				guard.yVel = 0
			} else {
				//log.Panicf("Unknown guard state: %v",guard)
			}
			continue; //Try again in new direction
		} else if guardMap[guard.x][guard.y].Passable && !guardMap[guard.x][guard.y].Visited {
			//log.Printf("Visited new tile: %d, %d",guard.x, guard.y)
			copyObj := guardMap[guard.x][guard.y]
			copyObj.Visited = true
			guardMap[guard.x][guard.y] = copyObj
			visitCount += 1
		} else {
			//log.Printf("Weird tile: %d, %d", guard.x, guard.y)
		}
	}
	log.Printf("Visited tiles: %d", visitCount)

}

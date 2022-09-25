package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/AllOtherUserNamesTaken/automatic-invention/automaticinvention"
)

func PtOrPts(score int64) string {
	if score == 1 {
		return "pt"
	}
	return "pts"
}

func main() {
	path := os.Args[1]
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	counter := new(automaticinvention.Counter).Init()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		game := automaticinvention.GameFromString(scanner.Text()).RankedGame()
		counter.Update(game)
	}

	array := counter.AsArray()
	automaticinvention.By(automaticinvention.SortByPerformance).Sort(array)

	for i, p := range array {
		fmt.Printf("%v. %v, %v %v\n", i+1, p.Team, p.Score, PtOrPts(p.Score))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

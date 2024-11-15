package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	filename := flag.String("csv", "problems.csv", "a csv for Qns and Ans")
	flag.Parse()

	file, err := os.Open(*filename)

	if err != nil {
		fmt.Println("Failed to open the provided csv file.")
		os.Exit(1)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	lines, err := reader.ReadAll()

	if err != nil {
		fmt.Println("Failed to parse the provided csv file.")
		os.Exit(1)
	}

	problems := Parseline(lines)

	counter := 0
	for i, prob := range problems {
		fmt.Printf("Problem is #%d: %s = ", i+1, prob.que)
		var answer string
		fmt.Scanf("%s", &answer)

		if answer == prob.ans {
			counter++
			fmt.Println("Correct Answer")
		} else {
			fmt.Println("Wrong Answer")
		}
	}
	fmt.Printf("you have scored %d out of %d\n", counter, len(problems))
}

type problem struct {
	que string
	ans string
}

func Parseline(lines [][]string) []problem {
	result := make([]problem, len(lines[1])+1)

	for i, line := range lines {
		result[i] = problem{
			que: line[0],
			ans: line[1],
		}
	}
	return result
}

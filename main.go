package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Input grades, onlyCount, maxMark, and weight
	var grades []int
	var onlyCount, maxMark int
	var weight float64

	fmt.Println("Enter your grades separated by spaces:")
	fmt.Println("Like 10 10 1 etc")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	gradeInput := scanner.Text()
	gradesStr := strings.Fields(gradeInput)
	for _, gradeStr := range gradesStr {
		grade, err := strconv.Atoi(gradeStr)
		if err != nil {
			fmt.Println("Invalid input. Please enter integers separated by spaces.")
			return
		}
		grades = append(grades, grade)
	}

	fmt.Println("Enter the number of maximum possible grades:")
	fmt.Scan(&onlyCount)

	fmt.Println("Enter the maximum mark for each grade:")
	fmt.Scan(&maxMark)

	fmt.Println("Enter the weight of this section of the module:")
	fmt.Scan(&weight)

	// Sort the grades in descending order and select the top amount defined by onlyCount
	sort.Slice(grades, func(i, j int) bool {
		return grades[i] > grades[j]
	})
	var sortedGrades []int
	if len(grades) < onlyCount {
		sortedGrades = grades
	} else {
		sortedGrades = grades[:onlyCount]
	}

	// Calculate the total marks obtained
	totalMarks := 0
	for _, grade := range sortedGrades {
		totalMarks += grade
	}

	// Calculate the maximum possible marks
	maxPossibleMarks := maxMark * onlyCount

	// Calculate the percentage achieved
	percentageAchieved := (float64(totalMarks) / float64(maxPossibleMarks)) * 100

	// Calculate the weighted percentage achieved for this section
	weightedPercentageAchieved := percentageAchieved * weight

	fmt.Printf("Percentage achieved for this section: %.2f%%\n", weightedPercentageAchieved)
}
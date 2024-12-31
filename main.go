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
	var grades []float64
	var onlyCount int
	var maxMark float64
	var weight float64

	fmt.Println("Enter your grades separated by spaces:")
	fmt.Println("Like 10 10.5 1 etc")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	gradeInput := scanner.Text()
	gradesStr := strings.Fields(gradeInput)
	for _, gradeStr := range gradesStr {
		grade, err := strconv.ParseFloat(gradeStr, 64)
		if err != nil {
			fmt.Println("Invalid input. Please enter numbers (integers or floats) separated by spaces.")
			return
		}
		grades = append(grades, grade)
	}

	fmt.Println("Enter the number of maximum possible grades:")
	fmt.Scan(&onlyCount)

	fmt.Println("Enter the maximum mark for each grade:")
	fmt.Scan(&maxMark)

	fmt.Println("Enter the weight of this section of the module (as a float, e.g., 0.2):")
	scanner.Scan() // Reuse the scanner to capture weight as text
	weightInput := scanner.Text()
	parsedWeight, err := strconv.ParseFloat(strings.TrimSpace(weightInput), 64)
	if err != nil {
		fmt.Println("Invalid input for weight. Please enter a decimal number.")
		return
	}
	weight = parsedWeight

	// Sort the grades in descending order and select the top amount defined by onlyCount
	sort.Slice(grades, func(i, j int) bool {
		return grades[i] > grades[j]
	})
	var sortedGrades []float64
	if len(grades) < onlyCount {
		sortedGrades = grades
	} else {
		sortedGrades = grades[:onlyCount]
	}

	// Calculate the total marks obtained
	totalMarks := 0.0
	for _, grade := range sortedGrades {
		totalMarks += grade
	}

	// Calculate the maximum possible marks
	maxPossibleMarks := maxMark * float64(onlyCount)

	// Calculate the percentage achieved
	percentageAchieved := (totalMarks / maxPossibleMarks) * 100

	// Calculate the weighted percentage achieved for this section
	weightedPercentageAchieved := percentageAchieved * weight

	fmt.Printf("Percentage achieved for this section: %.2f%%\n", weightedPercentageAchieved)
}
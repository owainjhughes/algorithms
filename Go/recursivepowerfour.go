package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// get number of cases
	scanner.Scan()
	n, err := strconv.Atoi(scanner.Text())
	if err != nil || n < 1 || n > 100 {
		return
	}

	results := make([]string, 0, n)

	count := 0
	for count < n && scanner.Scan() {
		//get number of integers
		x, err := strconv.Atoi(scanner.Text())
		if err != nil || x <= 0 || x > 100 {
			continue
		}

		// get the integers
		scanner.Scan()
		numStr := scanner.Text()
		numbers := strings.Fields(numStr)

		// make sure its the correct number of integers
		if len(numbers) != x {
			results = append(results, "-1")
			count++
			continue
		}

		// fun calculation function with the sum and numbers
		sum := 0
		processNumbers(&sum, numbers)

		results = append(results, strconv.Itoa(sum))
		count++
	}

	printResults(results)
}

func processNumbers(sum *int, numbers []string) {
	index := 0
	processNumbersRecursive(sum, numbers, &index)
}

func processNumbersRecursive(sum *int, numbers []string, index *int) {
	// base case
	if *index >= len(numbers) {
		return
	}

	// process current int
	num, err := strconv.Atoi(numbers[*index])
	if err == nil && num <= 0 {
		// calculate power of four only for non-positive integers
		*sum += int(math.Pow(float64(num), 4))
	}

	*index++
	processNumbersRecursive(sum, numbers, index)
}

func printResults(results []string) {
	index := 0
	printResultsRecursive(results, &index)
}

func printResultsRecursive(results []string, index *int) {
	// base case
	if *index >= len(results) {
		return
	}

	fmt.Println(results[*index])

	*index++
	printResultsRecursive(results, index)
}

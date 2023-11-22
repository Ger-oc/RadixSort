package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func radixSort(arr []int) []int {
	max := getMax(arr)
	for exp := 1; max/exp > 0; exp *= 10 {
		countingSort(arr, exp)
	}
	return arr
}

func getMax(arr []int) int {
	max := arr[0]
	for _, val := range arr {
		if val > max {
			max = val
		}
	}
	return max
}

func countingSort(arr []int, exp int) {
	n := len(arr)
	output := make([]int, n)
	count := make([]int, 19) // 19 because it includes negative values
	for i := 0; i < n; i++ {
		index := (arr[i]/exp)%10 + 9 // Shift by 9 to handle negative values
		count[index]++
	}
	for i := 1; i < 19; i++ {
		count[i] += count[i-1]
	}
	for i := n - 1; i >= 0; i-- {
		index := (arr[i]/exp)%10 + 9
		output[count[index]-1] = arr[i]
		count[index]--
	}
	for i := 0; i < n; i++ {
		arr[i] = output[i]
	}
}

func main() {
	numbers, err := readNumbersFromCSV("numbers.csv")
	if err != nil {
		log.Fatalf("Error reading numbers: %v", err)
	}
	start := time.Now()
	sortedNumbers := radixSort(numbers)
	elapsed := time.Since(start)
	fmt.Printf("RadixSort took %s\n", elapsed)
	err = writeNumbersToCSV("out.csv", sortedNumbers)
	if err != nil {
		log.Fatalf("Error writing numbers: %v", err)
	}
	if isSorted(sortedNumbers) {
		fmt.Println("The numbers are sorted correctly")
	} else {
		fmt.Println("The numbers are not sorted correctly")
	}
}
func readNumbersFromCSV(filename string) ([]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	var numbers []int
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		for _, value := range record {
			number, err := strconv.Atoi(value)
			if err != nil {
				return nil, err
			}
			numbers = append(numbers, number)
		}
	}
	return numbers, nil
}
func writeNumbersToCSV(filename string, numbers []int) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	for _, number := range numbers {
		err := writer.Write([]string{strconv.Itoa(number)})
		if err != nil {
			return err
		}
	}
	return nil
}
func isSorted(numbers []int) bool {
	for i := 1; i < len(numbers); i++ {
		if numbers[i-1] > numbers[i] {
			return false
		}
	}
	return true
}

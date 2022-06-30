package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := ioutil.ReadFile("data.txt")
	if err != nil {
		fmt.Println(err)
	}
	var numbers []float64
	lines := strings.Split(string(file), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		num, _ := strconv.ParseFloat(line, 64)
		numbers = append(numbers, num)
	}

	// if \

	sort.Float64s(numbers)
	fmt.Println("Average :", math.Round((average(numbers))))
	fmt.Println("Median :", math.Round(median(numbers)))
	fmt.Println("Variance :", math.Round(variance(numbers)))
	fmt.Println("Standard deviation :", math.Round(standardDeviation(numbers)))
}

func average(numbers []float64) float64 {
	var sum float64
	for _, number := range numbers {
		sum += number
	}
	return sum / float64(len(numbers))
}

func median(numbers []float64) float64 {
	var median float64
	if len(numbers)%2 == 0 {
		median = (numbers[len(numbers)/2] + numbers[len(numbers)/2-1]) / 2
	} else {
		median = numbers[len(numbers)/2]
	}
	return median
}

func variance(numbers []float64) float64 {
	mean := average(numbers)
	var variance float64
	for _, number := range numbers {
		variance += (number - mean) * (number - mean)
	}
	return variance / float64(len(numbers))
}

func standardDeviation(numbers []float64) float64 {
	return math.Sqrt(variance(numbers))
}

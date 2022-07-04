package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
)

// calcul Moyen
func average(numbers []float64) float64 {
	var somme float64
	for _, number := range numbers {
		somme += number
	}
	return somme / float64(len(numbers))
}

//calcul Median
func median(numbers []float64) float64 {
	var median float64
	if len(numbers)%2 == 0 {
		median = (numbers[len(numbers)/2] + numbers[len(numbers)/2-1]) / 2
	} else {
		median = numbers[len(numbers)/2]
	}
	return median
}


//calcul Variance 
func variance(numbers []float64) float64 {
	mean := average(numbers)
	var variance float64
	for _, number := range numbers {
		variance += (number - mean) * (number - mean)
	}
	return variance / float64(len(numbers))
}


//calcul Ecart-types
func standardDeviation(numbers []float64) float64 {
	return math.Sqrt(variance(numbers))
}



func main() {
	file, err := ioutil.ReadFile("data.txt")
	if err != nil {
		fmt.Println(err)
	}
	var numbers []float64
	lines := strings.Split(string(file), "\n") // Pour décomposer les informations dans "file" en une liste de sous-chaîne en séparant les informations par des retour à la ligne 
	for _, line := range lines {
		if line == "" {
			continue
		}
		num, _ := strconv.ParseFloat(line, 64)
		numbers = append(numbers, num)
	}

	
	sort.Float64s(numbers)
	fmt.Println("Average :", math.Round((average(numbers))))
	fmt.Println("Median :", math.Round(median(numbers)))
	fmt.Println("Variance :", int(variance(numbers)))
	fmt.Println("Standard deviation :", math.Round(standardDeviation(numbers)))
}

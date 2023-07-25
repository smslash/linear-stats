package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Error: incorrect number of arguments")
	}

	s := os.Args[1]

	file, err := os.Open(s)
	if err != nil {
		log.Fatalln("Error:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	data := []float64{}
	for scanner.Scan() {
		value, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			log.Fatalln("Error:", err)
		}
		data = append(data, value)
	}

	if len(data) < 2 {
		log.Fatalln("Error: not enough data for calculations")
	}

	a, b := Regression(data)
	r := Pearson(data)

	fmt.Printf("Linear Regression Line: y = %.6fx + %.6f\n", a, b)
	fmt.Printf("Pearson Correlation Coefficient: %.10f\n", r)
}

func Regression(data []float64) (float64, float64) {
	n := float64(len(data))

	var sumX float64
	var sumY float64
	var sumXY float64
	var sumXX float64

	for i, y := range data {
		x := float64(i)
		sumX += x
		sumY += y
		sumXY += x * y
		sumXX += x * x
	}

	a := (n*sumXY - sumX*sumY) / (n*sumXX - sumX*sumX)
	b := (sumY - a*sumX) / n

	return a, b
}

func Pearson(data []float64) float64 {
	n := float64(len(data))

	var sumX float64
	var sumY float64
	var sumXY float64
	var sumXX float64
	var sumYY float64

	for i, y := range data {
		x := float64(i)
		sumX += x
		sumY += y
		sumXY += x * y
		sumXX += x * x
		sumYY += y * y
	}

	num := n*sumXY - sumX*sumY
	den := (n*sumXX - sumX*sumX) * (n*sumYY - sumY*sumY)

	if den == 0 {
		return 0
	}

	return num / math.Sqrt(den)
}

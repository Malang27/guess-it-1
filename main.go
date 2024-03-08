package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var numbers []float64
	debut := 0
	fin := 0
	reader := bufio.NewScanner(os.Stdin)
	for reader.Scan() {
		nbr, _ := strconv.Atoi(reader.Text())
		numbers = append(numbers, float64(nbr))
		if len(numbers) <= 10 {
			debut = int(nbr) - 100
			fin = int(nbr) + 100
		} else {
			average := Average(numbers)
			stand_dev := Standard_Dev(numbers)
			debut = Round(average - 1.28*stand_dev)
			fin = Round(average + 1.28*stand_dev)
		}
		fmt.Printf("%d %d\n", debut, fin)

	}

}
func Average(tab []float64) float64 {
	somme := Sum(tab)
	Average := somme / float64(len(tab))
	return Average
}
func Sum(tab []float64) float64 {
	sum := 0.0

	for i := 0; i < len(tab); i++ {
		dif := 0.0
		if i > 0 {
			dif = tab[i] - tab[i-1]
		}
		if dif < 10000.0 && dif > -10000.0 {
			sum += tab[i]

		} /*  else {
			sum += tab[0]
		} */
	}
	return sum
}
func Variance(tab []float64) float64 {
	sum_des_dif := 0.0
	moyenne := Average(tab)

	for _, val := range tab {
		n := val - moyenne
		if n < 10000.0 && n > -10000.0 {
			sum_des_dif += math.Pow(n, 2)

		}
	}
	variance := sum_des_dif / float64(len(tab))
	return variance
}
func Standard_Dev(tab []float64) float64 {
	vari := Variance(tab)
	return math.Sqrt(vari)
}
func Round(val float64) int {
	decimal := math.Mod(val, 1)
	if decimal <= 0.5 {
		return int(val - decimal)
	} else {
		return int(val - decimal + 1)
	}

}

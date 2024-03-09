package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

// Constantes utilisées dans le programme
const (
	initialWindowSize = 10
	stdDevMultiplier  = 1.28
)

func main() {
	var numbers []float64
	var isFirstIteration = true

	// Scanner pour lire depuis l'entrée standard
	reader := bufio.NewScanner(os.Stdin)
	for reader.Scan() {
		// Convertir l'entrée en un nombre entier
		nbr, err := strconv.Atoi(reader.Text())
		if err != nil {
			fmt.Println("Entrée non valide. Veuillez entrer un entier valide.")
			continue
		}

		// Ajouter le nombre à la liste des nombres
		numbers = append(numbers, float64(nbr))
		var debut, fin int

		// Vérifier s'il s'agit de la première itération
		if isFirstIteration {
			// Définir la fenêtre initiale
			debut = nbr - 100
			fin = nbr + 100
			isFirstIteration = false
		} else {
			// Calculer la moyenne et l'écart type
			average := average(numbers)
			standDev := standardDeviation(numbers)
			// Calculer la nouvelle fenêtre en fonction de la moyenne et de l'écart type
			debut = round(average - stdDevMultiplier*standDev)
			fin = round(average + stdDevMultiplier*standDev)
		}

		// Afficher la plage calculée
		fmt.Printf("%d %d\n", debut, fin)
	}
}

// Calculer la moyenne d'un tableau de nombres
func average(tab []float64) float64 {
	somme := sum(tab)
	return somme / float64(len(tab))
}

// Calculer la somme d'un tableau de nombres
func sum(tab []float64) float64 {
	sum := 0.0
	for i := 0; i < len(tab); i++ {
		dif := 0.0
		// Calculer la différence entre les éléments successifs
		if i > 0 {
			dif = tab[i] - tab[i-1]
		}
		// Ajouter l'élément à la somme si la différence est dans la plage spécifiée
		if dif < 10000.0 && dif > -10000.0 {
			sum += tab[i]
		}
	}
	return sum
}

// Calculer la variance d'un tableau de nombres
func variance(tab []float64) float64 {
	sumDesDif := 0.0
	moyenne := average(tab)

	for _, val := range tab {
		n := val - moyenne
		// Ajouter la différence au carré à la somme si elle est dans la plage spécifiée
		if n < 10000.0 && n > -10000.0 {
			sumDesDif += math.Pow(n, 2)
		}
	}
	return sumDesDif / float64(len(tab))
}

// Calculer l'écart type d'un tableau de nombres
func standardDeviation(tab []float64) float64 {
	vari := variance(tab)
	return math.Sqrt(vari)
}

// Arrondir un nombre flottant à l'entier le plus proche
func round(val float64) int {
	decimal := math.Mod(val, 1)
	if decimal <= 0.5 {
		return int(val - decimal)
	}
	return int(val - decimal + 1)
}

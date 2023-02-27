package main

import "fmt"

func allumette(nb_allumette int) {
	var allumette_prise int
	fmt.Println("Combien d'allumette voulez-vous enlever ?")
	fmt.Scanln(&allumette_prise)

	nb_allumette = nb_allumette - allumette_prise
	if nb_allumette <= 1 {
		print("Perdu !")
	} else {
		print("Il reste ", nb_allumette, " allumette(s). \n")
		allumette(nb_allumette)
	}
}

func main() {
	var nb_allumette int
	fmt.Println("Combien d'allumette voulez-vous au départ ?") // n le nombre d'allumette
	fmt.Scanln(&nb_allumette)
	allumette(nb_allumette)
}

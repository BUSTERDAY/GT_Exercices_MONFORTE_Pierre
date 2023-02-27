package main

import "fmt"

func allumette(nb_allumette int) {
	var allumette_prise int
	fmt.Scanln(&allumette_prise)
	fmt.Println("Combien d'allumette voulez-vous enlever ?")

	nb_allumette = nb_allumette - allumette_prise
	print("Il reste", nb_allumette, "allumette(s)")
	if nb_allumette == 1 {
		print("Perdu !")
	} else {
		allumette(nb_allumette)
	}
}

func main() {
	var nb_allumette int
	fmt.Scanln(&nb_allumette)
	fmt.Println("Combien d'allumette voulez-vous au départ ?") // n le nombre d'allumette
	allumette(nb_allumette)
}

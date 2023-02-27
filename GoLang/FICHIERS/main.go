package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	var choice int

	for {
		fmt.Println("Que souhaitez-vous faire ?")
		fmt.Println("1. Récupérer tout le texte contenu dans un fichier .txt")
		fmt.Println("2. Ajouter du texte dans un fichier .txt")
		fmt.Println("3. Supprimer tout le contenu du fichier .txt")
		fmt.Println("4. Remplacer le contenu du fichier .txt par du texte donné par l'utilisateur")
		fmt.Println("5. Quitter")

		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fileContent, err := ioutil.ReadFile("texte.txt")
			if err != nil {
				fmt.Println("Erreur lors de la lecture du fichier :", err)
			} else {
				fmt.Println(string(fileContent))
			}
		case 2:
			fmt.Println("Entrez le texte à ajouter :")
			var textToAdd string
			fmt.Scanln(&textToAdd)

			file, err := os.OpenFile("texte.txt", os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println("Erreur lors de l'ouverture du fichier :", err)
			} else {
				_, err = fmt.Fprintln(file, textToAdd)
				if err != nil {
					fmt.Println("Erreur lors de l'écriture dans le fichier :", err)
				}
				file.Close()
			}
		case 3:
			err := ioutil.WriteFile("texte.txt", nil, 0644)
			if err != nil {
				fmt.Println("Erreur lors de la suppression du contenu du fichier :", err)
			} else {
				fmt.Println("Contenu du fichier supprimé avec succès.")
			}
		case 4:
			fmt.Println("Entrez le nouveau texte :")
			var newText string
			fmt.Scanln(&newText)

			err := ioutil.WriteFile("texte.txt", []byte(newText), 0644)
			if err != nil {
				fmt.Println("Erreur lors de l'écriture dans le fichier :", err)
			} else {
				fmt.Println("Nouveau texte écrit dans le fichier avec succès.")
			}
		case 5:
			fmt.Println("Au revoir !")
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

package hangman

import (
	"bufio"
	"fmt"
	"os"
)

func Affichage() {
	var ligne []string
	file, _ := os.Open("draw/affiche.txt")
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		ligne = append(ligne, fileScanner.Text())
	}
	for _, i := range ligne {
		fmt.Println(i)
	}
}

func AffichageFin() {
	var ligne []string
	file, _ := os.Open("draw/gg.txt")
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		ligne = append(ligne, fileScanner.Text())
	}
	for _, i := range ligne {
		fmt.Println(i)
	}
}

func AffBlank(blank string) {
	var lignes []string
	var index int
	var mot [15][30]string
	file, _ := os.Open("draw/standard.txt")
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		lignes = append(lignes, fileScanner.Text())
	}
	alpha := " !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~"
	for i := 0; i < len(blank); i++ {
		for j := 0; i < len(alpha)-1; j++ {
			if blank[i] == alpha[j] {
				index = j
				break
			}
		}
		lettre := lignes[9*(index)+1 : 9*(index+1)]
		for index, j := range lettre {
			mot[i][index] = j
		}
	}
	for i := 0; i < 9; i++ {
		for j := range blank {
			fmt.Printf(mot[j][i])
			fmt.Printf("  ")
		}
		fmt.Printf("\n")
	}
}

func Mentors() {
	var ligne []string
	file, _ := os.Open("draw/mentor.txt")
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		ligne = append(ligne, fileScanner.Text())
	}
	for _, i := range ligne {
		fmt.Println(i)
	}
}

func AffHangman(nb_err int) {
	var ligne []string
	if nb_err == 0 {
		fmt.Println("")
	} else {
		file, _ := os.Open("draw/hangman.txt")
		defer file.Close()
		fileScanner := bufio.NewScanner(file)
		for fileScanner.Scan() {
			ligne = append(ligne, fileScanner.Text())
		}
		hang := ligne[8*(nb_err-1) : 8*(nb_err-1)+7]
		if nb_err == 11 {
			hang = ligne[71:80]
		}
		for _, i := range hang {
			fmt.Println(i)
		}
	}
}

func Affscoreboard() {
	var ligne []string
	fmt.Println("-------------------------------")
	fmt.Println("NAME           ERRORS     WORD")
	file, _ := os.Open("scoreboard.txt")
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		ligne = append(ligne, fileScanner.Text())
	}
	for i := range ligne {
		infos := SplitWhiteSpaces(string(ligne[i]))
		var spaces string = "     "
		for j := 0; j < 10-len(infos[0]); j++ {
			spaces += " "
		}
		fmt.Println(infos[0] + spaces + infos[1] + "          " + infos[2])
	}
	fmt.Println("")
}

func AffRules() {
	fmt.Println("RULES OF THE GAME :")
	fmt.Println("	- The goal is to guess the sercret word before José finish hung !")
	fmt.Println("	- For that you can try letters (without accents) or words to save him")
	fmt.Println("	- If you make a mistake, José will be threaten")
	fmt.Println("	- You can type HELP for more")
	fmt.Println("")
}

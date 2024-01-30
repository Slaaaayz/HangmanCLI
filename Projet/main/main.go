package main

import (
	"fmt"
	"hangman"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	// Préparation du jeu
	var hard bool
	voy := "aeiouy"
	if len(os.Args) >= 3 {
		if os.Args[1] == "--startWith" {
			hangman.Affichage()
			StartnStopRecup()
		} else if os.Args[1] == "--add" {
			Addwords()
		}
	} else {
		if len(os.Args) == 2 {
			if os.Args[1] == "--hard" {
				hard = true
			}
			if os.Args[1] == "--add" {
				fmt.Println("--add need at least one argument !")
				os.Exit(0)
			}
		}
		hangman.Affichage()
		var pseudo string
		fmt.Println("Choose a username : ")
		fmt.Scanln(&pseudo)
		Rules()
		fmt.Println("Welcome to Hangman " + pseudo + "!")
		var fichier string
		if hard {
			fichier = "words/wordsA3.txt"
		} else {
			if Custom() {
				fichier = "words/customwords.txt"
			} else {
				fichier = ChooseLevel() // Demande au joueur quel niveau de difficulté il veut essayer
				if Anglais() {
					fichier = fichier[:11] + "A" + fichier[11:]
				}
			}
		}
		file, _ := os.ReadFile(fichier) // On lit le fichier sélectionné en fnct du niveau
		words := hangman.Splits(string(file), "\n")
		var ranword int
		if len(words) > 1 {
			rand.Seed(time.Now().UnixNano())
			ranword = rand.Intn(len(words) - 1)
		}
		// Choisis un index au pif
		word := words[ranword] // Et prend le mot à l'index ranword
		var blank string
		for i := 0; i < len(word)-1; i++ { // Stocke le bon nombre de tirets dans une string
			blank += "_"
		}
		if hard {
			blank = FillBlank(blank, len(word)/3-1, word)
		} else {
			blank = FillBlank(blank, len(word)/2-1, word) // Révèle len(word)/2 lettres aléatoires de notre mot
		}
		hangman.AffBlank(blank)
		var tab string
		fmt.Println("Good luck, you have 10 attempts")
		Play(word[:len(word)-1], blank, 0, tab, tab, hard, voy, pseudo) // Lancement du jeu
	}

}

func Play(word string, blank string, err int, wrongLetters string, wrongWords string, hard bool, voy string, name string) {
	var letter string
	var win bool
	var check bool
	var loose bool
	var try bool
	fmt.Println("Choose a letter or guess a word: ")
	fmt.Scanln(&letter) // Récupère la réponse du joueur
	if letter == "STOP" {
		fmt.Println("Save...")
		fmt.Println("Saved ! See you !")
		StartnStopSave(word, blank, err, wrongLetters, wrongWords, hard, voy, name)
	}
	if letter == word { //Si il a deviné le mot
		win = true //Gagné
	} else if letter == "mentor" || letter == "mentors" {
		hangman.Mentors()
		fmt.Println("\non vous aime <3")
		win = true
	} else if letter == "HELP" {
		fmt.Println("  List of possible commands :")
		fmt.Println("	• STOP : Stop the game and save your progression ")
		fmt.Println("	• SCOREBOARD : Print the progression of the all players")
		fmt.Println("	• RESET : Clear scoreboard")
		fmt.Println("	• RULES : Print rules of game")
		fmt.Println("	• RESETCUSTOM: Delete custom words")
	} else if letter == "SCOREBOARD" {
		hangman.Affscoreboard()
	} else if letter == "RESET" {
		ResetScoreBoard()
	} else if letter == "RESETCUSTOM" {
		ResetCustom()
	} else if letter == "RULES" {
		hangman.AffRules()
	} else { //Sinon
		if len(letter) == 1 && hard {
			var stockerr int
			stockerr, voy = voyelle(voy, letter)
			err += stockerr
		}
		for i := range word { // On parcours le mot
			if letter == string(word[i]) { //Si il a deviné un lettre
				blank = string(blank[:i]) + letter + string(blank[i+1:]) //On Modifie blank
				check = true
			}
		}
		if !check && (len(letter) == 1 || Accent(letter)) { // Si il tenté une lettre mais que c'est faux
			try = false
			if IsAlpha(letter) {
				fmt.Println("Choose a letter that is in the alphabet, without accent !")
			}
			for i := range wrongLetters {
				if letter == string(wrongLetters[i]) {
					fmt.Println("Your already tried this letter")
					fmt.Println("")
					try = true
				}
			}
			if !try && !(Accent(letter)) {
				wrongLetters += letter + " "
				err++ // On lui rajoute une erreur
				fmt.Println("Not present in the word, " + strconv.Itoa(10-err) + " attempts remaining")
				fmt.Printf("\n")
			} else if hard {
				err++
			}
		} else if !check { // Si il a tenté un mot mais qu'il a faux
			wordalreadyTried := hangman.Splits(string(wrongWords), " ") // Stocke les mots du fichier sélectionné dans une liste
			try := false
			for i := range wordalreadyTried {
				if letter == string(wordalreadyTried[i]) { // Si il a déjà tenté ce mot
					fmt.Println("You already tried this word")
					fmt.Printf("\n")
					try = true
					break
				}
			}
			if !try {
				err += 2 // On lui rajoute 2 erreurs
				if err != 11 {
					fmt.Println("Wrong word, nice try " + strconv.Itoa(10-err) + " attempts remaining")
					fmt.Printf("\n")
				}
				wrongWords += letter + " "
			} else if hard {
				err++
			}
		}
	}
	if err >= 10 { // Si il a fait plus de 9 erreurs
		loose = true //Perdu
	}
	if blank == word { // Si le mot est complété
		win = true // Gagné
	}
	if win || loose { // Si c'est la fin de la partie (gagné ou perdu)
		hangman.AffHangman(err) //On affiche le pendu
		if win {                // Bravo si c'est gagné
			fmt.Printf("\n")
			fmt.Println("CONGRATULATIONS")
			fmt.Println("The word : ")
			hangman.AffBlank(word)
			Scoreboard(name, err, word)
			fmt.Printf("\n")
			hangman.AffichageFin()
		} else { // Sinon t vrmt fonfon
			fmt.Printf("\n")
			fmt.Println("Oh Snap !")
			fmt.Println("The word was : " + word)
			fmt.Printf("\n")
			hangman.AffichageFin()
		}
		var answer string
		fmt.Println("Do you want replay ? (y/n)")
		fmt.Scanln(&answer)
		if answer == "y" {
			main()
		} else if answer == "n" {
			fmt.Println("See you !")
		}
	} else { // C'est pas la fin de la partie
		hangman.AffBlank(blank)
		// Affiche les lettres et les mots tentés mais faux
		if len(wrongLetters) > 0 && !hard {
			fmt.Printf("\n")
			fmt.Println("Tried letters : " + wrongLetters)
		}
		if len(wrongWords) > 0 && !hard {
			fmt.Printf("\n")
			fmt.Println("Tried words : " + wrongWords)
		}
		hangman.AffHangman(err)                                           //On affiche le pendu
		Play(word, blank, err, wrongLetters, wrongWords, hard, voy, name) // Et on relance un tour
	}
}

func FillBlank(blank string, n int, word string) string { // Révèle des lettres aléatoires de notre mot
	for i := 0; i < n; i++ {
		ind := rand.Intn(len(word) - 1)
		c := word[ind]
		blank = string(blank[:ind]) + string(c) + string(blank[ind+1:])
	}
	return string(blank)
}

func ChooseLevel() string { // Laisse le joueur choisir le level
	var nbr string
	fmt.Printf("\n")
	fmt.Println("Choose a level between 1 (easy word) and 3 (difficult word): ")
	fmt.Scanln(&nbr)
	if nbr != "1" && nbr != "2" && nbr != "3" { // Tant que le level choisi est invalide
		return ChooseLevel() // On repose la question
	}
	// Et on return le fichier correspondant au level choisi
	if nbr == "1" {
		return "words/words1.txt"
	} else if nbr == "2" {
		return "words/words2.txt"
	} else {
		return "words/words3.txt"
	}
}

func StartnStopRecup() {
	var stock string
	var tb []string
	file, _ := os.ReadFile("save.txt")
	if os.Args[1] == "--startWith" {
		stock = string(file)
	}
	if stock == "" {
		fmt.Println("No game saved")
	} else {
		tb = hangman.Splits(stock, "\n")
		var wl string
		var ww string
		for i := 3; i <= len(tb)-1; i++ {
			if len(tb[i]) == 1 {
				wl += tb[i] + " "
			} else {
				ww += tb[i] + " "
			}
		}
		errors, _ := strconv.Atoi(tb[2])
		var hard bool
		var voy string
		name := tb[5]
		if tb[len(tb)-2] == "T" {
			hard = true
			voy = tb[len(tb)-1]
		}
		fmt.Println("Welcome Back ! ")
		fmt.Println(tb[1])
		os.Remove("save.txt")
		os.Create("save.txt")
		fmt.Println("Good luck, you have " + strconv.Itoa(10-errors) + " attempts left !")
		Play(tb[0], tb[1], errors, wl, ww, hard, voy, name)
	}
}

func StartnStopSave(word string, blank string, err int, wrongLetters string, wrongWords string, hard bool, voy string, name string) {
	file, _ := os.OpenFile("save.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	defer file.Close()
	_, err2 := file.WriteString(word + "\n" + blank + "\n" + strconv.Itoa(err) + "\n")
	_, err3 := file.WriteString(wrongLetters + "\n" + wrongWords + "\n" + name)
	var err4 error
	if hard {
		_, err4 = file.WriteString("\nT\n" + voy)
	} else {
		_, err4 = file.WriteString("\nF")
	}
	if err2 != nil || err3 != nil || err4 != nil {
		panic(err2)
	}
	os.Exit(1)
}

func voyelle(voy string, letter string) (int, string) {
	var verif bool
	for i, elt := range voy {
		if string(elt) == letter {
			verif = true
			if i == len(voy)-1 {
				voy = voy[:i]
			} else {
				voy = voy[:i] + voy[i+1:]
			}
		}
	}
	if verif {
		if len(voy) < 3 {
			fmt.Println("Tu essaie trop de voyelles, tu es pénalisé")
			return 1, voy
		}
	}
	return 0, voy
}

func Scoreboard(name string, err int, word string) {
	file, _ := os.OpenFile("scoreboard.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	defer file.Close()
	_, err2 := file.WriteString(name + " " + strconv.Itoa(err) + " " + word + "\n")
	if err2 != nil {
		panic(err2)
	}
}

func ResetScoreBoard() {
	var verif string
	fmt.Println("Are you sure you want to reset the scoreboard ? (y/n)")
	fmt.Scanln(&verif)
	if verif == "y" {
		os.Remove("scoreboard.txt")
		os.Create("scoreboard.txt")
		fmt.Println("Scoreboard reset succeded !")
	} else if verif == "n" {
		fmt.Println("Scoreboard reset aborted !")
	} else {
		ResetScoreBoard()
	}
}

func ResetCustom() {
	var verif string
	fmt.Println("Are you sure you want to delete the custom words ? (y/n)")
	fmt.Scanln(&verif)
	if verif == "y" {
		os.Remove("words/customwords.txt")
		os.Create("words/customwords.txt")
		fmt.Println("Custom words delete succeded !")
	} else if verif == "n" {
		fmt.Println("Custom words delete aborted !")
	} else {
		ResetCustom()
	}
}

func IsAlpha(s string) bool {
	for _, i := range s {
		if !((i >= 97 && i <= 122) || (i >= 65 && i <= 90)) {
			return false
		}
	}
	return true
}

func Accent(letter string) bool {
	accents := "àâäéèêëîïôöùûüÿç"
	for _, elt := range accents {
		if letter == string(elt) {
			return true
		}
	}
	return false
}

func Rules() {
	var answer string
	fmt.Println("Do you know the rules of Hangman ? (y/n)")
	fmt.Scanln(&answer)
	if answer == "n" {
		hangman.AffRules()
	}
	if answer != "y" && answer != "n" {
		Rules()
	}
}

func Anglais() bool {
	var answer string
	fmt.Println("Do you want to play with french words ? (y/n)")
	fmt.Scanln(&answer)
	if answer == "n" {
		return true
	} else if answer == "y" {
		return false

	} else {
		return Anglais()
	}
}

func Addwords() {
	for i := 2; i < len(os.Args); i++ {
		if IsAlpha(os.Args[i]) {
			file, _ := os.OpenFile("words/customwords.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
			defer file.Close()
			_, err := file.WriteString(os.Args[i] + ".")
			if i != len(os.Args)-1 {
				_, err1 := file.WriteString("\n")
				if err1 != nil {
					panic(err1)
				}
			}
			if err != nil {
				panic(err)
			}
		} else {
			fmt.Println("The word " + os.Args[i] + " is not valid. No accent, no special characters please !")
		}
	}
}

func Custom() bool {
	file, _ := os.ReadFile("words/customwords.txt")
	stock := string(file)
	if stock != "" {
		var ans string
		fmt.Println("Do you want to play with the customs words ? (y/n)")
		fmt.Scanln(&ans)
		if ans == "y" {
			return true
		} else if ans == "n" {
			return false
		} else {
			return Custom()
		}
	}
	return false
}

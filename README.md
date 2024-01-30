# <div align="center">HANGMAN
## SOMMAIRE
- [I. Comment installer le Hangman](#i-comment-installer-le-hangman)
    - [Executer le Hangman dans un shell](#executer-le-hangman-dans-un-shell)
- [II. Fonctionnement du Hangman](#ii-fonctionnement-du-hangman)
- [III. Comment jouer au Hangman](#iii-comment-jouer-au-hangman)
- [IV. Les commandes qui peuvent être utilisé dans le Hangman](#iv-les-commandes-qui-peuvent-être-utilisé-dans-le-hangman)
- [V. Les "Easter Eggs" du Hangman](#v-les-easter-eggs-du-hangman)
- [VI. Si le Hangman ne fonctionne pas](#vi-si-le-hangman-ne-fonctionne-pas)
- [VII. Comment fonctionne le code du Hangman](#vii-comment-fonctionne-le-code-du-hangman)

## I. Comment installer le Hangman
Pour installer le Hangman, il faut d'abord cloner le repository sur votre ordinateur. Pour cela, il faut ouvrir un terminal et taper la commande suivante :
```bash
git clone https://ytrack.learn.ynov.com/git/cmaxime/hangman-classic
```
Ensuite, il faut se rendre sur VsCode et ouvrir le dossier hangman. Une fois le dossier ouvert, il faut ouvrir un terminal et taper la commande suivante :
```go
cd .\Projet\
```
Puis pour lancer le jeu taper la commande ci-dessous :
```go
go run .\main\main.go
```
### Executer le Hangman dans un shell  
Après avoir clone le repository, se rendre dans le dossier hangman avec la commande suivante :
```powershell
cd hangman-classic/Projet
```
 Une fois le dossier ouvert, il faut taper la commande suivante :
```powershell
go build .\main\main.go
```
Un fichier éxecutable sera alors crée il suffit maintenant de taper :
```powershell
.\main.exe
```
## II. Fonctionnement du Hangman
Le Hangman est un jeu de pendu. Le but est de trouver le mot caché en devinant les lettres qui le composent. Pour cela, le joueur a droit à 10 essais. Si le joueur trouve le mot, il gagne la partie. Sinon, il perd la partie.
## III. Comment jouer au Hangman 
Pour jouer au Hangman, il faut d'abord lancer le jeu. Ensuite, il faut choisir un pseudo puis niveau parmi les 3 proposés. Une fois le niveau choisi, le joueur doit deviner le mot caché en proposant des lettres ou un mot. 

**Exemple de partie gagnée &#x1F609; :**   

![Alt text](<Projet/image/hang win.PNG>)

 
**Exemple de partie perdue &#x1F60F;:** 

![Alt text](<Projet/image/hang loose.PNG>)


## IV. Les commandes qui peuvent être utilisé dans le Hangman  
Voici les commandes à utiliser en cours de partie :  
- STOP : Permet de d'arreter le jeu et de créer un point de progression de la ou la partie s'est arretée.
- SCOREBOARD : Permet d'afficher le tableau des scores.
- RESET : Permet de vider à 0 le tableau des scores.
- RULES : Permet d'afficher les règles du jeu
- RESETCUSTOM : Permet de vider la liste de mot que le joueur à ajouté 
- HELP : Permet d'afficher les commandes disponibles.

⚠️ ATTENTION ces commandes sont à marquer en majuscule dans le terminal
## V. Les "Easter Eggs" du Hangman
- Si le joueur entre "STOP" à n'importe quel moment du jeu dans le terminal, il quitte le jeu et il obtient une save qu'il pourra relancer plus tard avec la commande suivante : 
```go
go run .\main\main.go --startWith save.txt
```
- Si le joueur souhaite jouer au hangman en version Hard il suffira alors de taper la commande : 
```go
go run .\main\main.go --hard
```
- Si le joueur souhaite ajouter son propre mot au Hangman alors il faut taper dans le terminal :
```go
go run .\main\main.go --add (mot à ajouter)
```
- Si le joueur entre le mot mentor ou mentors dans le terminal alors il a directement gagné la partie.
## VI. Si le Hangman ne fonctionne pas
Si le Hangman ne fonctionne pas, il faut d'abord vérifier que vous avez bien installé Go sur votre ordinateur. 

Si Go est bien installé, il faut vérifier que vous avez bien installé le repository du Hangman.

Si le repository est bien installé, il faut vérifier que vous avez bien ouvert le dossier hangman sur VsCode.

Si le Hangman de s'execute pas après ces étapes alors supprimer le dossier hangman et recommencer l'installation du Hangman.

Si le Hangman présente un bug dans le terminal utiliser "CTRL + c" ou la commande suivante pour le relancer et permettre d'avoir une sauvgarde:
```go
STOP
```
```go
go run .\main\main.go --startWith save.txt
```

Si le Hangman ne fonctionne toujours pas, il faut contacter les auteurs du Hangman.
## VII. Comment fonctionne le code du Hangman
**Le code du Hangman est composé de 3 fichier .go:**
- main : contient le main.go qui lance le jeu
- display : contient les fonctions du jeu qui permette l'affichage graphique 
- split : Décompose une string en un tableau de strings 
 
## <div align="right">Les auteurs du Hangman
<div align="right">SAUTEREAU DU PART Diane  
<div align="right">COLOMBAN-FERNANDEZ Luna  
<div align="right">CHORT Maxime
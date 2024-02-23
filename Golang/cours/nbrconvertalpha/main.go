package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	my_arr := os.Args[1:]
	// Récupère les arguments de la ligne de commande à partir du deuxième argument (index 1) et les stocke dans my_arr.
	ln := 0
	for i := range my_arr {
		ln = i
	}
	// Calcule le nombre d'arguments dans my_arr et stocke le résultat dans ln.
	if ln >= 1 {
		// Vérifie s'il y a au moins un argument (après "--upper").
		if my_arr[0] == "--upper" {
			// Vérifie si le premier argument est "--upper".
			z01.PrintRune(0)
			// Si "--upper" est présent, imprime un espace.
			for i := 1; i <= ln; i++ {
				// Parcourt les arguments suivants après "--upper".
				num := 0
				// Initialise une variable num pour stocker la valeur numérique de l'argument.
				for _, w := range my_arr[i] {
					// Parcourt chaque caractère de l'argument.
					num = num*10 + int(w-'0')
					// Convertit la séquence de caractères en entier.
				}
				if num >= 1 && num <= 26 {
					z01.PrintRune('A' + rune(num-1))
					// Si l'entier est dans la plage 1-26, imprime la lettre majuscule correspondante.
				} else {
					z01.PrintRune(' ')
					// Sinon, imprime un espace.
				}
			}
		} else {
			// Si le premier argument n'est pas "--upper", alors il n'y a pas d'option "--upper".
			for i := 0; i <= ln; i++ {
				// Parcourt tous les arguments.
				myNum := 0
				// Initialise une variable myNum pour stocker la valeur numérique de l'argument.
				for _, w := range my_arr[i] {
					// Parcourt chaque caractère de l'argument.
					myNum = myNum*10 + int(w-'0')
					// Convertit la séquence de caractères en entier.
				}
				if myNum >= 1 && myNum <= 26 {
					z01.PrintRune('a' + rune(myNum-1))
					// Si l'entier est dans la plage 1-26, imprime la lettre minuscule correspondante.
				} else {
					z01.PrintRune(' ')
					// Sinon, imprime un espace.
				}
			}
		}
	}
	z01.PrintRune('\n')
	// Imprime un saut de ligne à la fin
}

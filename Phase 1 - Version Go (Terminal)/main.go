package main

import (
	"fmt"
	"time"
)

func main() {
	var menuchoice float64
	sleep := time.Duration(3) * time.Second
	for {
		ClearScreen()
		fmt.Println("Bonjour que voulez vous convertir aujourd'hui ?")
		fmt.Println("1 - Convertir des Kilomètres en Miles")
		fmt.Println("2 - Convertir des Kilogrammes en Livres")
		fmt.Println("3 - Convertir des degrés Celsius en Fahrenheit")
		fmt.Println("4 - Quitter")
		_, err := fmt.Scan(&menuchoice)
		if err != nil {
			fmt.Println("❌ Choix impossible, réessayez.")
			time.Sleep(sleep)
			continue
		}
		switch menuchoice {
		case 1:
			for {
				ClearScreen()
				fmt.Println("Combien de Kilomètres voulez vous convertir ?")
				fmt.Println("Rentrez le nombre de Kilomètres que vous voulez convertir en Miles:")
				fmt.Println("0 : Pour revenir en arrière")

				_, err := fmt.Scan(&menuchoice)
				if err != nil {
					fmt.Println("❌ Choix impossible, réessayez.")
					time.Sleep(sleep)
					continue
				}

				if menuchoice == 0 {
					break
				} else {
					fmt.Println(menuchoice, "km équivalent à :", Kilometers2Miles(menuchoice), "miles.") // j'au utilisé Println pour garder la précision des résultats mais si on veut des résultats arrondis il faut utiliser Printf
					time.Sleep(sleep)
				}
			}
		case 2:
			for {
				ClearScreen()
				fmt.Println("Combien de Kilogrammes voulez vous convertir ?")
				fmt.Println("Rentrez le nombre de Kilogrammes que vous voulez convertir en lbs/pounds:")
				fmt.Println("0 : Pour revenir en arrière")

				_, err := fmt.Scan(&menuchoice)
				if err != nil {
					fmt.Println("❌ Choix impossible, réessayez.")
					time.Sleep(sleep)
					continue
				}

				if menuchoice == 0 {
					break
				} else {
					fmt.Println(menuchoice, "kg équivalent à :", Kilograms2pounds(menuchoice), "lbs/pounds.")
					time.Sleep(sleep)
				}
			}
		case 3:
			for {
				ClearScreen()
				fmt.Println("Combien de degrés Celsius voulez vous convertir ?")
				fmt.Println("Rentrez le nombre de degrés Celsius que vous voulez convertir en Fahreneit:")
				fmt.Println("0 : Pour revenir en arrière")

				_, err := fmt.Scan(&menuchoice)
				if err != nil {
					fmt.Println("❌ Choix impossible, réessayez.")
					time.Sleep(sleep)
					continue
				}

				if menuchoice == 0 {
					break
				} else {
					fmt.Println(menuchoice, "°C équivalent à :", Celsius2Fahreneit(menuchoice), "degrés Fahrenheit.")
					time.Sleep(sleep)
				}
			}
		case 4:
			return
		default:
			fmt.Println("❌ Choix impossible, réessayez.")
			time.Sleep(sleep)
		}
	}
}

func Kilometers2Miles(nb float64) float64 {
	nb = nb / 1.609344
	return nb
}

func Kilograms2pounds(nb float64) float64 {
	nb = nb / 0.45359237
	return nb
}

func Celsius2Fahreneit(nb float64) float64 {
	nb = (nb*9)/5 + 32
	return nb
}

func ClearScreen() {
	fmt.Print("\033[H\033[J")
}

package main

import (
	"fmt"
	"errors"
)

func dzielenie(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Dzielenie przez zero jest niedozwolone.")
	}
	return a / b, nil
}

func main() {
	var dzielnik int
	fmt.Print("Podaj dzielnik: ")
	_, err := fmt.Scan(&dzielnik)
	if err != nil {
		fmt.Println("Błąd: Podano nieprawidłową wartość.")
		return
	}

	wynik, err := dzielenie(10, dzielnik)
	if err != nil {
		fmt.Println("Błąd: ", err)
		return
	}else{
		fmt.Println("Wynik: ", wynik)
	}
}
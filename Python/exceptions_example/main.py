try:
    dzielnik = int(input("Podaj dzielnik: "))
    wynik = 10 / dzielnik
    print(f"Wynik: {wynik}")
except ZeroDivisionError:
    print("Błąd: Dzielenie przez zero jest niedozwolone.")
except ValueError:
    print("Błąd: Wprowadź liczbę całkowitą.")
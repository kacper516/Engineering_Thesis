#include <iostream>
#include <stdexcept>

int main()
{
    try
    {
        int dzielnik{};
        std::cout << "Podaj dzielnik: ";
        std::cin >> dzielnik;
        if (dzielnik == 0)
            throw std::runtime_error("Dzielenie przez zero jest niedozwolone.");
        int wynik = 10 / dzielnik;
        std::cout << "Wynik: " << wynik << std::endl;
    }
    catch (const std::exception &e)
    {
        std::cerr << "Blad: " << e.what() << std::endl;
    }
}
use std::io

fn main() {
    println!("Podaj dzielnik: ")
    let mut input = String::new();
    io::stdin().read_line(&mut input).unwrap();
    let dzielnik: i32 = match input.trim().parse(){
        Ok(num) => num,
        Err(_) => {
            println!("Błąd: Podano nieprawidłową wartość.");
            return;
        }
    };

    if dzielnik == 0 {
        println!("Błąd: Dzielenie przez zero jest niedozwolone.");
    } else {
        let wynik = 10 / dzielnik;
        println!("Wynik dzielenia: {}", wynik);
    }
}
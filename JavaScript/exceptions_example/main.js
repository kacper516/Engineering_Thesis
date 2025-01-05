process.stdout.write("Podaj dzielnik: ");
process.stdin.on("data", (data) => {
  try {
    const dzielnik = parseInt(data.toString().trim());
    if (isNaN(dzielnik)) {
      throw new Error("Podano nieprawidłową wartość.");
    }
    if (dzielnik === 0) {
      throw new Error("Dzielenie przez zero jest niedozwolone.");
    }
    const wynik = 10 / dzielnik;
    console.log(`Wynik: ${wynik}`);
  } catch (error) {
    console.log(`Błąd: ${error.message}`);
  } finally {
    process.stdin.pause();
  }
});

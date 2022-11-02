package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 	bet        int
// 	lastGame   string
// 	savedGames []string
// 	counter    int = 0
// )

func main() {
	rand.Seed(time.Now().UnixNano())

	minE := 5
	maxE := 10
	minO := 3
	maxO := 5
	minS := 3
	maxS := 4

	game := NewGame(Lotofacil)

	_, err := game.checkBet(15)
	if err != nil {
		fmt.Println(err)
	}

	err = game.Game(minE, maxE, minO, maxO, minS, maxS)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(game)
	// 	for {
	// 		input := greetings()

	// 		switch input {
	// 		case "1":
	// 			lastGame = formattedOutput(lotofacil.Game([]int{}, 6))
	// 			counter++
	// 			fmt.Println()
	// 			fmt.Printf("Jogo %v: %v\n", counter, lastGame)
	// 			fmt.Println()
	// 			savedGames = append(savedGames, lastGame)

	// 		case "2":
	// 			lastGame = formattedOutput(notSoRandGame())
	// 			counter++
	// 			fmt.Println()
	// 			fmt.Printf("Jogo %v: %v\n", counter, lastGame)
	// 			fmt.Println()
	// 			savedGames = append(savedGames, lastGame)

	// 		case "3":
	// 			for i, v := range savedGames {
	// 				fmt.Println()
	// 				fmt.Printf("Jogo %v: %v\n", i+1, v)
	// 			}
	// 		case "4":
	// 			writeFile(savedGames)
	// 			fmt.Println()
	// 			fmt.Println("Jogos Salvos!")
	// 			fmt.Println()
	// 		case "5":
	// 			return
	// 		}
	// 	}
	// }

	// // Main menu
	// func greetings() string {
	// 	possibleValues := [5]string{"1", "2", "3", "4", "5"}
	// 	for {
	// 		fmt.Println()
	// 		fmt.Println("Selecione a opcao desejada:")
	// 		fmt.Println("1. Gerar novo jogo completo aleatorio.")
	// 		fmt.Println("2. Completar jogo com numeros inseridos manualmente.")
	// 		fmt.Println("3. Exibir jogos realizados.")
	// 		fmt.Println("4. Salvar jogos em um documento de texto.")
	// 		fmt.Println("5. Sair.")
	// 		fmt.Println()
	// 		scanner := bufio.NewScanner(os.Stdin)
	// 		scanner.Scan()
	// 		input := scanner.Text()
	// 		for _, v := range possibleValues {
	// 			if input == v {
	// 				return input
	// 			}
	// 		}
	// 		fmt.Println("Opcao nao encontrada, favor tentar novamente.")
	// 	}
	// }

	// // Generates a random 15 numbers game with some numbers already provided by user input.
	// func notSoRandGame() []int {
	// 	// Initializing other variables.
	// 	scanner := bufio.NewScanner(os.Stdin)
	// 	var randomGame []int
	// 	var inputInt []int
	// 	var input []string

	// UserInput:
	// 	for {
	// 		fmt.Println()
	// 		fmt.Println("Digite os numeros que deseja jogar que o programa completara os valores restantes, respeitando as seguintes regras:")
	// 		fmt.Println("- Jogar os numeros separados por espacos, exemplo \"2 3 6 12...\"")
	// 		fmt.Println("- Jogar de 1 a 14 numeros")
	// 		fmt.Println("- Jogar numeros entre 1 e 25.")
	// 		fmt.Println()
	// 		scanner.Scan()

	// 		// Insert input elements into a slice removing white spaces.
	// 		input = strings.Fields(scanner.Text())

	// 		// Check lenght of input.
	// 		if len(input) == 0 || len(input) > 14 {
	// 			fmt.Println("Quantidade de numeros nao permitida, favor tente novamente.")
	// 			continue UserInput
	// 		}

	// 		// Check if valid ints and append to a slice of ints
	// 		inputInt = nil
	// 		for _, v := range input {
	// 			intv, err := strconv.Atoi(v)
	// 			if err != nil {
	// 				fmt.Printf("Valor invalido inserido [%s], favor tente novamente.\n", v)
	// 				continue UserInput
	// 			}
	// 			if !contains(lotofacil.Numbers(), intv) {
	// 				fmt.Printf("Valores devem estar entre %d e %d, favor tente novamente.\n", lotofacil.minNumb, lotofacil.maxNumb)
	// 				continue UserInput
	// 			}
	// 			inputInt = append(inputInt, intv)
	// 		}
	// 		break
	// 	}

	// 	// Merges user input with possible numbers.
	// 	randomGame = lotofacil.Game(inputInt, 6)
	// 	return randomGame
	// }

	// // Checks if a slice `m` of contains `n`.
	// func contains(m []int, n int) bool {
	// 	for _, v := range m {
	// 		if v == n {
	// 			return true
	// 		}
	// 	}
	// 	return false
	// }

	// // Formats slices of ints to final string output to be stored on savedGames.
	// func formattedOutput(slice []int) string {
	// 	var str []string
	// 	var final string
	// 	var sep string = " - "

	// 	// Sorting the slice ascending
	// 	sort.Ints(slice)

	// 	// Converts the slice of ints into a new slice of strings
	// 	for _, v := range slice {
	// 		str = append(str, strconv.Itoa(v))
	// 	}

	// 	// Join these strings with a '-' separator for final game.
	// 	final = strings.Join(str, sep)
	// 	return final
	// }

	// func writeFile(slice []string) {
	// 	loc, _ := time.LoadLocation("America/Recife")
	// 	current_time := time.Now().In(loc)
	// 	time := current_time.Format("Jan 2, 2006 at 3:04 PM")

	// 	f, _ := os.Create("Jogos.txt")

	// 	defer f.Close()

	// 	_, err := f.WriteString(time + "\n")
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	for i := range slice {
	// 		text := "Jogo " + strconv.Itoa(i) + ": " + slice[i] + "\n"
	// 		_, err2 := f.WriteString(text)
	// 		if err2 != nil {
	// 			panic(err2)
	// 		}
	// 	}
}

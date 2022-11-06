package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/pedropaccola/go-lotofacil/lotto"
)

var (
	lastGame   string
	savedGames []string
	counter    int
)

func init() {
	fmt.Println("Escolha o jogo que deseja apostar:")
	for i := lotto.Lotofacil; i <= lotto.Quina; i++ {

	}

}

func main() {
	rand.Seed(time.Now().UnixNano())

	// s := lotto.GameSettings{
	// 	Lt:      lotto.Lotofacil,
	// 	MinEven: 5,
	// 	MaxEven: 10,
	// 	MinOdd:  3,
	// 	MaxOdd:  5,
	// 	MinSeq:  3,
	// 	MaxSeq:  4,
	// }

	// jogo := lotto.NewGame(s)

	// _, err := jogo.CheckBet(15)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// err = jogo.Game()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(jogo)

	for {
		menu()

		input, err := read(0, 5)
		if err != nil {
			fmt.Println(err)
			continue
		}

		switch input {
		case 1:
			lastGame = fmt.Sprintf()
			counter++
			fmt.Println()
			fmt.Printf("Jogo %v: %v\n", counter, lastGame)
			fmt.Println()
			savedGames = append(savedGames, lastGame)

		case 2:
			lastGame = formattedOutput(notSoRandGame())
			counter++
			fmt.Println()
			fmt.Printf("Jogo %v: %v\n", counter, lastGame)
			fmt.Println()
			savedGames = append(savedGames, lastGame)

		case 3:
			for i, v := range savedGames {
				fmt.Println()
				fmt.Printf("Jogo %v: %v\n", i+1, v)
			}
		case 4:
			writeFile(savedGames)
			fmt.Println()
			fmt.Println("Jogos Salvos!")
			fmt.Println()
		case 5:
			return
		}
	}
}

// Read input
func read(min, max int) (int, error) {
	var input int
	fmt.Scan(&input)
	if input < min || input > max {
		return 0, fmt.Errorf("opcao nao encontrada, favor tentar novamente")
	}
	return input, nil

}

// Main menu
func menu() {
	fmt.Println()
	fmt.Println("Selecione a opcao desejada:")
	fmt.Println("1. Gerar novo jogo completo aleatorio.")
	fmt.Println("2. Completar jogo com numeros inseridos manualmente.")
	fmt.Println("3. Exibir jogos realizados.")
	fmt.Println("4. Salvar jogos em um documento de texto.")
	fmt.Println("5. Sair.")
	fmt.Println()
}

package game

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/pedropaccola/go-lotofacil/lotto"
)

type Game struct {
	settings    lotto.BetSettings
	bet         *lotto.Bet
	betQuantity int
	savedBets   []string
}

func NewGame() *Game {
	return &Game{
		savedBets:   []string{},
		betQuantity: 1,
	}
}

func (g *Game) Start() {
	var lottoIota int
	for {
		g.clearScr()

		fmt.Println("=====================================================")
		fmt.Println()
		fmt.Println("         GERADOR DE APOSTAS DE LOTERIA")
		fmt.Println()
		fmt.Println("=====================================================")
		fmt.Println()
		fmt.Println("Escolha o jogo que deseja apostar:")
		fmt.Println()

		i := 1
		for l := lotto.LottoName(1); l < lotto.LastLotto; l++ {
			fmt.Printf("\t%d) %s\n", i, l)
			i++
		}
		fmt.Println()
		fmt.Printf("\t%d) Sair\n", i)

		inp, err := g.readInput(1, i)
		if err != nil {
			g.errorHandler(err)
			continue
		}
		if inp == i {
			os.Exit(0)
		} else {
			lottoIota = inp
		}
		break
	}
	g.settings = lotto.BetSettings{
		Lt:      lotto.LottoName(lottoIota),
		Bet:     0,
		MaxEven: 100,
		MaxOdd:  100,
		MaxSeq:  15,
	}
	g.configGame()
}

func (g *Game) configGame() {
	for {
		g.clearScr()

		fmt.Println("=====================================================")
		fmt.Println("              CUIDADO COM OS MAXIMOS!!!")
		fmt.Println("     Sua aposta pode nao ser gerada corretamente")
		fmt.Println("Valores padroes (100) para evitar quaisquer restricoes")
		fmt.Println("=====================================================")
		fmt.Println()
		fmt.Printf("Voce esta apostando na (%s).\n", g.settings.Lt)
		fmt.Println("Selecione a opcao caso deseje alterar alguma restricao:")
		fmt.Println()
		fmt.Printf("\t%d) Aposta de (%d) dezenas.<--- OBRIGATORIO\n", 1, g.settings.Bet)
		fmt.Printf("\t%d) Maximo de (%d) numeros pares.\n", 2, g.settings.MaxEven)
		fmt.Printf("\t%d) Maximo de (%d) numeros impares.\n", 3, g.settings.MaxOdd)
		fmt.Printf("\t%d) Maximo de (%d) numeros em sequencia.\n", 4, g.settings.MaxSeq)
		fmt.Println()
		fmt.Printf("\t%d) Numero de jogos que deseja gerar (%d)\n", 5, g.betQuantity)
		fmt.Println()
		fmt.Println("Selecione um dos valores acima para configurar o jogo ou")
		fmt.Println("selecione:")
		fmt.Println()
		fmt.Printf("\t%d) Realizar aposta!\n", 6)
		fmt.Printf("\t%d) Voltar\n", 7)
		fmt.Println()
		fmt.Printf("\t%d) Sair\n", 8)

		inp, err := g.readInput(1, 8)
		if err != nil {
			g.errorHandler(err)
			continue
		}
		switch inp {
		case 1:
			fmt.Println("Quantas dezenas deseja apostar?")
			inp2, err := g.readInput(1, 100)
			if err != nil {
				g.errorHandler(err)
				continue
			}
			g.settings.Bet = inp2
			continue
		case 2:
			fmt.Println("Qual o maximo de pares desejados?")
			inp2, err := g.readInput(1, 100)
			if err != nil {
				g.errorHandler(err)
				continue
			}
			g.settings.MaxEven = inp2
			continue
		case 3:
			fmt.Println("Qual o maximo de impares desejados?")
			inp2, err := g.readInput(1, 100)
			if err != nil {
				g.errorHandler(err)
				continue
			}
			g.settings.MaxOdd = inp2
			continue
		case 4:
			fmt.Println("Qual o maximo de numeros em sequencia desejados?")
			inp2, err := g.readInput(1, 100)
			if err != nil {
				g.errorHandler(err)
				continue
			}
			g.settings.MaxSeq = inp2
			continue
		case 5:
			fmt.Println("Quantos jogos deseja gerar? (Max: 10 jogos)")
			inp2, err := g.readInput(1, 10)
			if err != nil {
				g.errorHandler(err)
				continue
			}
			g.betQuantity = inp2
			continue
		case 6:
			g.bet, err = lotto.NewBet(g.settings)
			if err != nil {
				g.errorHandler(err)
				continue
			}
			g.makeBet()
		case 7:
			g.Start()
		case 8:
			os.Exit(0)
		default:
			continue
		}
		break
	}

}

func (g *Game) makeBet() {
	for {
		g.clearScr()
		g.savedBets = []string{}

		fmt.Println("=====================================================")
		fmt.Println()
		fmt.Println("                   JOGOS GERADOS")
		fmt.Println()
		fmt.Println("=====================================================")
		fmt.Println()
		for i := 0; i < g.betQuantity; i++ {
			if err := g.bet.Generate(); err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println(g.bet)
			g.savedBets = append(g.savedBets, g.bet.String())
		}
		fmt.Println()
		fmt.Println("O que deseja fazer?")
		fmt.Println()
		fmt.Printf("\t%d) Gerar mais (%d) jogos\n", 1, g.betQuantity)
		fmt.Printf("\t%d) Salvar os jogos\n", 2)
		fmt.Println()
		fmt.Printf("\t%d) Configurar a (%s) novamente\n", 3, g.settings.Lt)
		fmt.Printf("\t%d) Voltar para a tela inicial\n", 4)
		fmt.Println()
		fmt.Printf("\t%d) Sair\n", 5)

		inp, err := g.readInput(1, 5)
		if err != nil {
			g.errorHandler(err)
			continue
		}
		switch inp {
		case 1:
			g.makeBet()
		case 2:
			err := g.saveFile()
			if err != nil {
				g.errorHandler(err)
				continue
			}
		case 3:
			g.configGame()
		case 4:
			g.Start()
		case 5:
			os.Exit(0)
		default:
			continue
		}
	}

}

// HELPER FUNCTIONS
func (g *Game) readInput(minValue, maxValue int) (int, error) {
	var input int
	fmt.Scan(&input)
	if input < minValue || input > maxValue {
		return 0, fmt.Errorf("opcao nao encontrada, favor tentar novamente")
	}
	return input, nil
}

func (g *Game) clearScr() {
	fmt.Println("\033[2J")
	fmt.Println("\033[H")
	c := exec.Command("cmd", "/c", "cls")
	c.Stdout = os.Stdout
	c.Run()
}

func (g *Game) errorHandler(err error) {
	fmt.Println()
	fmt.Println("ERRO:")
	fmt.Println()
	fmt.Printf("----> %s\n", err)
	fmt.Println()
	fmt.Println("Aguarde 5 segundos")
	for i := 1; i < 5; i++ {
		fmt.Printf("%d... ", i)
		time.Sleep(1 * time.Second)
	}
}

func (g *Game) saveFile() error {
	file, err := os.OpenFile("jogos.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, v := range g.savedBets {
		file.WriteString(time.Now().Format("02/01/2006 15:04:05") + v + "\n")
	}
	return nil
}

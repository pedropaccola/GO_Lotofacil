package lotto

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"strings"
)

type BetSettings struct {
	Lt      LottoName
	Bet     int
	MaxEven int
	MaxOdd  int
	MaxSeq  int
}

type Bet struct {
	settings BetSettings

	min          int
	max          int
	possibleBets []int
	numbers      []int
}

func (b *Bet) String() string {
	strArr := make([]string, len(b.numbers))
	for i, v := range b.numbers {
		strArr[i] = strconv.Itoa(v)
	}
	str := strings.Join(strArr, ", ")

	return fmt.Sprintf("Aposta de %d dezenas: %s", b.settings.Bet, str)
}

func NewBet(s BetSettings) (*Bet, error) {
	b := &Bet{
		settings: s,
	}

	switch b.settings.Lt {
	case 1: //lotofacil
		b.min = 1
		b.max = 25
		b.possibleBets = []int{15, 16, 17, 18, 19, 20}
	case 2: //lotomania
		b.min = 1
		b.max = 100
		b.possibleBets = []int{50}
	case 3: //megasena
		b.min = 1
		b.max = 60
		b.possibleBets = []int{6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	case 4: //quina
		b.min = 1
		b.max = 80
		b.possibleBets = []int{5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	default: //erro
		return nil, fmt.Errorf("opcao (%s) nao disponivel", b.settings.Lt)
	}
	if err := b.validBet(); err != nil {
		return nil, err
	}
	return b, nil
}

func (b *Bet) validBet() error {
	// Basic validation
	if b.settings.MaxEven+b.settings.MaxOdd < b.settings.Bet {
		return fmt.Errorf(
			"impossivel aposta de (%d) pares e (%d) impares com (%d) dezenas",
			b.settings.MaxEven, b.settings.MaxOdd, b.settings.Bet)
	}
	if b.settings.MaxSeq < 1 || b.settings.MaxEven < 0 || b.settings.MaxOdd < 0 {
		return fmt.Errorf(
			"impossivel sequencia de 0 digitos ou valores negativos")
	}

	// Check if bet is in possible range
	for _, bet := range b.possibleBets {
		if b.settings.Bet == bet {
			return nil
		}
	}

	return fmt.Errorf(
		"impossivel apostar (%d) dezenas na %s", b.settings.Bet,
		b.settings.Lt)

}

func (b *Bet) Generate() error {
	b.numbers = make([]int, 0, b.settings.Bet)
	found := make(map[int]bool)

	for len(b.numbers) <= b.settings.Bet {
		i := rand.Intn(b.max) + 1
		// Map uniques and check if generated all possible values.
		if found[i] {
			if len(found) < b.max {
				continue
			} else {
				return fmt.Errorf("nao foi possivel gerar o jogo")
			}
		}
		found[i] = true
		if !b.generateConstraints(i) {
			continue
		} else {
			b.numbers = append(b.numbers, i)
		}
	}
	sort.Slice(b.numbers, func(i, j int) bool {
		return b.numbers[i] < b.numbers[j]
	})
	return nil
}

func (b *Bet) generateConstraints(check int) bool {
	evens := 0
	odds := 0
	seqs := 1
	count := 0
	copyArr := make([]int, 0, len(b.numbers)) //Copy of b.numbers

	// First number generated
	if len(b.numbers) == 0 {
		if check%2 == 0 {
			if b.settings.MaxEven > 0 {
				return true
			}
		} else {
			if b.settings.MaxOdd > 0 {
				return true
			}
		}
		return false
	}

	// Check odds/evens
	for _, v := range b.numbers {
		copyArr = append(copyArr, v)
		if v%2 == 0 {
			evens++
		} else {
			odds++
		}
	}
	if check%2 == 0 && evens >= b.settings.MaxEven {
		return false
	}
	if check%2 != 0 && odds >= b.settings.MaxOdd {
		return false
	}

	// Check sequences
	copyArr = append(copyArr, check)
	sort.Slice(copyArr, func(i, j int) bool {
		return copyArr[i] < copyArr[j]
	})
	for i := range copyArr {
		if i != 0 && copyArr[i] == copyArr[i-1]+1 {
			count++
		} else {
			count = 1
		}

		if count > seqs {
			seqs = count
		}
	}
	if seqs > b.settings.MaxSeq {
		return false
	}
	return true
}

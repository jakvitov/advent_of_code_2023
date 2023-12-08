package day7

import (
	"advent_of_code_2023/day1"
	"advent_of_code_2023/day2"
	"strconv"
)

const INPUT_FILE string = "day7/test_input.txt"

// For easily creating partially ordered set among combinations
const (
	ACE          int = 6
	FOURCE       int = 5
	FULL_HOUSE   int = 4
	TRIPLE       int = 3
	TWO_PAIR     int = 2
	ONE_PAIR     int = 1
	HIGHEST_CARD int = 0
)

type card int

func getCard(char int32) card {
	if day1.IsNum(char) {
		//Char to int conversion rate
		return card(char - 48)
	}
	switch char {
	case 'T':
		return card(10)
	case 'J':
		return card(11)
	case 'Q':
		return card(12)
	case 'K':
		return card(13)
	case 'A':
		return card(14)
	default:
		panic("Not valid character in card.")
	}
}

type hand struct {
	bid         int
	cards       []card
	combination int
}

func getCombination(cards []card) int {
	cardsMap := make(map[card]int, 5)
	//Creating a map with card to its frequency
	for _, crd := range cards {
		val, found := cardsMap[crd]
		if found {
			cardsMap[crd] = val + 1
		} else {
			cardsMap[crd] = 1
		}
	}
	maxCount := 1
	secondMaxCount := 1
	for _, value := range cardsMap {
		if value > maxCount {
			secondMaxCount = maxCount
			maxCount = value
		} else if value > secondMaxCount {
			secondMaxCount = value
		}
	}
	switch maxCount {
	case 5:
		return ACE
	case 4:
		return FOURCE
	case 3:
		if secondMaxCount == 2 {
			return FULL_HOUSE
		}
		return TRIPLE
	case 2:
		if secondMaxCount == 2 {
			return TWO_PAIR
		}
		return ONE_PAIR
	default:
		return HIGHEST_CARD
	}
}

func parseHand(line string) *hand {
	halves := day2.TokenizeString(line, ' ')
	bidSize, _ := strconv.Atoi(halves[1])
	cards := make([]card, 5)
	for i, char := range halves[0] {
		cards[i] = getCard(char)
	}
	return &hand{
		bid:         bidSize,
		cards:       cards,
		combination: getCombination(cards),
	}
}

// h1 > h2
func isFirstHigher(first, second hand) bool {
	if first.combination == second.combination {
		for i := 0; i < 5; i++ {
			if first.cards[i] == second.cards[i] {
				continue
			}
			return first.cards[i] > second.cards[i]
		}
	}
	return first.combination > second.combination
}

// Merge step of the whole operation
func merge(hnds1, hnds2 []hand) []hand {
	result := make([]hand, len(hnds1)+len(hnds2))
	index1 := 0
	index2 := 0
	for index1 < len(hnds1) && index2 < len(hnds2) {
		if !isFirstHigher(hnds1[index1], hnds2[index2]) {
			result[index1+index2] = hnds1[index1]
			index1 += 1
		} else {
			result[index1+index2] = hnds2[index2]
			index2 += 1
		}
	}
	//We append the rest of the remaining array at the end of the result
	for index1 != len(hnds1) || (index2 != len(hnds2)) {
		if index1 < len(hnds1) {
			result[index1+index2] = hnds1[index1]
			index1 += 1
		} else {
			result[index1+index2] = hnds2[index2]
			index2 += 1
		}
	}
	return result
}

// Implementation of merge sort to n*lon(n) sort the field
func sort(hnds []hand) []hand {
	if len(hnds) == 1 {
		return hnds
	}
	half := len(hnds) / 2
	return merge(sort(hnds[:half]), sort(hnds[half:]))
}

func GetTotalWinnings() int64 {
	result := int64(0)
	lines := day1.ReadFileAsLines(INPUT_FILE)
	hands := make([]hand, len(lines))
	for i, line := range lines {
		hands[i] = *parseHand(line)
	}
	hands = sort(hands)
	for i, hnd := range hands {
		result += int64(hnd.bid * (i + 1))
	}
	return result
}

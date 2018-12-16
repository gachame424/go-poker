package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

var drawnNumAllArray []int
var quotientArray []int
var remainderArray []int

func main() {
	var drawnNumArray []int
	drawnNumArray = drawNums(drawnNumArray, 5)
	display(drawnNumArray, 0)
	drawnNumArray = changeCard(drawnNumArray)
	display(drawnNumArray, 1)
	fmt.Println(checkPoker())
}

func drawNums(drawnNumArray []int, count int) []int {
	var drawnNum = 0
	for i := 0; i < count; i++ {
		drawnNum = getPokerNum()
		drawnNumAllArray = append(drawnNumAllArray, drawnNum)
		drawnNumArray = append(drawnNumArray, drawnNum)
	}
	return drawnNumArray
}

func getPokerNum() int {
	const MAX = 52
	rand.Seed(time.Now().UnixNano())
	var drawnNum = 0
	for {
		drawnNum = rand.Intn(MAX)
		if contains(drawnNum) == false {
			break
		}
		if contains(drawnNum) == true {
			continue
		}
	}
	return drawnNum
}

func contains(num int) bool {
	for _, v := range drawnNumAllArray {
		if num == v {
			return true
		}
	}
	return false
}

func display(drawnNumArray []int, checkPokerFlg int) {
	suits := [4]string{"♠︎", "♣︎", "♢", "♡"}
	num := [13]string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	fmt.Println()
	fmt.Println(" No.1   No.2   No.3   No.4   No.5")
	fmt.Println("┌────┐ ┌────┐ ┌────┐ ┌────┐ ┌────┐")
	sort.Sort(sort.IntSlice(drawnNumArray))
	for i := 0; i < 5; i++ {
		var quotient = drawnNumArray[i] / 4
		var remainder = drawnNumArray[i] % 4
		var displayNum = num[quotient]
		if displayNum != "10" {
			displayNum = " " + displayNum
		}
		fmt.Printf("|%s %s| ", suits[remainder], displayNum)
		if checkPokerFlg == 1 {
			quotientArray = append(quotientArray, quotient)
			remainderArray = append(remainderArray, remainder)
		}
	}
	fmt.Println()
	fmt.Println("└────┘ └────┘ └────┘ └────┘ └────┘")
}

func changeCard(drawnNumArray []int) []int {
	fmt.Println("変更するカードNoは？？")
	fmt.Println("(複数の場合はカンマ区切りで)")
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	stdinCardNo := strings.TrimSpace(stdin.Text())
	if regexp.MustCompile("^$").Match([]byte(stdinCardNo)) {
		return drawnNumArray
	}
	if !regexp.MustCompile("^(\\d,){0,4}\\d$").Match([]byte(stdinCardNo)) {
		fmt.Println("ちゃんと指示に従え。馬鹿野郎。")
		os.Exit(0)
	}
	changeCardNo := strings.Split(stdinCardNo, ",")
	var cardNo = 0
	var cardNoArray []int
	for i := 0; i < len(changeCardNo); i++ {
		cardNo, _ = strconv.Atoi(changeCardNo[i])
		cardNoArray = append(cardNoArray, cardNo)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(cardNoArray)))
	drawnNumArray = dropCard(cardNoArray, drawnNumArray)
	return drawNums(drawnNumArray, 5-len(drawnNumArray))
}

func dropCard(cardNoArray []int, drawnNumArray []int) []int {
	var cardNo = 0
	for i := 0; i < len(cardNoArray); i++ {
		cardNo = cardNoArray[i]
		if cardNo < 0 || 5 < cardNo {
			fmt.Println("No.1~5しかねぇだろ。馬鹿野郎。")
			os.Exit(0)
		}
		cardNo = cardNo - 1
		drawnNumArray = unset(drawnNumArray, cardNo)
	}
	return drawnNumArray
}

func unset(s []int, i int) []int {
	if i >= len(s) {
		return s
	}
	return append(s[:i], s[i+1:]...)
}

func checkPoker() string {
	if isFullHouse() {
		return "FullHouse"
	}
	if isFourCard() {
		return "FourCard"
	}
	if isThreeCard() {
		return "ThreeCard"
	}
	if isRoyalStraightFlush() {
		return "RoyalStraightFlush"
	}
	if isStraightFlush() {
		return "StraightFlush"
	}
	if isStraight() {
		return "Straight"
	}
	if isFlush() {
		return "Flush"
	}
	if isTwoPair() {
		return "TwoPair"
	}
	if isOnePair() {
		return "OnePair"
	}
	return "HighCard"
}

func isOnePair() bool {
	if quotientArray[0] == quotientArray[1] ||
		quotientArray[1] == quotientArray[2] ||
		quotientArray[2] == quotientArray[3] ||
		quotientArray[3] == quotientArray[4] {
		return true
	}
	return false
}

func isTwoPair() bool {
	if (quotientArray[0] == quotientArray[1] &&
		(quotientArray[2] == quotientArray[3] ||
			quotientArray[3] == quotientArray[4])) ||
		quotientArray[1] == quotientArray[2] &&
			quotientArray[3] == quotientArray[4] {
		return true
	}
	return false
}

func isThreeCard() bool {
	if (quotientArray[0] == quotientArray[1] &&
		quotientArray[1] == quotientArray[2]) ||
		(quotientArray[1] == quotientArray[2] &&
			quotientArray[2] == quotientArray[3]) ||
		(quotientArray[2] == quotientArray[3] &&
			quotientArray[3] == quotientArray[4]) {
		return true
	}
	return false
}

func isStraight() bool {
	if quotientArray[0]+1 == quotientArray[1] &&
		quotientArray[1]+1 == quotientArray[2] &&
		quotientArray[2]+1 == quotientArray[3] &&
		quotientArray[3]+1 == quotientArray[4] {
		return true
	}

	return false
}

func isFlush() bool {
	if remainderArray[0] == remainderArray[1] &&
		remainderArray[1] == remainderArray[2] &&
		remainderArray[2] == remainderArray[3] &&
		remainderArray[3] == remainderArray[4] {
		return true
	}
	return false
}

func isFullHouse() bool {
	if (quotientArray[0] == quotientArray[1] &&
		quotientArray[1] == quotientArray[2] &&
		quotientArray[3] == quotientArray[4]) ||
		(quotientArray[0] == quotientArray[1] &&
			quotientArray[2] == quotientArray[3] &&
			quotientArray[3] == quotientArray[4]) {
		return true
	}
	return false
}

func isFourCard() bool {
	if quotientArray[0] == quotientArray[1] &&
		quotientArray[2] == quotientArray[3] &&
		quotientArray[0] == quotientArray[2] {
		return true
	}
	if quotientArray[1] == quotientArray[2] &&
		quotientArray[3] == quotientArray[4] &&
		quotientArray[1] == quotientArray[3] {
		return true
	}
	return false
}

func isStraightFlush() bool {
	if isStraight() && isFlush() {
		return true
	}
	return false
}

func isRoyalStraightFlush() bool {
	if !isFlush() {
		return false
	}
	if quotientArray[0] == 0 &&
		quotientArray[1] == 9 &&
		quotientArray[2] == 10 &&
		quotientArray[3] == 11 &&
		quotientArray[4] == 12 {
		return true
	}
	return false
}

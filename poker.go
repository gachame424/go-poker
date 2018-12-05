package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
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
	drawnNumArray = drawNums(drawnNumArray, 5-len(drawnNumArray))
	display(drawnNumArray, 1)
	checkPoker(drawnNumArray)
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
	fmt.Println("変更するカードNoは？？(複数の場合はカンマ区切りで)")
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	stdinCardNo := stdin.Text()
	cardNoArrayTmp := strings.Split(stdinCardNo, ",")
	var cardNo = 0
	var cardNoArray []int
	for i := 0; i < len(cardNoArrayTmp); i++ {
		cardNo, _ = strconv.Atoi(cardNoArrayTmp[i])
		cardNoArray = append(cardNoArray, cardNo)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(cardNoArray)))
	return dropCard(cardNoArray, drawnNumArray)
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

func checkPoker(drawnNumArray []int) {}

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func calculateMinimumBuses() {
	var numFamilies int
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Input the number of families: ")
	fmt.Scanln(&numFamilies)

	fmt.Print("Input the number of members in the family (separated by a space): ")
	inputFamilies, _ := reader.ReadString('\n')
	inputFamilies = strings.TrimSpace(inputFamilies)

	// Memproses input menjadi slice integer
	familyStrings := strings.Fields(inputFamilies)
	if len(familyStrings) != numFamilies {
		fmt.Println("Input must be equal with count of family")
		return
	}

	familyMembers := make([]int, len(familyStrings))
	for i, fam := range familyStrings {
		val, err := strconv.Atoi(fam)
		if err != nil || val < 1 {
			fmt.Println("Invalid input. Please enter valid integers greater than 0.")
			return
		}
		familyMembers[i] = val
	}

	sort.Sort(sort.Reverse(sort.IntSlice(familyMembers)))

	buses := 0
	left, right := 0, len(familyMembers)-1

	for left <= right {
		if familyMembers[left]+familyMembers[right] <= 4 {
			right--
		}
		left++
		buses++
	}

	fmt.Printf("Minimum bus required is: %d\n", buses)
}

func main() {
	calculateMinimumBuses()
}

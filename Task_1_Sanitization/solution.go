package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	//"log"
	)



func main() {
	file, _ := os.ReadFile("./input/input19.txt")
	allText := string(file)
	text := strings.Split(allText, "\n")
	
	for index, line := range(text) {		//to access the cell you gotta use "text[index]" because "line" is just a value
		text[index] = strings.TrimSpace(line)		//to take out new line --> strings.TrimRight(yourString, "\n")
	}

	stringN := strings.Split(text[0], " ")[0]			//WORDS TO CHECK
	stringM := strings.Split(text[0], " ")[1]			//BANNED WORDS

	N, _ := strconv.Atoi(strings.TrimSpace(stringN))
	M, _ := strconv.Atoi(strings.TrimSpace(stringM))

	var wordsToCheckSet = make([]string, N)		 
	
	for i:=0; i<N; i++ {		// CELLS
		for j:=M+1; j<=len(text)-1; j++ {		// WORDS, "-1" because "len" starts from 1
			if i == j-M-1 {
				wordsToCheckSet[i] = text[j]
			}
		} 
	}

	var bannedWordsSet = make([]string, M)
	
	for i:=0; i<M; i++ {		// CELLS
		for j:=1; j<=M+1; j++ {		// WORDS skipping the first line
			if i == j-1 {
				bannedWordsSet[i] = text[j]
			}
		} 
	}
	
	//fmt.Println("BannedWordsSet: \n", bannedWordsSet, "\n\n", "wordsToCheckSet: \n", wordsToCheckSet)
	counter := 0
	for i:=0; i<N; i++ {
		for j:=0; j<M; j++ {
			if strings.Contains(wordsToCheckSet[i], bannedWordsSet[j]) == true {
				counter = counter +1
			}	
		}
		if counter == 0 {
			fmt.Println("SAFE")
		} else {
			fmt.Println("BANNED")
		}
		counter = 0
	}
}

	
	  

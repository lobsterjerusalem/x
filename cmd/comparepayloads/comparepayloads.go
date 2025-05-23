package main

import (
	"os"
	"strconv"
	"strings"
	"fmt"
)

func main(){
	useJaccard := false
	threshold := .8
	if len(os.Args) < 3 {
		fmt.Println("Usage: comparehexpayloads <file one> <file two> <optional: jaccard similarity threshold (float) e.g: .8>")
		fmt.Println("Note: The files should be mostly the same this was made to compare hex of payloads that are one byte off in order to find length offsets.")
		fmt.Println("The output is derived by running xxd on the payloads and only comparing the hex")
		os.Exit(1)
	}

	if len(os.Args) == 4 {
		useJaccard = true
		f, err := strconv.ParseFloat(os.Args[3], 64)
		if err != nil {
			fmt.Printf("ERR: the jaccard arg should be a float, a good number might be .8, err=%s", err)
			os.Exit(1)
		}
		threshold = f
	}

	fileOne := os.Args[1]
	fileTwo := os.Args[2]
    _,err := os.Stat(fileOne) 
	if err != nil {
		fmt.Printf("Could not open %s, err=%s\n", fileOne, err)
		os.Exit(1)
	}
    _,err = os.Stat(fileTwo) 
	if err != nil {
		fmt.Printf("Could not open %s, err=%s\n", fileTwo, err)
		os.Exit(1)
	}

	fileOneData, err := os.ReadFile(fileOne)
	if err != nil {
		fmt.Printf("Could not read %s, err=%s\n", fileOne, err)
		os.Exit(1)
	}
	fileTwoData, err := os.ReadFile(fileTwo)
	if err != nil {
		fmt.Printf("Could not read %s, err=%s\n", fileTwo, err)
		os.Exit(1)
	}

	if len(fileOneData) < len(fileTwoData) {
		fmt.Println("Note: The second file passed should be the smaller file for best results")
		os.Exit(1)
	}

	fileOneLines := strings.Split(string(fileOneData), "\n")
	fileTwoLines := strings.Split(string(fileTwoData), "\n")

	skipped := 0
	mismatched := 0
	total := len(fileOneLines)
	for n := range len(fileOneLines) {
		fOneStr := string(fileOneLines[n])
		fTwoStr := string(fileTwoLines[n])

		// too short to substring out
		if len(fOneStr) < 8 || len(fTwoStr) < 8 {
			skipped++
			continue
		}

		sim := jaccardSimilarity(fOneStr, fTwoStr) 
		if useJaccard {
			if sim < threshold {
				fmt.Printf("\tjaccard: %f for line %d\n%s\n%s\n\n", sim, n, fOneStr, fTwoStr)
			}
			continue
		}

		// we anticipate the fact that one byte being off
		if strings.Contains(fOneStr, fTwoStr[4:len(fTwoStr)-2]) {
			continue
		}
		mismatched++
		fmt.Printf("%s\n%s\n\n", fOneStr, fTwoStr)
	}

	fmt.Printf("Summary: total: %d, skipped %d, mismatched: %d\n", total, skipped, mismatched)
}

// This is just here, not really using it now, but may later as an option
func jaccardSimilarity(a, b string) float64 {
	setA := make(map[string]struct{})
	setB := make(map[string]struct{})

	for _, word := range strings.Fields(a) {
		setA[word] = struct{}{}
	}
	for _, word := range strings.Fields(b) {
		setB[word] = struct{}{}
	}

	intersection := 0
	for word := range setA {
		if _, exists := setB[word]; exists {
			intersection++
		}
	}
	union := len(setA) + len(setB) - intersection

	if union == 0 {
		return 0
	}
	return float64(intersection) / float64(union)
}


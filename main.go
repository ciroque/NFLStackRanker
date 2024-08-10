package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// Item represents an item with a name and a score
type Item struct {
	Name  string
	Score int
}

// CompareItems compares two items and adjusts their scores based on the comparison result.
func CompareItems(a, b *Item) int {
	clearScreen()
	fmt.Printf("Compare %s vs %s\n", a.Name, b.Name)
	fmt.Println("Which is better? (1 for", a.Name, ", 2 for", b.Name, ")")
	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		a.Score++
		return 1
	case 2:
		b.Score++
		return -1
	default:
		fmt.Println("Invalid input. Skipping comparison.")
		return 0
	}
}

// MergeSortRanks recursively sorts and ranks items using pairwise comparisons
func MergeSortRanks(items []*Item) []*Item {
	if len(items) <= 1 {
		return items
	}

	// Split the list into two halves
	mid := len(items) / 2
	left := MergeSortRanks(items[:mid])
	right := MergeSortRanks(items[mid:])

	// Merge the two halves with pairwise comparisons
	return merge(left, right)
}

// Merge two halves, comparing each pair
func merge(left, right []*Item) []*Item {
	result := make([]*Item, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if CompareItems(left[i], right[j]) == 1 {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Append remaining items
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

// clearScreen clears the terminal screen
func clearScreen() {
	switch runtime.GOOS {
	case "linux", "darwin": // Unix-like systems
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows": // Windows systems
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		fmt.Println("Unsupported platform, cannot clear screen.")
	}
}

func main() {
	// List of NFL teams to rank
	items := []*Item{
		{Name: "Arizona Cardinals", Score: 0},
		{Name: "Atlanta Falcons", Score: 0},
		{Name: "Baltimore Ravens", Score: 0},
		{Name: "Buffalo Bills", Score: 0},
		{Name: "Carolina Panthers", Score: 0},
		{Name: "Chicago Bears", Score: 0},
		{Name: "Cincinnati Bengals", Score: 0},
		{Name: "Cleveland Browns", Score: 0},
		{Name: "Dallas Cowboys", Score: 0},
		{Name: "Denver Broncos", Score: 0},
		{Name: "Detroit Lions", Score: 0},
		{Name: "Green Bay Packers", Score: 0},
		{Name: "Houston Texans", Score: 0},
		{Name: "Indianapolis Colts", Score: 0},
		{Name: "Jacksonville Jaguars", Score: 0},
		{Name: "Kansas City Chiefs", Score: 0},
		{Name: "Las Vegas Raiders", Score: 0},
		{Name: "Los Angeles Chargers", Score: 0},
		{Name: "Los Angeles Rams", Score: 0},
		{Name: "Miami Dolphins", Score: 0},
		{Name: "Minnesota Vikings", Score: 0},
		{Name: "New England Patriots", Score: 0},
		{Name: "New Orleans Saints", Score: 0},
		{Name: "New York Giants", Score: 0},
		{Name: "New York Jets", Score: 0},
		{Name: "Philadelphia Eagles", Score: 0},
		{Name: "Pittsburgh Steelers", Score: 0},
		{Name: "San Francisco 49ers", Score: 0},
		{Name: "Seattle Seahawks", Score: 0},
		{Name: "Tampa Bay Buccaneers", Score: 0},
		{Name: "Tennessee Titans", Score: 0},
		{Name: "Washington Commanders", Score: 0},
	}

	// Perform the ranking
	rankedItems := MergeSortRanks(items)

	// Print the ranked items
	fmt.Println("Ranked NFL Teams:")
	for i, item := range rankedItems {
		fmt.Printf("%d: %s (Score: %d)\n", i+1, item.Name, item.Score)
	}
}

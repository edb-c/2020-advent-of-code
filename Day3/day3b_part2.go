package main

/*
1 - Get it running - check
2 - Get it running correctly & submit - check - Yay!
3 - Refactor - tbd  - It's still ugly, I know...
*/

/*
--- Day 3: Toboggan Trajectory ---
From your starting position at the top-left,
check the position that is right 3 and down 1.
Then, check the position that is right 3 and down 1 from there,
and so on until you go past the bottom of the map.

The locations you'd check in the above example are marked here with O
where there was an open square and X where there was a tree.

*/

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type TraversePattern struct {
	Right int
	Down  int
}

func main() {
	lines := GetLines("./puzzle_input3.txt")

	treeCountMultiplied := 1

	traversePatterns := []TraversePattern{
		TraversePattern{Right: 1, Down: 1},
		TraversePattern{Right: 3, Down: 1},
		TraversePattern{Right: 5, Down: 1},
		TraversePattern{Right: 7, Down: 1},
		TraversePattern{Right: 1, Down: 2},
	}

	for _, aPattern := range traversePatterns {
		treeCountMultiplied *= checkSlopes(lines, aPattern)
		fmt.Println("aPattern", aPattern, treeCountMultiplied)
	}

	fmt.Println("treeCountMultiplied", treeCountMultiplied)
} //end main

func GetLines(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	return strings.Split(string(data), "\n")
}

func checkSlopes(lines []string, traversePattern TraversePattern) int {
	treeCount := 0
	treeSpotY := 0
	isTree := `#`

	// fmt.Println("Length", len(lines))

	for treeSpotX := 0; treeSpotX < len(lines); treeSpotX += traversePattern.Down {

		aLine := string(lines[treeSpotX])

		if string(aLine[treeSpotY]) == isTree {

			// fmt.Println(treeSpotX,treeSpotY, string(aLine[treeSpotY]),isTree, string(aLine[treeSpotY]) == isTree)

			treeCount++
		}
		//	fmt.Println("Test", treeSpotX,treeSpotY , traversePattern.Right , len(aLine), (treeSpotY + traversePattern.Right) % len(aLine))
		treeSpotY = (treeSpotY + traversePattern.Right) % len(aLine)
	} ///end for i
	fmt.Println("treeCount", treeCount)
	return treeCount
}

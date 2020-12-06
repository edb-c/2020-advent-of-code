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

func main() {
	lines := GetLines("./puzzle_input3.txt")
	treeCount := 0
	treeSpotX := 3
	treeSpotY := 0
	isTree := `#`

	// fmt.Println("Length", len(lines))
	for i, aLine := range lines {
		//fmt.Println("i", i, treeSpotX,treeSpotY )
			if string(lines[i][treeSpotY]) == isTree {
				treeCount++
				fmt.Println(string(lines[i]), i,",",treeSpotY,"   For treeSpotX, treeSpotY", treeSpotX, treeSpotY,"Tree",string(lines[i][treeSpotY]) )
			} 

		treeSpotY = (treeSpotX + treeSpotY) % len(aLine) 

	} ///end for i
	fmt.Println("treeCount", treeCount)
} //end main

func GetLines(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	return strings.Split(string(data), "\n")
}

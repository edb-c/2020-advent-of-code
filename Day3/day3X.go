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

UPDATE - Initial thoughts were incorrect because value of treeSpotX did not reset to zero
once it reached end of row line.
Initial thoughts:
Hmm, needed steps seem to be:
 read the file/get lines

	check each character, to see if its a X - tree

	line i	find tree in character
	0 		3
	1		3
	2		6
	3		9
	4		12
​
	Pattern ?

*/

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	lines := GetLines("./puzzle_input3.txt")
	treeCount := 0
	treeSpotX := 0
	treeSpotY := 0
	isTree := `#`
	// fmt.Println("Length", len(lines))

	for i, aLine := range lines {
		//fmt.Println("i", i, treeSpotX,treeSpotY )

		for i2 := 0; i2 < len(aLine); i2 = i2 + 3 {

			if i >= 1 {
				if treeSpotY > len(aLine)-1 {
					treeSpotX = 0
					treeSpotY = 0
				}

				//	fmt.Println("line", "column", i, i2, string(lines[i][i2]))
				if i2 == treeSpotY && string(lines[i][i2]) == isTree {
					treeCount = treeCount + 1
					fmt.Println(string(lines[i]), i, ",", i2, "   For treeSpotX, treeSpotY", treeSpotX, treeSpotY, i2, "Tree", string(lines[i][i2]))

				}

			} //end if i1> 0
			//fmt.Println("for line i",i, "treeSpotX, treeSpotY",treeSpotX, treeSpotY,"i,i2",i,i2,treeCount)

			//	fmt.Println("For treeSpotX, treeSpotY", treeSpotX, treeSpotY, "line", i,i2, string(lines[i][i2]))
		} //end for i2

		treeSpotX = treeSpotX + 1
		treeSpotY = treeSpotY + 3

	} ///end for i
	fmt.Println("treeCount", treeCount)
} //end main

func GetLines(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	return strings.Split(string(data), "\n")
}

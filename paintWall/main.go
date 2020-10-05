package main
import "fmt"

func main() {
	// 0
	fmt.Println(paintWall([]int{}))
	fmt.Println(paintWall([]int{0}))
	fmt.Println(paintWall([]int{1}))

	// 2
	fmt.Println(paintWall([]int{1,0,0,1,0,1}))
	fmt.Println(paintWall([]int{0,0,0,1,0,1}))

	// 1
	fmt.Println(paintWall([]int{1,1,0,1,0,0}))

	// 4
	fmt.Println(paintWall([]int{1,1,0,1,0,1,0,0,0,1,0,1}))

}

func paintWall(colors []int) (changes int) {
	colorsLen := len(colors)

	// 递归边界条件
	if colorsLen <= 1 {
		return 0
	}

	// 如果第一个为红，则对所有蓝色，全刷红。
	if colors[0] == 0 {
		for j:=1; j<colorsLen; j++{
			if colors[j] == 1 {
				changes += 1
			}
		}
		return changes
	}
	
	var j int = 0
	var k int = 0
	var toAllRedChanges int = 0
	var thisPartToBlueChanges int = 0

	// 如果第一个为蓝色
	// 走递归，算出刷红色和蓝色的最小值

	// 先到第一个红颜色块
	for j=1; j<colorsLen; j++ {
		if colors[j] == 0 {
			toAllRedChanges = paintWall(colors[j:])
			break
		}
	}
	
	if j==colorsLen{
		return 0
	}

	for k=j+1;k<colorsLen;k++ {
		if colors[k] == 1 {
			thisPartToBlueChanges = k - j + paintWall(colors[k:])
			break
		}
	}

	// 取两者较小值
	if toAllRedChanges < thisPartToBlueChanges {
		changes = toAllRedChanges
	} else {
		changes = thisPartToBlueChanges
	}

	return changes
}
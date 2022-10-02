package main

import (
	"fmt"
	"sort"
)

func main() {
	//入力を受ける
	var numberOfSheets int
	fmt.Scan(&numberOfSheets)

	mochies := make([]int, numberOfSheets)
	for i := 0; i < numberOfSheets; i++ {
		fmt.Scan(&mochies[i])
	}

	//降順で並べ替える
	sort.Sort(sort.Reverse(sort.IntSlice(mochies)))

	count := 0
	stock := -1
	for _, v := range mochies {
		//同じ数字だったらスキップ
		if v == stock {
			continue
		}
		//違う数字だったらcountを増やす
		count++
		stock = v
	}

	fmt.Println(count)
}

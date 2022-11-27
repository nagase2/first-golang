package main

import (
	"fmt"
)

func main() {

	one := []int32{2, 1, 5, 3, 4}
	two := []int32{2, 5, 1, 3, 4}
	minimumBribes(one)
	minimumBribes(two)
}

func sumSlice(arr []int) int {
	sum := 0
	for _, n := range arr {
		sum += n
	}
	return sum
}

// お金を払うと横の人を買収して一個前にすすめる
// 買収できるのは一人２回まで。２回以上買収しした人がいたらTooChaoticと出力してそこで処理を終了する
// 最後に、全体で何度買収行為が合ったのかを数えて出力する
func minimumBribes(q []int32) {
	n := len(q)
	bribes := make([]int, n, n)
	i, j := 0, 1

	for i < n-1 && j < n {
		l, r := q[i], q[j]
		// 入れ替え対象か？　2,1  3,4 など
		if l > r {
			bribes[l-1] += 1
			if bribes[l-1] > 2 {
				fmt.Println("Too chaotic")
				return
			}
			q[i], q[j] = q[j], q[i]
			//次に読むところを設定。入れ替えが発生した場合、それが更に後ろに戻される場合を考慮してのこと（1,5,6,3,4など）
			// 毎回最初から読んでもよいがじかんかかるのでこのほうが合理的
			if i == 0 {
				fmt.Println("🐮i:",i,"j:",j)
				i, j = i+1, j+1
			} else {
				fmt.Println("😸i:",i,"j:",j)
				i, j = i-1, j-1
			}
		} else {
			
			fmt.Println("🈲i:",i,"j:",j)
			i, j = i+1, j+1
		}
	}

	sum := sumSlice(bribes)
	fmt.Println(sum)
}

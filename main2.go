package main

import (
	"fmt"
	"time"
)

// func main() {
// 	src := []int{1, 2, 3, 4, 5}
// 	dst := []int{}

// 	for _, s := range src {
// 		go func(s int) {
// 			// 何か(重い)処理をする
// 			result := s * 2

// 			// ここでdstを読み取ってresultを追加している
// 			// dstを読み取った後に他のgoroutineで値が変わっててもその値を上書きしてしまう
// 			dst = append(dst, result)
// 		}(s)
// 	}

// 	time.Sleep(time.Second)
// 	fmt.Println(dst)
// }

func main() {
	ch1 := make(chan int)

	go func() {
		defer close(ch1)
		for {
			ch1 <- 1
		}
	}()

	timeout := time.After(1 * time.Second)

	// このforループを1秒間ずっと実行し続ける
	for {
		select {
		case s := <-ch1:
			fmt.Println(s)
		case <-timeout:
			fmt.Println("time out")
			return
		default:
			fmt.Println("default")
			time.Sleep(time.Millisecond * 100)
		}
	}
}

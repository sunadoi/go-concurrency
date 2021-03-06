package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("what is today's lucky number?")
	// go getLuckyNum()

	// time.Sleep(time.Second * 5) // これを書かないとメインgoroutineが終了して、他のgoroutineの終了を待たずにプログラムが終了する

	// 代わりにWaitGroupを使ってgoroutineの終了を待つ
	// var wg sync.WaitGroup
	// wg.Add(1)

	// go func() {
	// 	defer wg.Done()
	// 	getLuckyNum()
	// }()

	// wg.Wait()

	// チャネルを使って送受信を行う
	c := make(chan int)
	go getLuckyNum(c)

	num := <-c
	fmt.Printf("Today's your lucky number is %d!\n", num)

	close(c)
}

func getLuckyNum(c chan<- int) {
	fmt.Println("...")

	// 占いにかかる時間はランダム
	rand.Seed(time.Now().Unix())
	time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)

	num := rand.Intn(10)
	// fmt.Printf("Today's your lucky number is %d!\n", num)
	c <- num
}
package main

import (
	"fmt"
	"testing"
)

// 查看cpufile，可以看到耗时
// go test -v -run="TestCountAndSay$" -cpuprofile=cpu1.out
// go tool pprof -svg cpu1.out > cpu1.svg
func TestCountAndSay(t *testing.T) {
	for i := 1; i <= 30; i++ {
		fmt.Printf("m: %d,\tres: %s\n", i, CountAndSay(i))
	}
}

// go test -v -run="TestCountAndSay2" -cpuprofile=cpu2.out
// go tool pprof -svg cpu2.out > cpu2.svg
func TestCountAndSay2(t *testing.T) {
	for j := 1; j <= 30; j++ {
		fmt.Printf("n: %d,\tres: %s\n", j, CountAndSay2(j))
	}
}

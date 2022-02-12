package main

import "fmt"

type MemoizedFunc func(int) int

var factorialMemoized MemoizedFunc

func Memoize(functionToExec MemoizedFunc) MemoizedFunc {
	cache := make(map[int]int)

	return func(key int) int {
		if val, found := cache[key]; found {
			return val
		}
		temp := functionToExec(key)
		cache[key] = temp
		return temp
	}

}

func factorial(x int) int {
	if x == 1 {
		return 1
	} else {
		return x * factorialMemoized(x -1)
	}
}

func main() {

	factorialMemoized = Memoize(factorial)

	fmt.Println(factorialMemoized(2))
	fmt.Println(factorialMemoized(3))
	fmt.Println(factorialMemoized(4))
}

package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	fib := 0
	return func() int {
		fib++
		fn_1 := 1
		fn_2 := 0
		for i := 1; i < fib; i++ {
			temp := fn_2
			fn_2 = fn_1
			fn_1 = temp + fn_2
		}
		return fn_2
	}
}

/*
func fibonacci() func() int {
    fn_2, fn_1 := 0, 1
    return func() int {
        res := fn_2
        fn_2, fn_1 = fn_1, fn_2+fn_1
        return res
    }
}
*/

func main() {
	f := fibonacci()
	for i := 0; i < 30; i++ { //10
		fmt.Println(f())
	}
}

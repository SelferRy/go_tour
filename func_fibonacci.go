package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	fib := 0
	return func() int {
		// fmt.Println("FIB", fib)
		/*if fib < 2 {
			// fmt.Println("HERE", fib)
			result := fib
			fib++
			return result
		}*/
		fib++
		var fn_1 int = 1
		var fn_2 int = 0
		for i := 1; i < fib; i++ {
			temp := fn_2 // fib_n2 val
			fn_2 = fn_1
			fn_1 = temp + fn_2
			// fmt.Println("temp = ", temp, "fib_n2 = ", fib_n2, "fib_n1 = ", fib_n1)
		}
		// fmt.Println(fib_n1, fib_n2)
		return fn_2 // fib_n1 + fib_n2
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

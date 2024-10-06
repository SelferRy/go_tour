package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type MyString string

func (s MyString) TwoFirst() []string {
	var res []string
	res = append(res, string(s[0]), string(s[1]))
	// for i :=0; i < 2; i++ {
	//	append(res, string(s[0]), string(s[1])
	return res
}

func (s MyString) ShowTwoFirst() {
	for i := 0; i < 2; i++ {
		fmt.Println(string(s[i]))
		fmt.Println(s[i])
	}
}

// func (s MyString) OwnTwoFirst() {
//	var res []string
//	for i :=0; i < 2; i++ {
//		res = append(res, s[i])
//	}
//}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
	s := MyString("some string")
	fmt.Println(s.TwoFirst())
	s.ShowTwoFirst()
}

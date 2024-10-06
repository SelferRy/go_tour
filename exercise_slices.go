package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	val := uint8((dx + dy) / 2)
	inner := make([]uint8, dx)
	for i := range inner {
		inner[i] = val
	}
	arr := make([][]uint8, dy)
	for i := range arr {
		arr[i] = inner
	}
	return arr
}

func main() {
	pic.Show(Pic)
}

/*func Pic(dx, dy int) [][]uint8 {
	fmt.Println(dx, dy, uint8(dx), uint8(dy))
	val := uint8(dx) // uint8((dx + dy) / 2)
	fmt.Println("val", val)
	inner := make([]uint8, dx)
	fmt.Println("inner-first:", inner)
	for i := range inner {
		inner[i] = val
		fmt.Println("cycle", inner[i])
	}
	fmt.Println("inner-finish:", inner)
	arr := make([][]uint8, dy)
	fmt.Println("arr-first:", arr)
	for i := range arr {
		arr[i] = inner
	}
	fmt.Println("arr-finished:", arr)
	return arr
}*/

package sorttext

import "fmt"

func Calculate(a string, b, c int) {
	if a == "mul" {
		result := b * c
		fmt.Println(result)
	}
	if a == "add" {
		result := b + c
		fmt.Println(result)
	}
	if a == "sub" {
		result := b - c
		fmt.Println(result)
	}
	if a == "div" {
		if c == 0 {
			fmt.Println(" value 'C' not be 0")
			return
		}
		result := float32(float32(b) / float32(c))
		fmt.Println(result)
	}
}

//dãy fibonaci
func FibonacciTo(number int) {
	a := 0
	b := 1
	for a < number {
		fmt.Println(a)
		c := a + b
		a = b
		b = c
	}
}

//sắp xếp
func BubbleSort(sort string, x []int) []int {
	if sort == "bigtosmall" {
		for i := 0; i < len(x); i++ {
			for j := i + 1; j < len(x); j++ {
				if x[i] < x[j] {
					thirdParam := x[i]
					x[i] = x[j]
					x[j] = thirdParam
				}
			}
		}
	}
	if sort == "smalltobig" {
		for i := 0; i < len(x); i++ {
			for j := i + 1; j < len(x); j++ {
				if x[i] > x[j] {
					thirdParam := x[i]
					x[i] = x[j]
					x[j] = thirdParam
				}
			}
		}
	}
	return x
}

func FindMinMaxNumber(find string, x []int) int {
	// var result int
	var result int

	arr := BubbleSort("bigtosmall", x)
	if find == "max" {
		return arr[0]
	}
	if find == "min" {
		return arr[len(x)-1]

	}
	return result
}
func Average(x []int) {
	result := 0
	for _, v := range x {
		result += v
	}
	fmt.Println(result / len(x))
}

//merge sort
func Merge(l, r []int) []int {
	ret := make([]int, 0, len(l)+len(r))
	for len(l) > 0 || len(r) > 0 {
		if len(l) == 0 {
			return append(ret, r...)
		}
		if len(r) == 0 {
			return append(ret, l...)
		}
		if l[0] <= r[0] {
			ret = append(ret, l[0])
			l = l[1:]
		} else {
			ret = append(ret, r[0])
			r = r[1:]
		}
	}
	return ret
}

func MergeSort(s []int) []int {
	if len(s) <= 1 {
		return s
	}
	n := len(s) / 2
	l := MergeSort(s[:n])
	r := MergeSort(s[n:])
	return Merge(l, r)
}
func PrimeNumber(a []int) {
	error := 0
	for _, v := range a {
		if v < 2 {
			error += 1
		}
		for i := 2; i < v/2; i++ {
			if v%i == 0 {
				error += 1
			}
		}
	}
	if error != 0 {
		fmt.Println("không phải dãy số nguyên tố")

	} else {
		fmt.Println("là dãy số  nguyên tố")
	}
}

//quickSort
func Partition(arr []int, low, high int) int {

	p := arr[high]

	for j := low; j < high; j++ {
		if arr[j] < p {
			arr[j], arr[low] = arr[low], arr[j]
			low++
		}
	}

	arr[low], arr[high] = arr[high], arr[low]
	return low
}

func QuickSort(arr []int, low, high int) []int {

	if low > high {
		return arr
	}

	p := Partition(arr, low, high)
	QuickSort(arr, low, p-1)
	QuickSort(arr, p+1, high)
	return arr
}

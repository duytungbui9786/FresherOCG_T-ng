package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/techmaster.vn/helloworld/sorttext"
)

func readText(file string) []int {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("fall")
	}
	arr := strings.Split(string(data), ",")
	var newarray []int
	for _, v := range arr {
		a, _ := strconv.Atoi(v)
		newarray = append(newarray, a)
	}
	return newarray
}

func TestCalculate(x int) (result int) {
	result = x + 2
	return result
}

const (
	nameKey = "result"
)

func main() {
	// server.RunServer()
	var a []int
	a = readText("text.txt")
	// fmt.Println(a)
	sorttext.FibonacciTo(100)
	fmt.Println(sorttext.BubbleSort("bigtosmall", a))
	fmt.Println(sorttext.BubbleSort("smalltobig", a))
	fmt.Println(sorttext.MergeSort(a))
	fmt.Println(sorttext.QuickSort(a, 0, len(a)-1))
	sorttext.PrimeNumber(a)

}

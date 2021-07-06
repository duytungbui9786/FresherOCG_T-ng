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
	var a []int
	a = readText("text.txt")
	// fmt.Println(a)
	// sorttext.FibonacciTo(100)
	// fmt.Println(sorttext.BubbleSort("bigtosmall", a[]))
	// fmt.Println(sorttext.BubbleSort("smalltobig", a))
	fmt.Println(sorttext.QuickSort(a, 0, len(a)-1))
	// server.RunServer()

	// //crawl-data
	// resp, err := http.Get("https://www.imdb.com/chart/top/?ref_=nv_mv_250")
	// if err != nil {
	// 	panic(err)
	// }
	// defer resp.Body.Close()
	// fmt.Println("Response status:", resp.Status)
	// scanner := bufio.NewScanner(resp.Body)
	// for i := 0; scanner.Scan() && i < 2000; i++ {
	// 	res := strings.Contains(scanner.Text(), "collection-detail__product-details text-align-left")
	// 	if res == true {
	// 		// fmt.Println(scanner.Text())
	// 		for i := 0; scanner.Scan() && i < 5; i++ {
	// 			fmt.Println(scanner.Text())
	// 		}
	// 	}
	// }
	// if err := scanner.Err(); err != nil {
	// 	panic(err)
	// }
}

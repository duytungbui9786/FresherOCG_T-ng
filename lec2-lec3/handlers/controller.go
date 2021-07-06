package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

const (
	nameKey = "result"
)

type Message struct {
	result int
}

func SumWithParam(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// return "OKOK"
	// json.NewEncoder(w).Encode("OKOK")
	params := req.URL.Query()
	// get "name" query

	if yourName, ok := params[nameKey]; ok {
		text := strings.Split(yourName[0], ",")
		// var arrConvert []int
		var result int
		for _, val := range text {

			value, _ := strconv.Atoi(val)
			result += value
		}
		// a := Message{result: result}
		json.NewEncoder(w).Encode(result)

	} else {
		fmt.Fprintf(w, "summary = %d", "infinity")
	}
}
func MultiplyWithParam(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	params := req.URL.Query()
	// get "name" query

	if yourName, ok := params[nameKey]; ok {
		text := strings.Split(yourName[0], ",")
		// var arrConvert []int
		result := 1
		for _, val := range text {

			value, _ := strconv.Atoi(val)
			result *= value
		}
		json.NewEncoder(w).Encode(result)
	} else {
		fmt.Fprintln(w, `pleased add params`)
	}
}
func MinuWithParam(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	params := req.URL.Query()
	// get "name" query

	if yourName, ok := params[nameKey]; ok {
		text := strings.Split(yourName[0], ",")
		// var arrConvert []int
		var result int
		value1, _ := strconv.Atoi(text[0])
		value2, _ := strconv.Atoi(text[1])
		result = value1 - value2
		json.NewEncoder(w).Encode(result)
	} else {
		fmt.Fprintln(w, `pleased add params`)
	}
}
func DivideWithParam(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	params := req.URL.Query()
	// get "name" query

	if yourName, ok := params[nameKey]; ok {
		text := strings.Split(yourName[0], ",")
		// var arrConvert []int
		var result float32
		value1, _ := strconv.Atoi(text[0])
		value2, _ := strconv.Atoi(text[1])
		if value2 == 0 {
			json.NewEncoder(w).Encode("infinity")
			return
		}
		result = float32(value1) / float32(value2)
		json.NewEncoder(w).Encode(result)
	} else {
		fmt.Fprintln(w, `pleased add params`)
	}
}

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var abc = []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', ' '}
var length int
var values string
var numbs []string
var fW = 0
var sW = 0
var W = 0
var res []byte

func main() {
	// scanner
	in := bufio.NewScanner(os.Stdin)

	// reading request
	fmt.Println("Введите запрос.")
	fmt.Scanln(&length)
	in.Scan()
	values = in.Text()

	// checking the length of the request and splitting the string
	if length >= 0 && length <= 500 {
		numbs = strings.Split(values, " ")
	} else {
		fmt.Print("Запрос вне диапозона.")
	}

	// value conversion cycle
	for _, value := range numbs {
		val, err := strconv.Atoi(value)
		if err != nil {
			fmt.Print("Ошибка конвертации. ", err)
		}
		if W == 0 {
			res = append(res, valueW(val))
			W = val
			fW = val
		} else {
			sW = val
			if fW <= sW {
				W = sW - fW
				res = append(res, valueW(W))
				fW = val
			} else {
				W = fW - sW
				res = append(res, valueW(W))
				fW = val
			}
		}
	}

	// response output
	fmt.Print(string(res))
}

// getting character by degree
func valueW(i int) byte {
	degree := 0.0
	for math.Pow(2, degree) != float64(i) {
		degree++
	}
	ch := abc[int(degree)]
	return ch
}

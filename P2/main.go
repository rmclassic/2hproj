package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	stdinReader := bufio.NewReader(os.Stdin)
	buff, _, _ := stdinReader.ReadLine()
	opCount, err := strconv.Atoi(string(buff))
	if err != nil {
		println("operation count must be integer")
		return
	}

	result := ReadAndDoOperation(opCount)

	println(result)
}

func ReadAndDoOperation(count int) int {
	stdinReader := bufio.NewReader(os.Stdin)
	buff, _, _ := stdinReader.ReadLine()

	op := NormalizeString(string(buff))
	nums := strings.Split(op, " ")
	result, err := DoOperation(nums)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}

	if count == 1 {
		return result
	} else {
		return result + ReadAndDoOperation(count-1)
	}
}

func NormalizeString(op string) string {
	op = strings.ReplaceAll(op, "+", " ")
	op = strings.ReplaceAll(op, ",", " ")
	op = strings.TrimSpace(op)
	r, _ := regexp.Compile(" {2,}")
	op = string(r.ReplaceAll([]byte(op), []byte(" ")))
	return op
}

func DoOperation(nums []string) (int, error) {
	if len(nums) == 0 {
		return 0, nil
	}

	num := nums[0]
	n, err := strconv.Atoi(num)
	if err != nil {
		return 0, errors.New("invalid operation specified")
	}

	result, err := DoOperation(nums[1:])
	if err != nil {
		return 0, err
	}

	return n + result, nil
}

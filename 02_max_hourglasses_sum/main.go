package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY arr as parameter.
 */
func hourglassSum(arr [][]int32) int32 {
	var width, height = getMatrixlength(arr)
	var sum, maxSum int32
	maxSum = hourglassIndexSum(arr, 0, 0)
	indexWidthLimitForHourglass := (width / 2)
	indexHeightLimitForHourglass := (height / 2)
	for i := int32(0); i <= indexWidthLimitForHourglass; i += 1 {
		for j := int32(0); j <= indexHeightLimitForHourglass; j += 1 {
			sum = hourglassIndexSum(arr, i, j)
			if sum > maxSum {
				maxSum = sum
			}
		}
	}
	return maxSum
}

/*
 *
 * Given the array and a base index
 * Returns the sum of hourglass for that index base
 * If there is a not complete hourglass into the index, throws a panic
 */
func hourglassIndexSum(arr [][]int32, i, j int32) int32 {
	var iLimit, jLimit, sum int32
	iLimit = i + 3
	jLimit = j + 3
	jOriginal := j
	for ; i < iLimit; i += 1 {
		for j = jOriginal; j < jLimit; j += 1 {
			if (iLimit-i) == 2 && ((jLimit-j == 3) || (jLimit-j == 1)) {
				sum += 0
			} else {
				sum += arr[i][j]
			}
		}
	}
	return sum
}

func getMatrixlength(arr [][]int32) (int32, int32) {
	width := len(arr)
	height := len(arr[0])
	return int32(width), int32(height)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	var arr [][]int32
	for i := 0; i < 6; i++ {
		arrRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var arrRow []int32
		for _, arrRowItem := range arrRowTemp {
			arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
			checkError(err)
			arrItem := int32(arrItemTemp)
			arrRow = append(arrRow, arrItem)
		}

		if len(arrRow) != 6 {
			panic("Bad input")
		}

		arr = append(arr, arrRow)
	}

	result := hourglassSum(arr)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

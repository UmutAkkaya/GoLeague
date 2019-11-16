package api

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// Given an array of array flattens it by a level
func FlattenMatrix(records [][]string) []string {
	var response []string
	for _, row := range records {
		response = append(response, row...)
	}
	return response
}

// Applies the given operation to the elements of the matrix (records)
// Takes in response which is the initial value
func ApplyOperation(records [][]string, initialValue int, operation func(x int, y int) int) (int, error) {
	var result = initialValue
	for _, row := range records {
		for _, num := range row {
			parsed, err := strconv.Atoi(num)
			if err != nil {
				return 0, err
			}
			result = operation(result, parsed)
		}
	}
	return result, nil
}

// Transpose a given matrix
func InvertMatrix(records [][]string) [][]string {
	numRow := len(records) // 3

	if numRow == 0 {
		return [][]string{}
	}

	numColumn := len(records[0]) // 4
	invertedMatrix := make([][]string, numColumn)

	for i := 0; i < numColumn; i++ {
		invertedMatrix[i] = make([]string, numRow)
		for j := 0; j < numRow; j++ {
			invertedMatrix[i][j] = records[j][i]
		}
	}
	return invertedMatrix
}

func Add(x int, y int) int {
	return x + y
}

func Multiply(x int, y int) int {
	return x * y
}

// Given an array of array returns a comma seperated string representation
func MatrixToString(matrix [][]string) string {
	var response string
	for _, row := range matrix {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}

	return response
}

// Given a requests with a file returns the values in the file
func ReadRecords(request *http.Request) ([][]string, error) {
	file, _, err := request.FormFile("file")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, err
	}
	return records, nil
}
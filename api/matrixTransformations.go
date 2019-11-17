package api

import (
	"fmt"
	"strings"
)

/*
FlattenMatrix : Given an array of array flattens it by a level
*/
func FlattenMatrix(records [][]string) []string {
	var response []string
	for _, row := range records {
		response = append(response, row...)
	}
	return response
}

/*
InvertMatrix : Transpose a given matrix
*/
func InvertMatrix(records [][]string) [][]string {
	numRow := len(records)

	if numRow == 0 {
		return [][]string{}
	}

	numColumn := len(records[0])
	invertedMatrix := make([][]string, numColumn)

	for i := 0; i < numColumn; i++ {
		invertedMatrix[i] = make([]string, numRow)
		for j := 0; j < numRow; j++ {
			invertedMatrix[i][j] = records[j][i]
		}
	}
	return invertedMatrix
}

/*
MatrixToString : Given an array of array returns a comma seperated string representation
*/
func MatrixToString(matrix [][]string) string {
	var response string
	for _, row := range matrix {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}

	return response
}
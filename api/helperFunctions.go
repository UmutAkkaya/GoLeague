package api

import (
	"encoding/csv"
	"net/http"
	"strconv"
)

/*
ApplyOperation : Applies the given operation to the elements of the matrix (records)
Takes in response which is the initial value
*/
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

/*
Add : adds up given 2 numbers and returns them
 */
func Add(x int, y int) int {
	return x + y
}

/*
Multiply : multiplies given 2 numbers and returns them
*/
func Multiply(x int, y int) int {
	return x * y
}

/*
ReadRecords : Given a requests with a file returns the values in the file
 */
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
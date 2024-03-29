package api

import (
	"fmt"
	"net/http"
	"strings"
)

/*
EchoHandler : Handler method for 'echo' call. Prints out the given matrix
 */
func EchoHandler(w http.ResponseWriter, r *http.Request) {
	records, err := ReadRecords(r)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	var response = MatrixToString(records)
	fmt.Fprint(w, response)
}

/*
InvertHandler : Handler for the 'invert' call. Transposes the given matrix.
 */
func InvertHandler(w http.ResponseWriter, r *http.Request) {
	records, err := ReadRecords(r)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}

	inverted := InvertMatrix(records)

	var response = MatrixToString(inverted)
	fmt.Fprint(w, response)
}

/*
FlattenHandler : Handler for the 'flatten' call. Returns an array 1 level flattened down.
 */
func FlattenHandler(w http.ResponseWriter, r *http.Request) {
	records, err := ReadRecords(r)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}

	response := FlattenMatrix(records)

	fmt.Fprint(w, strings.Join(response, ","))
}

/*
SumHandler : Handler for the 'sum' call. Adds up the numbers in the matrix.
 */
func SumHandler(w http.ResponseWriter, r *http.Request) {
	records, err := ReadRecords(r)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}

	var response int
	response, err = ApplyOperation(records, response, Add)

	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}

	fmt.Fprint(w, response)
}

/*
MultiplyHandler : Handler for the 'multiply' call. Multiplies the numbers in the matrix.
 */
func MultiplyHandler(w http.ResponseWriter, r *http.Request) {
	records, err := ReadRecords(r)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}

	if len(records) == 0 {
		fmt.Fprint(w, 0)
		return
	}

	var response = 1
	response, err = ApplyOperation(records, response, Multiply)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}

	fmt.Fprint(w, response)
}


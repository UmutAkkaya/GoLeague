package api

import (
	"fmt"
	"net/http"
	"strings"
)

func EchoHandler(w http.ResponseWriter, r *http.Request) {
	records, err := ReadRecords(r)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	var response = MatrixToString(records)
	fmt.Fprint(w, response)
}

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

func FlattenHandler(w http.ResponseWriter, r *http.Request) {
	records, err := ReadRecords(r)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}

	response := FlattenMatrix(records)

	fmt.Fprint(w, strings.Join(response, ","))
}

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

func MultiplyHandler(w http.ResponseWriter, r *http.Request) {
	records, err := ReadRecords(r)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
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


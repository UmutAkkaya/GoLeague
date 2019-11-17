package apiTests

import (
	"../api"
	"net/http"
	"testing"
)

func TestSum(t *testing.T) {
	tables := []struct {
		num1 int
		num2 int
		expected int
	}{
		{1, 1, 2},
		{1, 2, 3},
		{2, 2, 4},
		{-5, 2, -3},
	}

	for _, table := range tables {
		total := api.Add(table.num1, table.num2)
		if total != table.expected {
			t.Errorf("Sum of (%d+%d) was incorrect, got: %d, want: %d.", table.num1, table.num2, total, table.expected)
		}
	}
}

func TestMultiply(t *testing.T) {
	tables := []struct {
		num1 int
		num2 int
		expected int
	}{
		{1, 1, 1},
		{0, 2, 0},
		{2, 2, 4},
		{-5, 2, -10},
	}

	for _, table := range tables {
		total := api.Multiply(table.num1, table.num2)
		if total != table.expected {
			t.Errorf("Multiplication of (%d*%d) was incorrect, got: %d, want: %d.", table.num1, table.num2, total, table.expected)
		}
	}
}


func TestApplyOperation(t *testing.T) {
	tables := []struct {
		input [][]string
		expected int
	}{
		{[][]string{{"1", "1", "1"}, {"3", "-4", "5"}, {"-2", "3", "4"}},
			13},

		{[][]string{{}}, 1},
	}

	for _, table := range tables {
		result, _ := api.ApplyOperation(table.input, 1, api.Add)
		if result != table.expected {
			t.Errorf("ApplyOperation failed, got: %v, want: %v.", result, table.expected)
		}
	}
}

func TestReadRecords(t *testing.T) {
	expected := [][]string{[]string{"a", "2", "c"}, []string{"4", "e", "6"}, []string{"g", "8", "i"}}

	var request http.Request
	request.FormFile("../inputs/testFile.csv")

	results, _ := api.ReadRecords(&request)
	if !arrayEquals(results, expected) {
		t.Errorf("ApplyOperation failed, got: %v, want: %v.", results, expected)
	}
}

func arrayEquals(arr1 [][]string, arr2 [][]string) bool {
	for i := range arr1 {
		for j := range arr1[i] {
			if arr1[i][j] != arr2[i][j] {
				return false
			}
		}
	}

	return true
}
package testing

import (
	"league/api"
	"net/http"
	"testing"
)

func TestSum(t *testing.T) {
	tables := []struct {
		x int
		y int
		n int
	}{
		{1, 1, 2},
		{1, 2, 3},
		{2, 2, 4},
		{-5, 2, -3},
	}

	for _, table := range tables {
		total := api.Add(table.x, table.y)
		if total != table.n {
			t.Errorf("Sum of (%d+%d) was incorrect, got: %d, want: %d.", table.x, table.y, total, table.n)
		}
	}
}

func TestMultiply(t *testing.T) {
	tables := []struct {
		x int
		y int
		n int
	}{
		{1, 1, 1},
		{0, 2, 0},
		{2, 2, 4},
		{-5, 2, -10},
	}

	for _, table := range tables {
		total := api.Multiply(table.x, table.y)
		if total != table.n {
			t.Errorf("Multiplication of (%d*%d) was incorrect, got: %d, want: %d.", table.x, table.y, total, table.n)
		}
	}
}

func TestMatrixToString(t *testing.T) {
	tables := []struct {
		x [][]string
		n string
	}{
		{[][]string{{"1", "1", "1"}, {"3", "-4", "5"}, {"-2", "3", "4"}},
			"1,1,1\n3,-4,5\n-2,3,4\n"},

		{[][]string{{}}, "\n"},

		{[][]string{{"a"}}, "a\n"},

		{[][]string{{"a", "5", "asda", "-4", "999", "-999"}}, "a,5,asda,-4,999,-999\n"},
	}

	for _, table := range tables {
		result := api.MatrixToString(table.x)
		if result != table.n {
			t.Errorf("MatrixToString failed, got: %v, want: %v.", result, table.n)
		}
	}
}

func TestApplyOperation(t *testing.T) {
	tables := []struct {
		x [][]string
		n int
	}{
		{[][]string{{"1", "1", "1"}, {"3", "-4", "5"}, {"-2", "3", "4"}},
			13},

		{[][]string{{}}, 1},
	}

	for _, table := range tables {
		result, _ := api.ApplyOperation(table.x, 1, api.Add)
		if result != table.n {
			t.Errorf("ApplyOperation failed, got: %v, want: %v.", result, table.n)
		}
	}
}

func TestFlatten(t *testing.T) {
	tables := []struct {
		x [][]string
		n []string
	}{
		{[][]string{{"1", "1", "1"}, {"3", "4", "5"}, {"2", "3", "4"}},
			[]string{"1", "1", "1", "3", "4", "5", "2", "3", "4"}},

		{[][]string{{"a", "-1", "999"}, {"0", "4", "5"}, {"2a", "3", "-99"}, {"4", "3", "99"}},
			[]string{"a", "-1", "999", "0", "4", "5", "2a", "3", "-99", "4", "3", "99"}},

		{[][]string{}, []string{}},
		{[][]string{{"a"}}, []string{"a"}},
	}

	for _, table := range tables {
		flattened := api.FlattenMatrix(table.x)
		if len(flattened) != len(table.n) {
			t.Errorf("Flattening was incorrect, got: %v, want: %v.", flattened, table.n)
		}

		for i, _ := range flattened {
			if flattened[i] != table.n[i] {
				t.Errorf("Flattening was incorrect, got: %v, want: %v.", flattened, table.n)
			}
		}
	}
}

func TestInvert(t *testing.T) {
	tables := []struct {
		x [][]string
		n [][]string
	}{
		{[][]string{{"1", "1", "1"}, {"-3", "4", "5"}, {"2s", "3", "a"}},
			[][]string{{"1", "-3", "2s"}, {"1", "4", "3"}, {"1", "5", "a"}}},

		{[][]string{{"a", "-1", "999"}, {"0", "4", "5"}, {"2a", "3", "-99"}, {"4", "3", "99"}},
			[][]string{{"a", "0", "2a", "4"}, {"-1", "4", "3", "3"}, {"999", "5", "-99", "99"}}},

		{[][]string{}, [][]string{}},
		{[][]string{{"a"}}, [][]string{{"a"}}},
	}

	for _, table := range tables {
		inverted := api.InvertMatrix(table.x)
		if len(inverted) != len(table.n) {
			t.Errorf("Inverted matrix doesnt match the length of expected, got: %v, want: %v.",
				len(inverted), len(table.n))
		}

		if !arrayEquals(inverted, table.n) {
			t.Errorf("Flattening was incorrect, got: %v, want: %v.", inverted, table.n)
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
	for i, _ := range arr1 {
		for j, _ := range arr1[i] {
			if arr1[i][j] != arr2[i][j] {
				return false
			}
		}
	}

	return true
}
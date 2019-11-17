package apiTests

import (
	"../api"
	"testing"
)

func TestMatrixToString(t *testing.T) {
	tables := []struct {
		input [][]string
		expected string
	}{
		{[][]string{{"1", "1", "1"}, {"3", "-4", "5"}, {"-2", "3", "4"}},
			"1,1,1\n3,-4,5\n-2,3,4\n"},

		{[][]string{{}}, "\n"},

		{[][]string{{"a"}}, "a\n"},

		{[][]string{{"a", "5", "asda", "-4", "999", "-999"}}, "a,5,asda,-4,999,-999\n"},
	}

	for _, table := range tables {
		result := api.MatrixToString(table.input)
		if result != table.expected {
			t.Errorf("MatrixToString failed, got: %v, want: %v.", result, table.expected)
		}
	}
}

func TestFlatten(t *testing.T) {
	tables := []struct {
		input [][]string
		expected []string
	}{
		{[][]string{{"1", "1", "1"}, {"3", "4", "5"}, {"2", "3", "4"}},
			[]string{"1", "1", "1", "3", "4", "5", "2", "3", "4"}},

		{[][]string{{"a", "-1", "999"}, {"0", "4", "5"}, {"2a", "3", "-99"}, {"4", "3", "99"}},
			[]string{"a", "-1", "999", "0", "4", "5", "2a", "3", "-99", "4", "3", "99"}},

		{[][]string{}, []string{}},
		{[][]string{{"a"}}, []string{"a"}},
	}

	for _, table := range tables {
		flattened := api.FlattenMatrix(table.input)
		if len(flattened) != len(table.expected) {
			t.Errorf("Flattening was incorrect, got: %v, want: %v.", flattened, table.expected)
		}

		for i := range flattened {
			if flattened[i] != table.expected[i] {
				t.Errorf("Flattening was incorrect, got: %v, want: %v.", flattened, table.expected)
			}
		}
	}
}

func TestInvert(t *testing.T) {
	tables := []struct {
		input [][]string
		expected [][]string
	}{
		{[][]string{{"1", "1", "1"}, {"-3", "4", "5"}, {"2s", "3", "a"}},
			[][]string{{"1", "-3", "2s"}, {"1", "4", "3"}, {"1", "5", "a"}}},

		{[][]string{{"a", "-1", "999"}, {"0", "4", "5"}, {"2a", "3", "-99"}, {"4", "3", "99"}},
			[][]string{{"a", "0", "2a", "4"}, {"-1", "4", "3", "3"}, {"999", "5", "-99", "99"}}},

		{[][]string{}, [][]string{}},
		{[][]string{{"a"}}, [][]string{{"a"}}},
	}

	for _, table := range tables {
		inverted := api.InvertMatrix(table.input)
		if len(inverted) != len(table.expected) {
			t.Errorf("Inverted matrix doesnt match the length of expected, got: %v, want: %v.",
				len(inverted), len(table.expected))
		}

		if !arrayEquals(inverted, table.expected) {
			t.Errorf("Flattening was incorrect, got: %v, want: %v.", inverted, table.expected)
		}
	}
}

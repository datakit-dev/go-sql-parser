package types

import "testing"

func Test_statementFrom(t *testing.T) {
	tests := []struct {
		input    string
		expected Statement
	}{
		{"select", SelectStatement},
		{"insert", InsertStatement},
		{"update", UpdateStatement},
		{"delete", DeleteStatement},
		{"create", CreateStatement},
		{"drop", DropStatement},
		{"alter", AlterStatement},
		{"foobar", UnknownStatement},
		{"", UnknownStatement},
	}

	for _, test := range tests {
		result := StatementFrom(test.input)
		if result != test.expected {
			t.Errorf("typeFrom(%s) returned %s, expected %s", test.input, result, test.expected)
		}
	}
}

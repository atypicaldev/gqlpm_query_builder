package builder

import (
	"testing"
)

// func createSpaceTab(n int) string {
// 	return strings.Repeat(" ", 5*n)
// }
// fmt.Sprintf("{\n%sset {\n%s<bob> <name> <Bob>\n%s}\n}", createSpaceTab(1), createSpaceTab(2), createSpaceTab(1)),

func TestMutationQueryBuilder(t *testing.T) {
	tests := []struct {
		inputs   [][]string
		expected string
		message  string
	}{
		{
			[][]string{{"bob", "name", "Bob"}},
			"{\n\tset {\n\n\t\t<bob> <name> <Bob>\n\n\t}\n}",
			"Expected :\n%s, \nGot:\n %v",
		},
		{
			[][]string{
				{"bob", "name", "Bob"},
				{"mary", "name", "Jill"},
			},
			"{\n\tset {\n\n\t\t<bob> <name> <Bob>\n\n\t\t<mary> <name> <Jill>\n\n\t}\n}",
			"Expected :\n%s, \nGot:\n %v",
		},
	}

	for _, test := range tests {
		b := NewMutationBuilder()

		for _, s := range test.inputs {
			b.AddTerm(s[0], s[1], s[2])
		}
		got := b.Build()

		if got != test.expected {
			t.Errorf("Expected:\n %s \nGot:\n %v", test.expected, got)
		}
	}
}

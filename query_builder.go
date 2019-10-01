package builder

import "text/template"

const queryTemplate = `
{
	query(func: {{.RootFunction}}){{.RootFilter}} {
		{{range .Fields}}
		{{expandField .}}
		{{end}}
	}
}
`

// QueryBuilder represents a struct that will allow the user to quickly build queries
type QueryBuilder struct {
	queryType    QueryType
	Fields       []Field
	RootFilter   string
	RootFunction string
}

// Field defines the outline for guts of a query
type Field struct {
	Alias, Name string
	Filters     []FilterFunction
	Expansion   []Field
}

var funcMap = template.FuncMap{
	"expandField": expandField,
}

func expandField(fields []Field) string {
	return "string"
}

// NewQueryBuilder returns an object to construct your graphql+- queries
func NewQueryBuilder() *QueryBuilder {
	return &QueryBuilder{queryType: Query}
}

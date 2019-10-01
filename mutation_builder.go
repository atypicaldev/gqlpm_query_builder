package builder

import (
	"fmt"
	"strings"
	"text/template"

	log "github.com/sirupsen/logrus"
)

const mutationTemplate = `
{
	set {
{{range .Terms}}
		<{{.Subject}}> <{{.Predicate}}> <{{.Value}}>
{{end}}
	}
}
`

// DgraphType is a predicate that is specific to the dgraph database
const DgraphType = "<dgraph.type>"

type term struct {
	subject   string
	predicate string
	value     string
}

// MutationBuilder is a struct used for building mutation queries
type MutationBuilder struct {
	queryType QueryType
	terms     []term
}

// NewMutationBuilder returns a new graphql+- mutation query builder
func NewMutationBuilder() *MutationBuilder {
	return &MutationBuilder{queryType: Set}
}

// AddTerm ads an RDF triple to the Proposed Query
func (b *MutationBuilder) AddTerm(subject, predicate, value string) *MutationBuilder {
	b.terms = append(b.terms, term{subject, predicate, value})
	return b
}

// Build returns the constructed RDF, graphql +- query
func (b *MutationBuilder) Build() string {
	return buildTemplate(b)
}

func buildTemplate(b *MutationBuilder) string {
	data := struct {
		Terms []struct {
			Subject, Predicate, Value string
		}
	}{
		collapseTerms(b.terms),
	}

	t, err := template.New(b.queryType.string()).Parse(b.queryType.getTemplate())
	if err != nil {
		log.Errorf("Error occurred while parsing template for builder of type %s", b.queryType.string())
		log.Fatalf("Error for parsing Template: %v", err)
	}

	var output strings.Builder

	err = t.Execute(&output, data)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}

	return strings.TrimSpace(output.String())
}

func collapseTerms(terms []term) []struct{ Subject, Predicate, Value string } {
	data := []struct{ Subject, Predicate, Value string }{}

	for _, t := range terms {
		data = append(data, struct{ Subject, Predicate, Value string }{Subject: t.subject, Predicate: t.predicate, Value: t.value})
	}

	return data
}

func baseFunc(s string) string {
	return fmt.Sprintf("{\n%s\n}", s)
}

func innerQuery(qt QueryType, terms []term) string {
	var query string
	switch qt {
	case Set:
		query = qt.string()
	case Delete:
	case Upsert:
	case Schema:
		log.Errorf("Unsupported Operation of type: %s", qt.string())
	default:
		log.Fatalf("Unnexpected QueryType provided: %v", qt)
	}
	return fmt.Sprintf("\t%s {\n%s\t}", query, concatTerms(terms))
}

func concatTerms(terms []term) string {
	result := ""

	for _, t := range terms {
		result += fmt.Sprintf("\t\t<%s> <%s> <%s>\n", t.subject, t.predicate, t.value)
	}

	return result
}

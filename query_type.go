package builder

import log "github.com/sirupsen/logrus"

// QueryType defines the type of graphQL+- query
type QueryType int

// Query defines the types of query
const (
	Query QueryType = 1 << iota
	Set
	Delete
	Upsert
	Schema
)

func (qt QueryType) string() string {
	switch qt {
	case Set:
		return "set"
	case Delete:
		return "delete"
	case Upsert:
		return "upsert"
	case Schema:
		return "schema"
	}
	return ""
}

func (qt QueryType) getTemplate() string {
	switch qt {
	case Set:
		return mutationTemplate
	case Delete:
	case Upsert:
	case Schema:
		log.Fatalf("No template implemented for %s", qt.string())
	}
	return ""
}

// FilterFunctionType defines values that describe the different filters
type FilterFunctionType int

// Defines the values used for FilterFunctionType
const (
	UID FilterFunctionType = 1 << iota
	UIDIn
	Has
	EQ
	LT
	GT
	LE
	GE
	AllOfTerms
	AnyOfTerms
	Match
	AllOfText
	AnyOfText
)

// FilterFunction provides a definition for the rdf filter functions to be constructed
type FilterFunction struct {
	Type FilterFunctionType
}

// FilterCombinator defines values that describe the operators to combine filters
type FilterCombinator int

// Defines the Values used for FilterCombinators
const (
	AND FilterCombinator = 1 << iota
	OR
	NOT
)

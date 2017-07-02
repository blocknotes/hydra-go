package main

import (
	"gopkg.in/mgo.v2/bson"
)

// Column - a column
type Column struct {
	Type        string `json:"type" bson:"type"`
	Validations string `json:"validations,omitempty" bson:"validations,omitempty"`
}

// Collection - a collection
type Collection struct {
	Name        string            `json:"name" bson:"name"`
	Singular    string            `json:"singular" bson:"singular"`
	Description string            `json:"description,omitempty" bson:"description,omitempty"`
	Columns     map[string]Column `json:"columns" bson:"columns"`
}

func (collection Collection) validate(data bson.M) []error {
	var results []error
	for k, v := range data {
		rules := collection.Columns[k].Validations
		ret := validate.Var(v, rules)
		if ret != nil {
			results = append(results, ret)
		}
	}

	return results
}

package main

// Project - a project
type Project struct {
	Name        string                `json:"name" bson:"name"`
	Endpoint    string                `json:"endpoint" bson:"endpoint"`
	Description string                `json:"description,omitempty" bson:"description,omitempty"`
	Collections map[string]Collection `json:"collections" bson:"collections"`
}

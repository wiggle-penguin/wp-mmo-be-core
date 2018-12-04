package model

type Item struct {
	ID         string            `json:"_id,omitempty" bson:"_id,omitempty"`
	Name       string            `json:"name,omitempty" bson:"name,omitempty"`
	Properties map[string]string `json:"properties,omitempty" bson:"properties,omitempty"`
}

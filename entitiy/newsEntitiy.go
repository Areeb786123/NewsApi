package entitiy

import "go.mongodb.org/mongo-driver/bson/primitive"

type NewsModels struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ImageUrl    string             `json:"imageUrl,omitempty"`
	AuthorName  string             `json:"authorName,omitempty"`
	Location    string             `json:"location,omitempty"`
	Title       string             `json:"title,omitempty"`
	Description string             `json:"description,omitempty"`
}

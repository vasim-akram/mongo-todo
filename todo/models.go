package todo

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	PENDING  = "pending"
	PROGRESS = "in_progress"
	DONE     = "done"
)

type Todo struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	Status      string             `json:"status" bson:"status"`
}

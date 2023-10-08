package model
import(
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Attendee struct {
    ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
    FirstName string            `json:"firstName" bson:"firstname"`
    LastName  string            `json:"lastName" bson:"lastname"`
    Year      string            `json:"year" bson:"year"`
    Nonveg    string            `json:"non-veg" bson:"nonveg"`
    Taken     string            `json:"taken" bson:"taken"`
    SerialNo  string            `json:"serialno" bson:"serialno"`
}


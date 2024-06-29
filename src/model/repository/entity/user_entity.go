package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserEntity struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
	Name     string             `bson:"name"`
	Age      int                `bson:"age"`
}

// func (ud *userDomain) GetJSONValue() (string, error) {
// 	b, err := json.Marshal(ud)
// 	if err != nil {
// 		return "", nil
// 	}
// 	return string(b), nil
// }

package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID    primitive.ObjectID `bson:"_id,emitonempty"`
	Name  string             `bson:"name"`
	Price int                `bson:"price"`
	Stock int                `bson:"stock"`
}

package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Compra struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	AlimentosAComprar []AlimentoCompra   `bson:"AlimentosCompra"`
	FechaCompra       time.Time          `bson:"FechaCompra"`
	CostoTotal        int                `bson:"CostoTotalCompra"`
	Usuario           string             `bson:"Usuario"`
}

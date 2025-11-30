package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Alimento struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty"`
	TipoAlimento       string             `bson:"Tipo"`
	MomentoAlimento    []string           `bson:"Momento"`
	NombreAlimento     string             `bson:"Nombre"`
	PrecioUnitario     int                `bson:"Precio"`
	CantidadActual     int                `bson:"CantidadActual"`
	CantidadMinima     int                `bson:"CantidadMinima"`
	FechaCreacion      time.Time          `bson:"FechaCreacion"`
	FechaActualizacion time.Time          `bson:"FechaActualizacion"`
	Usuario            string             `bson:"Usuario"`
}

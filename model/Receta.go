package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Receta struct {
	ID                  primitive.ObjectID `bson:"_id,omitempty"`
	MomentoReceta       string             `bson:"Momento"`
	NombreReceta        string             `bson:"Nombre"`
	AlimentosNecesarios []AlimentoReceta   `bson:"Alimentos"`
	FechaCreacion       time.Time          `bson:"FechaCreacion"`
	FechaActualizacion  time.Time          `bson:"FechaActualizacion"`
	Usuario             string             `bson:"Usuario"`
}

package model

type AlimentoReceta struct {
	ID                string `bson:"id"`
	NombreAlimento    string `bson:"NombreAlimento"`
	CantidadNecesaria int    `bson:"CantidadAlimento"`
}

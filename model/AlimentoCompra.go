package model

type AlimentoCompra struct {
	IDAlimento       string `bson:"IDAlimento"`
	NombreAlimento   string `bson:"NombreAlimento"`
	CantidadAComprar int    `bson:"CantidadAComprar"`
	Costo            int    `bson:"CostoUnitario"`
}

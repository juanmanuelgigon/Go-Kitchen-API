package repositories

import (
	"context"
	"fmt"

	"TPFINAL-GINCITO/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CompraRepositoryInterface interface {
	ObtenerCompras(usuario string) ([]model.Compra, error)
	InsertarCompra(compra model.Compra) (*mongo.InsertOneResult, error)
}

type CompraRepository struct {
	db DB
}

func NewCompraRepository(db DB) *CompraRepository {
	return &CompraRepository{
		db: db,
	}
}

func (repository CompraRepository) ObtenerCompras(usuario string) ([]model.Compra, error) {
	collection := repository.db.GetClient().Database("TPFINAL").Collection("Compras")
	filtro := bson.M{"Usuario": usuario}

	cursor, err := collection.Find(context.TODO(), filtro)

	defer cursor.Close(context.Background())

	var compras []model.Compra
	for cursor.Next(context.Background()) {
		var compra model.Compra
		err := cursor.Decode(&compra)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		compras = append(compras, compra)
	}

	return compras, err
}

func (repository CompraRepository) InsertarCompra(compra model.Compra) (*mongo.InsertOneResult, error) {
	collection := repository.db.GetClient().Database("TPFINAL").Collection("Compras")
	resultado, err := collection.InsertOne(context.TODO(), compra)
	return resultado, err
}

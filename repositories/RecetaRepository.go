package repositories

import (
	"context"
	"fmt"
	"time"

	"TPFINAL-GINCITO/model"
	"TPFINAL-GINCITO/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RecetaRepositoryInterface interface {
	ObtenerRecetas(usuario string) ([]model.Receta, error)
	ObtenerRecetaPorID(id string) (model.Receta, error)
	EliminarReceta(id primitive.ObjectID) (*mongo.DeleteResult, error)
	InsertarReceta(receta model.Receta) (*mongo.InsertOneResult, error)
	ModificarReceta(receta model.Receta) (*mongo.UpdateResult, error)
}

type RecetaRepository struct {
	db DB
}

func NewRecetaRepository(db DB) *RecetaRepository {
	return &RecetaRepository{
		db: db,
	}
}

func (repository RecetaRepository) ObtenerRecetas(usuario string) ([]model.Receta, error) {
	collection := repository.db.GetClient().Database("TPFINAL").Collection("Recetas")
	filtro := bson.M{"Usuario": usuario}
	cursor, err := collection.Find(context.TODO(), filtro)

	defer cursor.Close(context.Background())

	var recetas []model.Receta
	for cursor.Next(context.Background()) {
		var receta model.Receta
		err := cursor.Decode(&receta)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		recetas = append(recetas, receta)
	}

	return recetas, err
}

func (repository RecetaRepository) InsertarReceta(receta model.Receta) (*mongo.InsertOneResult, error) {
	collection := repository.db.GetClient().Database("TPFINAL").Collection("Recetas")
	resultado, err := collection.InsertOne(context.TODO(), receta)
	return resultado, err
}

func (repository RecetaRepository) ModificarReceta(receta model.Receta) (*mongo.UpdateResult, error) {
	collection := repository.db.GetClient().Database("TPFINAL").Collection("Recetas")

	filtro := bson.M{"_id": receta.ID}
	entidad := bson.M{
		"$set": bson.M{
			"Nombre":             receta.NombreReceta,
			"Momento":            receta.MomentoReceta,
			"FechaActualizacion": time.Now(),
			"Alimentos":          receta.AlimentosNecesarios,
		},
	}
	resultado, err := collection.UpdateOne(context.TODO(), filtro, entidad)

	return resultado, err
}

func (repository RecetaRepository) EliminarReceta(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	collection := repository.db.GetClient().Database("TPFINAL").Collection("Recetas")

	filtro := bson.M{"_id": id}

	resultado, err := collection.DeleteOne(context.TODO(), filtro)

	return resultado, err
}

func (repository RecetaRepository) ObtenerRecetaPorID(id string) (model.Receta, error) {

	collection := repository.db.GetClient().Database("TPFINAL").Collection("Recetas")
	objectID := utils.GetObjectIDFromStringID(id)
	filtro := bson.M{"_id": objectID}

	cursor, err := collection.Find(context.TODO(), filtro)

	defer cursor.Close(context.Background())

	var receta model.Receta

	for cursor.Next(context.Background()) {
		err := cursor.Decode(&receta)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
	return receta, err

}

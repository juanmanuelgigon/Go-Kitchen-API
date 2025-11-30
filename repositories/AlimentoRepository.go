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

type AlimentoRepositoryInterface interface {
	ObtenerAlimentos(usuario string) ([]model.Alimento, error)
	ObtenerAlimentoPorID(id string) (model.Alimento, error)
	EliminarAlimento(id primitive.ObjectID) (*mongo.DeleteResult, error)
	InsertarAlimento(alimento model.Alimento) (*mongo.InsertOneResult, error)
	ModificarAlimento(aula model.Alimento) (*mongo.UpdateResult, error)
}

type AlimentoRepository struct {
	db DB
}

func NewAlimentoRepository(db DB) *AlimentoRepository {
	return &AlimentoRepository{
		db: db,
	}
}

func (repository AlimentoRepository) ObtenerAlimentos(usuario string) ([]model.Alimento, error) {
	collection := repository.db.GetClient().Database("TPFINAL").Collection("Alimentos")
	filtro := bson.M{"Usuario": usuario}

	cursor, err := collection.Find(context.TODO(), filtro)

	defer cursor.Close(context.Background())

	var alimentos []model.Alimento
	for cursor.Next(context.Background()) {
		var alimento model.Alimento
		err := cursor.Decode(&alimento)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		alimentos = append(alimentos, alimento)
	}

	return alimentos, err
}

func (repository AlimentoRepository) InsertarAlimento(alimento model.Alimento) (*mongo.InsertOneResult, error) {
	collection := repository.db.GetClient().Database("TPFINAL").Collection("Alimentos")
	resultado, err := collection.InsertOne(context.TODO(), alimento)
	return resultado, err
}

func (repository AlimentoRepository) ModificarAlimento(alimento model.Alimento) (*mongo.UpdateResult, error) {
	collection := repository.db.GetClient().Database("TPFINAL").Collection("Alimentos")
	alimentoViejo, _ := repository.ObtenerAlimentoPorID(utils.GetStringIDFromObjectID(alimento.ID))
	if alimento.CantidadActual < 0 {
		alimento.CantidadActual = alimentoViejo.CantidadActual
	}
	if alimento.CantidadMinima < 0 {
		alimento.CantidadMinima = alimentoViejo.CantidadMinima
	}
	if alimento.PrecioUnitario < 0 {
		alimento.PrecioUnitario = alimentoViejo.PrecioUnitario
	}
	if len(alimento.MomentoAlimento) == 0 {
		alimento.MomentoAlimento = alimentoViejo.MomentoAlimento
	}
	if alimento.NombreAlimento == "" {
		alimento.NombreAlimento = alimentoViejo.NombreAlimento
	}
	if alimento.TipoAlimento == "" {
		alimento.TipoAlimento = alimentoViejo.TipoAlimento
	}

	filtro := bson.M{"_id": alimento.ID}
	entidad := bson.M{
		"$set": bson.M{
			"Nombre":             alimento.NombreAlimento,
			"Tipo":               alimento.TipoAlimento,
			"Momento":            alimento.MomentoAlimento,
			"Precio":             alimento.PrecioUnitario,
			"CantidadActual":     alimento.CantidadActual,
			"CantidadMinima":     alimento.CantidadMinima,
			"FechaActualizacion": time.Now(),
			"Usuario":            alimento.Usuario,
		},
	}
	resultado, err := collection.UpdateOne(context.TODO(), filtro, entidad)

	return resultado, err
}

func (repository AlimentoRepository) EliminarAlimento(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	collection := repository.db.GetClient().Database("TPFINAL").Collection("Alimentos")

	filtro := bson.M{"_id": id}

	resultado, err := collection.DeleteOne(context.TODO(), filtro)

	return resultado, err
}

func (repository AlimentoRepository) ObtenerAlimentoPorID(id string) (model.Alimento, error) {

	collection := repository.db.GetClient().Database("TPFINAL").Collection("Alimentos")
	objectID := utils.GetObjectIDFromStringID(id)
	filtro := bson.M{"_id": objectID}

	cursor, err := collection.Find(context.TODO(), filtro)

	defer cursor.Close(context.Background())
	var alimento model.Alimento

	for cursor.Next(context.Background()) {
		err := cursor.Decode(&alimento)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
	return alimento, err
}

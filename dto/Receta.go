package dto

import (
	"TPFINAL-GINCITO/model"
	"TPFINAL-GINCITO/utils"
	//"time"
)

type Receta struct {
	ID                  string
	MomentoReceta       string
	NombreReceta        string
	AlimentosNecesarios []AlimentoReceta
}

func NewReceta(receta model.Receta) *Receta {
	var listado []AlimentoReceta = NewListAlimentoReceta(receta.AlimentosNecesarios)
	return &Receta{
		ID:                  utils.GetStringIDFromObjectID(receta.ID),
		MomentoReceta:       receta.MomentoReceta,
		NombreReceta:        receta.NombreReceta,
		AlimentosNecesarios: listado,
	}
}

func (receta Receta) GetModel() model.Receta {
	var listado []model.AlimentoReceta = GetModel(receta.AlimentosNecesarios)
	return model.Receta{
		ID:                  utils.GetObjectIDFromStringID(receta.ID),
		MomentoReceta:       receta.MomentoReceta,
		NombreReceta:        receta.NombreReceta,
		AlimentosNecesarios: listado,
	}
}

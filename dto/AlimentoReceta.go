package dto

import (
	"TPFINAL-GINCITO/model"
)

type AlimentoReceta struct {
	IDAlimento        string
	NombreAlimento    string
	CantidadNecesaria int
}

func NewListAlimentoReceta(alimentoReceta []model.AlimentoReceta) []AlimentoReceta {
	var listado []AlimentoReceta
	i := 0
	for i < len(alimentoReceta) {
		AlimentoReceta := AlimentoReceta{
			IDAlimento:        alimentoReceta[i].ID,
			NombreAlimento:    alimentoReceta[i].NombreAlimento,
			CantidadNecesaria: alimentoReceta[i].CantidadNecesaria,
		}
		listado = append(listado, AlimentoReceta)
		i++
	}
	return listado
}

func GetModel(alimentoReceta []AlimentoReceta) []model.AlimentoReceta {
	var listado []model.AlimentoReceta
	i := 0
	for i < len(alimentoReceta) {
		AlimentoReceta := model.AlimentoReceta{
			ID:                alimentoReceta[i].IDAlimento,
			NombreAlimento:    alimentoReceta[i].NombreAlimento,
			CantidadNecesaria: alimentoReceta[i].CantidadNecesaria,
		}
		listado = append(listado, AlimentoReceta)
		i++
	}
	return listado
}

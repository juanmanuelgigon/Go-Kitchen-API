package dto

import (
	"TPFINAL-GINCITO/model"
	"TPFINAL-GINCITO/utils"
	"time"
)

type Compra struct {
	ID                string
	AlimentosAComprar []AlimentoCompra
	CostoTotal        int
}

func NewCompra(compra model.Compra) *Compra {
	var listado []AlimentoCompra = NewListAlimentoCompra(compra.AlimentosAComprar)
	return &Compra{
		ID:                utils.GetStringIDFromObjectID(compra.ID),
		AlimentosAComprar: listado,
		CostoTotal:        compra.CostoTotal,
	}
}

func (compra Compra) GetModel() model.Compra {
	var listado []model.AlimentoCompra = GetModelAlimentoCompra(compra.AlimentosAComprar)
	return model.Compra{
		ID:                utils.GetObjectIDFromStringID(compra.ID),
		AlimentosAComprar: listado,
		CostoTotal:        compra.CostoTotal,
		FechaCompra:       time.Now(),
	}
}

package dto

import (
	"TPFINAL-GINCITO/model"
	"TPFINAL-GINCITO/utils"
	"time"
)

type Alimento struct {
	IDAlimento      string
	TipoAlimento    string
	MomentoAlimento []string
	NombreAlimento  string
	PrecioUnitario  int
	CantidadActual  int
	CantidadMinima  int
}

func NewAlimento(alimento model.Alimento) *Alimento {
	return &Alimento{
		TipoAlimento:    alimento.TipoAlimento,
		MomentoAlimento: alimento.MomentoAlimento,
		NombreAlimento:  alimento.NombreAlimento,
		PrecioUnitario:  alimento.PrecioUnitario,
		IDAlimento:      utils.GetStringIDFromObjectID(alimento.ID),
		CantidadActual:  alimento.CantidadActual,
		CantidadMinima:  alimento.CantidadMinima,
	}
}

func (alimento Alimento) GetModel() model.Alimento {
	return model.Alimento{
		ID:              utils.GetObjectIDFromStringID(alimento.IDAlimento),
		TipoAlimento:    alimento.TipoAlimento,
		MomentoAlimento: alimento.MomentoAlimento,
		NombreAlimento:  alimento.NombreAlimento,
		PrecioUnitario:  alimento.PrecioUnitario,
		CantidadActual:  alimento.CantidadActual,
		CantidadMinima:  alimento.CantidadMinima,
		FechaCreacion:   time.Now(),
	}

}

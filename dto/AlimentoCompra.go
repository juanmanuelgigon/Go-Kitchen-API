package dto

import (
	"TPFINAL-GINCITO/model"
)

type AlimentoCompra struct {
	IDAlimento       string
	NombreAlimento   string
	CantidadAComprar int
	Costo            int
}

func NewListAlimentoCompra(alimentoCompra []model.AlimentoCompra) []AlimentoCompra {
	var listado []AlimentoCompra
	i := 0
	for i < len(alimentoCompra) {
		AlimentoCompra := AlimentoCompra{
			IDAlimento:       alimentoCompra[i].IDAlimento,
			NombreAlimento:   alimentoCompra[i].NombreAlimento,
			CantidadAComprar: alimentoCompra[i].CantidadAComprar,
			Costo:            alimentoCompra[i].Costo,
		}
		listado = append(listado, AlimentoCompra)
		i++
	}
	return listado
}

func GetModelAlimentoCompra(alimentoCompra []AlimentoCompra) []model.AlimentoCompra {
	var listado []model.AlimentoCompra
	i := 0
	for i < len(alimentoCompra) {
		AlimentoCompra := model.AlimentoCompra{
			IDAlimento:       alimentoCompra[i].IDAlimento,
			NombreAlimento:   alimentoCompra[i].NombreAlimento,
			CantidadAComprar: alimentoCompra[i].CantidadAComprar,
			Costo:            alimentoCompra[i].Costo,
		}
		listado = append(listado, AlimentoCompra)
		i++
	}
	return listado
}

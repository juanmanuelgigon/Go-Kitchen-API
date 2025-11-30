package services

import (
	"TPFINAL-GINCITO/dto"
	"TPFINAL-GINCITO/model"
	"TPFINAL-GINCITO/repositories"
	"TPFINAL-GINCITO/utils"
	"sort"
	"strings"
	"time"
)

type CompraInterface interface {
	ObtenerAlimentosCompra(filtroNombre, filtroTipo, usuario string) (*dto.Compra, *utils.CustomError)
	GenerarCompra(productos []string, usuario string) (*dto.Compra, *utils.CustomError)
	ObtenerDatosCompras(usuario string) ([]dto.ElementoGrafico, *utils.CustomError)
}

type CompraService struct {
	alimentoRepository repositories.AlimentoRepositoryInterface
	compraRepository   repositories.CompraRepositoryInterface
}

func NewCompraService(compraRepositorio repositories.CompraRepositoryInterface, alimentoRepositorio repositories.AlimentoRepositoryInterface) *CompraService {
	return &CompraService{
		alimentoRepository: alimentoRepositorio,
		compraRepository:   compraRepositorio,
	}
}
func (service *CompraService) ObtenerAlimentosCompra(filtroNombreAlimento, filtroTipoAlimento, usuario string) (*dto.Compra, *utils.CustomError) {
	alimentos, err := service.alimentoRepository.ObtenerAlimentos(usuario)
	if err != nil {
		return nil, utils.NewCustomError("ERR_015", "Error al obtener alimentos de la base de datos.")
	}

	var alimentosParaComprar []dto.AlimentoCompra
	for _, alimento := range alimentos {
		if alimento.CantidadActual < alimento.CantidadMinima {
			alimentoParaComprar := newAlimentoCompra(alimento)
			alimentosParaComprar = append(alimentosParaComprar, alimentoParaComprar)
		}
	}

	var alimentosCompra []dto.AlimentoCompra

	if len(alimentosParaComprar) != 0 {
		for _, alimento := range alimentosParaComprar {
			alim, err := service.alimentoRepository.ObtenerAlimentoPorID(alimento.IDAlimento)
			if err != nil {
				return nil, utils.NewCustomError("ERR_016", "Error al obtener el alimento con ID: "+alimento.IDAlimento)
			}
			if alim.Usuario != usuario {
				return nil, utils.NewCustomError("ERR_028", "Alimento no disponible para el usuario")
			}
			coincideNombre := filtroNombreAlimento == ""
			coincideTipo := filtroTipoAlimento == ""

			if filtroTipoAlimento != "" && strings.Contains(strings.ToLower(alim.TipoAlimento), strings.ToLower(filtroTipoAlimento)) {
				coincideTipo = true
			}
			if filtroNombreAlimento != "" && strings.Contains(strings.ToLower(alim.NombreAlimento), strings.ToLower(filtroNombreAlimento)) {
				coincideNombre = true
			}

			if coincideNombre && coincideTipo {
				alimentosCompra = append(alimentosCompra, alimento)
			}
		}
	}

	if len(alimentosCompra) == 0 {
		return &dto.Compra{}, nil
	}

	costoTotal := 0
	for _, alimentosCompra := range alimentosCompra {
		costoTotal += alimentosCompra.Costo
	}

	laCompra := dto.Compra{
		CostoTotal:        costoTotal,
		AlimentosAComprar: alimentosCompra,
	}

	return &laCompra, nil
}

func newAlimentoCompra(alimento model.Alimento) dto.AlimentoCompra {
	return dto.AlimentoCompra{
		IDAlimento:       utils.GetStringIDFromObjectID(alimento.ID),
		NombreAlimento:   alimento.NombreAlimento,
		CantidadAComprar: alimento.CantidadMinima - alimento.CantidadActual,
		Costo:            alimento.PrecioUnitario * (alimento.CantidadMinima - alimento.CantidadActual),
	}
}

func (service *CompraService) GenerarCompra(productosSeleccionados []string, usuario string) (*dto.Compra, *utils.CustomError) {
	compra, error := service.ObtenerAlimentosCompra("", "", usuario)
	if error != nil {
		return nil, error
	}

	if len(productosSeleccionados) != 0 {
		costoTotal := 0
		var listadoCompraPersonalizado []dto.AlimentoCompra
		for _, producto := range productosSeleccionados {
			producto = strings.TrimSpace(producto)
			for _, alimentoCompra := range compra.AlimentosAComprar {
				if producto == alimentoCompra.IDAlimento {
					listadoCompraPersonalizado = append(listadoCompraPersonalizado, alimentoCompra)
					costoTotal += alimentoCompra.Costo
				}
			}
		}
		compra.AlimentosAComprar = listadoCompraPersonalizado
		compra.CostoTotal = costoTotal
	}

	var compraInsertar model.Compra
	compraInsertar = compra.GetModel()
	compraInsertar.Usuario = usuario
	_, err := service.compraRepository.InsertarCompra(compraInsertar)

	if err != nil {
		return nil, utils.NewCustomError("ERR_010", "Error al insertar la compra en la base de datos")
	}

	for _, alimentoCompra := range compra.AlimentosAComprar {
		var alimentoM model.Alimento
		alimentoM, _ = service.alimentoRepository.ObtenerAlimentoPorID(alimentoCompra.IDAlimento)
		alimento := dto.Alimento{
			IDAlimento:     alimentoCompra.IDAlimento,
			CantidadActual: (alimentoCompra.CantidadAComprar + alimentoM.CantidadActual),
			PrecioUnitario: alimentoM.PrecioUnitario,
			CantidadMinima: alimentoM.CantidadMinima,
		}
		alimentoModel := alimento.GetModel()
		alimentoModel.Usuario = usuario
		_, err := service.alimentoRepository.ModificarAlimento(alimentoModel)
		if err != nil {
			return nil, utils.NewCustomError("ERR_019", "Error al modificar el alimento en la base de datos")
		}
	}
	return compra, nil
}

func (service *CompraService) ObtenerDatosCompras(usuario string) ([]dto.ElementoGrafico, *utils.CustomError) {
	compras, err := service.compraRepository.ObtenerCompras(usuario)
	if err != nil {
		return nil, utils.NewCustomError("ERR_029", "Error al obtener las compras de la base de datos.")
	}

	var resultado []dto.ElementoGrafico
	for _, compra := range compras {
		if compra.FechaCompra.Year() == time.Now().Year() {
			bandera := false
			for i := range resultado {
				if resultado[i].Tipo == compra.FechaCompra.Month().String() {
					resultado[i].Cantidad += compra.CostoTotal
					bandera = true
					break
				}
			}
			if !bandera {
				nuevo := dto.ElementoGrafico{
					Cantidad: compra.CostoTotal,
					Tipo:     compra.FechaCompra.Month().String(),
				}
				resultado = append(resultado, nuevo)
			}
		}
	}
	sort.Slice(resultado, func(i, j int) bool {
		monthOrder := map[string]int{
			"January":   1,
			"February":  2,
			"March":     3,
			"April":     4,
			"May":       5,
			"June":      6,
			"July":      7,
			"August":    8,
			"September": 9,
			"October":   10,
			"November":  11,
			"December":  12,
		}

		return monthOrder[resultado[i].Tipo] < monthOrder[resultado[j].Tipo]
	})
	return resultado, nil
}

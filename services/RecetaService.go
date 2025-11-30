package services

import (
	"TPFINAL-GINCITO/dto"
	"TPFINAL-GINCITO/model"
	"TPFINAL-GINCITO/repositories"
	"TPFINAL-GINCITO/utils"
	"sort"
	"strings"
)

type RecetaInterface interface {
	ObtenerRecetas(filtroNombreProducto, filtroTipoProducto, filtroMomento, usuario string) ([]*dto.Receta, *utils.CustomError)
	InsertarReceta(receta *dto.Receta, usuario string) (*dto.Receta, *utils.CustomError)
	ModificarReceta(receta *dto.Receta, usuario string) (*dto.Receta, *utils.CustomError)
	EliminarReceta(id string) (*dto.Receta, *utils.CustomError)
	ObtenerRecetaPorID(id, usuario string) (*dto.Receta, *utils.CustomError)
	ObtenerRecetasPorMomento(usuario string) ([]dto.ElementoGrafico, *utils.CustomError)
	ObtenerRecetasPorTipoAlimento(usuario string) ([]dto.ElementoGrafico, *utils.CustomError)
}
type RecetaService struct {
	recetaRepository   repositories.RecetaRepositoryInterface
	alimentoRepository repositories.AlimentoRepositoryInterface
}

func NewRecetaService(recetaRepo repositories.RecetaRepositoryInterface, alimentoRepo repositories.AlimentoRepositoryInterface) *RecetaService {
	return &RecetaService{
		recetaRepository:   recetaRepo,
		alimentoRepository: alimentoRepo,
	}
}

func (service *RecetaService) ObtenerRecetasPorTipoAlimento(usuario string) ([]dto.ElementoGrafico, *utils.CustomError) {
	recetasDB, err := service.recetaRepository.ObtenerRecetas(usuario)
	if err != nil {
		return nil, utils.NewCustomError("ERR_014", "Error al obtener las recetas de la base de datos")
	}
	var resultado []dto.ElementoGrafico
	for _, receta := range recetasDB {
		for _, alimento := range receta.AlimentosNecesarios {
			alim, _ := service.alimentoRepository.ObtenerAlimentoPorID(alimento.ID)
			var bandera bool = false
			for i := range resultado {
				if alim.TipoAlimento == resultado[i].Tipo {
					resultado[i].Cantidad = resultado[i].Cantidad + 1
					bandera = true
					break
				}
			}
			if !bandera {
				nuevo := dto.ElementoGrafico{
					Cantidad: 1,
					Tipo:     alim.TipoAlimento,
				}
				resultado = append(resultado, nuevo)
			}
		}
	}
	return resultado, nil
}

func (service *RecetaService) ObtenerRecetasPorMomento(usuario string) ([]dto.ElementoGrafico, *utils.CustomError) {
	recetasDB, err := service.recetaRepository.ObtenerRecetas(usuario)
	if err != nil {
		return nil, utils.NewCustomError("ERR_014", "Error al obtener las recetas de la base de datos")
	}
	var resultado []dto.ElementoGrafico
	for _, receta := range recetasDB {
		var bandera bool = false
		for i := range resultado {
			if receta.MomentoReceta == resultado[i].Tipo {
				resultado[i].Cantidad = resultado[i].Cantidad + 1
				bandera = true
				break
			}
		}
		if !bandera {
			nuevo := dto.ElementoGrafico{
				Cantidad: 1,
				Tipo:     receta.MomentoReceta,
			}
			resultado = append(resultado, nuevo)
		}
	}
	sort.Slice(resultado, func(i, j int) bool {
		momentoOrder := map[string]int{
			"Desayuno": 1,
			"Almuerzo": 2,
			"Merienda": 3,
			"Cena":     4,
		}

		return momentoOrder[resultado[i].Tipo] < momentoOrder[resultado[j].Tipo]
	})
	return resultado, nil
}

func (service *RecetaService) InsertarReceta(receta *dto.Receta, usuario string) (*dto.Receta, *utils.CustomError) {
	var alimentosActualizar []model.Alimento

	for _, alimento := range receta.AlimentosNecesarios {
		alim, err := service.alimentoRepository.ObtenerAlimentoPorID(alimento.IDAlimento)
		if err != nil {
			return nil, utils.NewCustomError("ERR_006", "Error al obtener el alimento con ID: "+alimento.IDAlimento)
		}

		if alim.CantidadActual < alimento.CantidadNecesaria {
			return nil, utils.NewCustomError("ERR_007", "Cantidad insuficiente del alimento con ID: "+alimento.IDAlimento)
		}

		esMomento := false
		for _, momento := range alim.MomentoAlimento {
			if momento == receta.MomentoReceta {
				esMomento = true
				break
			}
		}

		if !esMomento {
			return nil, utils.NewCustomError("ERR_008", "El alimento con ID: "+alimento.IDAlimento+" no es válido para el momento: "+receta.MomentoReceta)
		}
		alimentosActualizar = append(alimentosActualizar, alim)
	}

	for i, alimento := range receta.AlimentosNecesarios {
		alimentosActualizar[i].CantidadActual -= alimento.CantidadNecesaria
		_, err := service.alimentoRepository.ModificarAlimento(alimentosActualizar[i])
		if err != nil {
			return nil, utils.NewCustomError("ERR_009", "Error al actualizar la cantidad del alimento con ID: "+alimento.IDAlimento)
		}
	}
	recetaInsertar := receta.GetModel()
	recetaInsertar.Usuario = usuario
	_, err := service.recetaRepository.InsertarReceta(recetaInsertar)
	if err != nil {
		return nil, utils.NewCustomError("ERR_010", "Error al insertar la receta en la base de datos")
	}
	return receta, nil
}

func (service *RecetaService) ModificarReceta(receta *dto.Receta, usuario string) (*dto.Receta, *utils.CustomError) {
	for _, alim := range receta.AlimentosNecesarios {
		alimento, _ := service.alimentoRepository.ObtenerAlimentoPorID(alim.IDAlimento)
		var bandera bool = false
		for _, momentox := range alimento.MomentoAlimento {
			if momentox == receta.MomentoReceta {
				bandera = true
			}
		}
		if !bandera {
			return nil, utils.NewCustomError("ERR_025", "Error, no todos los alimentos son aptos para este momento")
		}
	}
	recetaModificar := receta.GetModel()
	recetaModificar.Usuario = usuario
	resultado, err := service.recetaRepository.ModificarReceta(recetaModificar)
	if err != nil {
		return nil, utils.NewCustomError("ERR_011", "Error al modificar la receta en la base de datos")
	}

	if resultado.MatchedCount == 0 {
		return nil, utils.NewCustomError("ERR_012", "No se encontró ninguna receta con ese ID para modificar")
	} else if resultado.ModifiedCount == 0 {
		return nil, utils.NewCustomError("ERR_013", "No se realizaron modificaciones en la receta")
	}
	return receta, nil
}

func (service *RecetaService) EliminarReceta(id string) (*dto.Receta, *utils.CustomError) {
	receta, err := service.recetaRepository.ObtenerRecetaPorID(id)
	if err != nil {
		return nil, utils.NewCustomError("ERR_001", "Error al obtener la receta. Receta no encontrada")
	}

	resultado, err := service.recetaRepository.EliminarReceta(utils.GetObjectIDFromStringID(id))
	if err != nil {
		return nil, utils.NewCustomError("ERR_002", "Error al eliminar la receta en la base de datos")
	}

	if resultado.DeletedCount == 0 {
		return nil, utils.NewCustomError("ERR_003", "No se eliminó ninguna receta. Posible ID inválido")
	}

	for _, alimento := range receta.AlimentosNecesarios {
		alimentoNecesario, err := service.alimentoRepository.ObtenerAlimentoPorID(alimento.ID)
		if err != nil {
			return nil, utils.NewCustomError("ERR_004", "Error al obtener el alimento con ID"+alimento.ID)
		}

		alimentoNecesario.CantidadActual += alimento.CantidadNecesaria
		_, err = service.alimentoRepository.ModificarAlimento(alimentoNecesario)
		if err != nil {
			return nil, utils.NewCustomError("ERR_005", "Error al modificar la cantidad del alimento con ID"+alimento.ID)
		}
	}
	receta2 := dto.NewReceta(receta)
	return receta2, nil
}

func (service *RecetaService) ObtenerRecetaPorID(id string, usuario string) (*dto.Receta, *utils.CustomError) {
	receta, err := service.recetaRepository.ObtenerRecetaPorID(id)
	if err != nil {
		return nil, utils.NewCustomError("ERR_017", "Error al obtener los alimentos de la base de datos")
	}
	if receta.Usuario != usuario {
		return nil, utils.NewCustomError("ERR_029", "Receta no disponible para el usuario")
	}
	recetaReturn := dto.NewReceta(receta)
	return recetaReturn, nil
}

func (service *RecetaService) ObtenerRecetas(filtroMomento, filtroNombreProducto, filtroTipoProducto, usuario string) ([]*dto.Receta, *utils.CustomError) {
	recetasDB, err := service.recetaRepository.ObtenerRecetas(usuario)
	if err != nil {
		return nil, utils.NewCustomError("ERR_014", "Error al obtener las recetas de la base de datos")
	}

	var recetas []*dto.Receta
	for _, recetaDB := range recetasDB {
		receta := dto.NewReceta(recetaDB)
		todosLosAlimentosSuficientes := true

		for _, alimento := range receta.AlimentosNecesarios {
			alim, err := service.alimentoRepository.ObtenerAlimentoPorID(alimento.IDAlimento)
			if err != nil {
				return nil, utils.NewCustomError("ERR_015", "Error al obtener el alimento con ID: "+alimento.IDAlimento)
			}
			if alim.CantidadActual < alimento.CantidadNecesaria {
				todosLosAlimentosSuficientes = false
				break
			}
			if alim.Usuario != usuario {
				todosLosAlimentosSuficientes = false
				break
			}
		}

		if todosLosAlimentosSuficientes {
			coincideMomento := filtroMomento == "" || strings.Contains(strings.ToLower(receta.MomentoReceta), strings.ToLower(filtroMomento))
			coincideNombre := filtroNombreProducto == ""
			coincideTipo := filtroTipoProducto == ""

			for _, alimento := range receta.AlimentosNecesarios {
				alim, err := service.alimentoRepository.ObtenerAlimentoPorID(alimento.IDAlimento)
				if err != nil {
					return nil, utils.NewCustomError("ERR_016", "Error al obtener el alimento con ID: "+alimento.IDAlimento)
				}
				if filtroTipoProducto != "" && strings.Contains(strings.ToLower(alim.TipoAlimento), strings.ToLower(filtroTipoProducto)) {
					coincideTipo = true
				}
				if filtroNombreProducto != "" && strings.Contains(strings.ToLower(alim.NombreAlimento), strings.ToLower(filtroNombreProducto)) {
					coincideNombre = true
				}
			}

			if coincideMomento && coincideNombre && coincideTipo {
				recetas = append(recetas, receta)
			}
		}
	}

	return recetas, nil
}

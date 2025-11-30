package services

import (
	"TPFINAL-GINCITO/dto"
	"TPFINAL-GINCITO/repositories"
	"TPFINAL-GINCITO/utils"
	"strings"
)

type AlimentoInterface interface {
	ObtenerAlimentos(filtroNombre string, usuario string) ([]*dto.Alimento, *utils.CustomError)
	InsertarAlimento(alimento *dto.Alimento, usuario string) (*dto.Alimento, *utils.CustomError)
	ModificarAlimento(alimento *dto.Alimento, usuario string) (*dto.Alimento, *utils.CustomError)
	EliminarAlimento(id string) (*dto.Alimento, *utils.CustomError)
	ObtenerAlimentoPorID(id, usuario string) (*dto.Alimento, *utils.CustomError)
}
type AlimentoService struct {
	alimentoRepository repositories.AlimentoRepositoryInterface
}

func NewAlimentoService(alimentoRepository repositories.AlimentoRepositoryInterface) *AlimentoService {
	return &AlimentoService{
		alimentoRepository: alimentoRepository,
	}
}
func (service *AlimentoService) ObtenerAlimentoPorID(id, usuario string) (*dto.Alimento, *utils.CustomError) {
	alim, err := service.alimentoRepository.ObtenerAlimentoPorID(id)
	if err != nil {
		return nil, utils.NewCustomError("ERR_017", "Error al obtener los alimentos de la base de datos")
	}
	if alim.Usuario != usuario {
		return nil, utils.NewCustomError("ERR_028", "Alimento no disponible para el usuario")
	}
	alimReturn := dto.NewAlimento(alim)
	return alimReturn, nil
}

func (service *AlimentoService) ObtenerAlimentos(filtroNombre string, usuario string) ([]*dto.Alimento, *utils.CustomError) {
	alimentosDB, err := service.alimentoRepository.ObtenerAlimentos(usuario)
	if err != nil {
		return nil, utils.NewCustomError("ERR_017", "Error al obtener los alimentos de la base de datos")
	}

	var alimentos []*dto.Alimento
	for _, alimentoDB := range alimentosDB {
		alimento := dto.NewAlimento(alimentoDB)
		if filtroNombre == "" || strings.Contains(alimento.NombreAlimento, filtroNombre) {
			alimentos = append(alimentos, alimento)
		}
	}

	return alimentos, nil
}

func (service *AlimentoService) InsertarAlimento(alimento *dto.Alimento, usuario string) (*dto.Alimento, *utils.CustomError) {
	if alimento.CantidadActual > 0 && alimento.CantidadMinima > 0 && len(alimento.MomentoAlimento) > 0 && alimento.NombreAlimento != "" && alimento.PrecioUnitario > 0 && alimento.TipoAlimento != "" {
		alimentoInsertar := alimento.GetModel()
		alimentoInsertar.Usuario = usuario
		_, err := service.alimentoRepository.InsertarAlimento(alimentoInsertar)
		if err != nil {
			return nil, utils.NewCustomError("ERR_018", "Error al insertar el alimento en la base de datos")
		} else {
			return alimento, nil
		}
	}
	return nil, utils.NewCustomError("ERR_024", "Alimento imposible de insertar")
}

func (service *AlimentoService) ModificarAlimento(alimento *dto.Alimento, usuario string) (*dto.Alimento, *utils.CustomError) {
	alimentoInsertar := alimento.GetModel()
	alimentoInsertar.Usuario = usuario
	resultado, err := service.alimentoRepository.ModificarAlimento(alimentoInsertar)
	if err != nil {
		return nil, utils.NewCustomError("ERR_019", "Error al modificar el alimento en la base de datos")
	}

	if resultado.MatchedCount == 0 {
		return nil, utils.NewCustomError("ERR_020", "No se encontró el alimento para modificar")
	} else if resultado.ModifiedCount == 0 {
		return nil, utils.NewCustomError("ERR_021", "No se realizaron cambios en el alimento")
	}

	return alimento, nil
}

func (service *AlimentoService) EliminarAlimento(id string) (*dto.Alimento, *utils.CustomError) {
	resultado, err := service.alimentoRepository.EliminarAlimento(utils.GetObjectIDFromStringID(id))
	if err != nil {
		return nil, utils.NewCustomError("ERR_022", "Error al eliminar el alimento en la base de datos")
	}

	if resultado.DeletedCount == 0 {
		return nil, utils.NewCustomError("ERR_023", "No se eliminó ningún alimento. Posible ID inválido")
	}
	alim, _ := service.alimentoRepository.ObtenerAlimentoPorID(id)
	alimentoReturn := dto.NewAlimento(alim)

	return alimentoReturn, nil
}

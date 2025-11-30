package handlers

import (
	"TPFINAL-GINCITO/dto"
	"TPFINAL-GINCITO/services"
	"TPFINAL-GINCITO/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AlimentoHandler struct {
	alimentoService services.AlimentoInterface
}

func NewAlimentoHandler(alimentoService services.AlimentoInterface) *AlimentoHandler {
	return &AlimentoHandler{
		alimentoService: alimentoService,
	}
}

func (handler *AlimentoHandler) ObtenerAlimentos(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))
	filtroNombre := c.Query("tipo")
	log.Printf("[handler:AlimentoHandler] Iniciando obtención de alimentos para el usuario: %s, Filtro: %s", user.Codigo, filtroNombre)
	alimentos, err := handler.alimentoService.ObtenerAlimentos(filtroNombre, user.Codigo)
	if alimentos == nil {
		alimentos = []*dto.Alimento{}
	}
	if err != nil {
		log.Printf("[handler:AlimentoHandler] Error al obtener alimentos para el usuario %s: %s", user.Codigo, err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Code,
			"message": err.Message,
		})
		return
	}
	log.Printf("[handler:AlimentoHandler] Alimentos obtenidos para el usuario: %s, Cantidad: %d", user.Codigo, len(alimentos))

	c.JSON(http.StatusOK, alimentos)
}
func (handler *AlimentoHandler) ObtenerAlimentoPorID(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))
	id := c.Param("id")
	log.Printf("[handler:AlimentoHandler] Iniciando obtención de alimento por ID: %s para el usuario: %s", id, user.Codigo)
	resultado, err := handler.alimentoService.ObtenerAlimentoPorID(id, user.Codigo)
	if err == nil {
		log.Printf("[handler:AlimentoHandler] Alimento: %s obtenido para el usuario: %s", resultado.NombreAlimento, user.Codigo)
		c.JSON(http.StatusCreated, resultado)
	} else {
		log.Printf("[handler:AlimentoHandler] Error al obtener alimento por ID: %s para el usuario: %s", id, user.Codigo)
		c.JSON(http.StatusBadRequest, err)
	}
}

func (handler *AlimentoHandler) InsertarAlimento(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))
	var alimento dto.Alimento
	log.Printf("[handler:AlimentoHandler] Iniciando inserción de alimento: %s para el usuario: %s", alimento.NombreAlimento, user.Codigo)
	if err := c.ShouldBindJSON(&alimento); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resultado, err := handler.alimentoService.InsertarAlimento(&alimento, user.Codigo)

	if err == nil {
		log.Printf("[handler:AlimentoHandler] Alimento: %s insertado con exito por el usuario: %s", alimento.NombreAlimento, user.Codigo)
		c.JSON(http.StatusCreated, resultado)
	} else {
		log.Printf("[handler:AlimentoHandler] Error al insertar el alimento: %s para el usuario: %s", alimento.NombreAlimento, user.Codigo)
		c.JSON(http.StatusBadRequest, err)
	}
}

func (handler *AlimentoHandler) ModificarAlimento(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))
	var alimento dto.Alimento
	log.Printf("[handler:AlimentoHandler] Iniciando modificación de alimento: %s para el usuario: %s", alimento.NombreAlimento, user.Codigo)
	if err := c.ShouldBindJSON(&alimento); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	alimento.IDAlimento = c.Param("id")
	resultado, err := handler.alimentoService.ModificarAlimento(&alimento, user.Codigo)

	if err == nil {
		log.Printf("[handler:AlimentoHandler] Alimento: %s modifcado exitosamente por el usuario: %s", alimento.NombreAlimento, user.Codigo)
		c.JSON(http.StatusCreated, resultado)
	} else {
		log.Printf("[handler:AlimentoHandler] Error al modificar el alimento: %s por el usuario: %s", alimento.NombreAlimento, user.Codigo)

		c.JSON(http.StatusBadRequest, err)
	}
}

func (handler *AlimentoHandler) EliminarAlimento(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))
	id := c.Param("id")
	log.Printf("[handler:AlimentoHandler] Iniciando eliminacion del alimento: %s para el usuario: %s", id, user.Codigo)
	resultado, err := handler.alimentoService.EliminarAlimento(id)

	if err == nil {
		log.Printf("[handler:AlimentoHandler] Alimento: %s eliminado por el usuario: %s", id, user.Codigo)
		c.JSON(http.StatusCreated, resultado)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}

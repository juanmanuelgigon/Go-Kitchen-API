package handlers

import (
	"TPFINAL-GINCITO/dto"
	"TPFINAL-GINCITO/services"
	"TPFINAL-GINCITO/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RecetaHandler struct {
	recetaService services.RecetaInterface
}

func NewRecetaHandler(recetaService services.RecetaInterface) *RecetaHandler {
	return &RecetaHandler{
		recetaService: recetaService,
	}
}

func (handler *RecetaHandler) ObtenerRecetas(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))
	filtroMomento := c.Query("momento")
	filtroTipoProducto := c.Query("tipo")
	filtroNombreProducto := c.Query("nombre")
	log.Printf("[handler:RecetaHandler] Filtros - Momento: %s, Tipo: %s, Nombre: %s",
		filtroMomento, filtroTipoProducto, filtroNombreProducto)

	recetas, _ := handler.recetaService.ObtenerRecetas(filtroMomento, filtroNombreProducto, filtroTipoProducto, user.Codigo)
	if recetas == nil {
		recetas = []*dto.Receta{}
	}

	log.Printf("[handler:RecetaHandler][method:ObtenerRecetas][cantidad:%d][user:%s]", len(recetas), user.Codigo)

	c.JSON(http.StatusOK, recetas)
}

func (handler *RecetaHandler) ObtenerRecetaPorID(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))
	id := c.Param("id")
	log.Printf("[handler:RecetaHandler] Iniciando obtencion de receta por el ID: %s para el usuario: %s", id, user.Codigo)
	resultado, err := handler.recetaService.ObtenerRecetaPorID(id, user.Codigo)
	if err == nil {
		log.Printf("[handler:RecetaHandler] Receta con el ID: %s obtenida para el usuario: %s", id, user.Codigo)
		c.JSON(http.StatusCreated, resultado)
	} else {
		log.Printf("[handler:RecetaHandler] Error al obtener la receta con el ID: %s para el usuario: %s", id, user.Codigo)
		c.JSON(http.StatusBadRequest, err)
	}
}

func (handler *RecetaHandler) InsertarReceta(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))
	var receta dto.Receta
	log.Printf("[handler:RecetaHandler] Iniciando insercion de receta: %s para el usuario: %s", receta.NombreReceta, user.Codigo)

	if err := c.ShouldBindJSON(&receta); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resultado, err := handler.recetaService.InsertarReceta(&receta, user.Codigo)
	if err == nil {
		log.Printf("[handler:RecetaHandler] Receta: %s insertada con exito por el usuario: %s", receta.NombreReceta, user.Codigo)
		c.JSON(http.StatusCreated, resultado)
	} else {
		log.Printf("[handler:RecetaHandler] Error en insercion de receta: %s para el usuario: %s", receta.NombreReceta, user.Codigo)
		c.JSON(http.StatusBadRequest, err)
	}
}

func (handler *RecetaHandler) ModificarReceta(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))
	var receta dto.Receta
	log.Printf("[handler:RecetaHandler] Iniciando modificacion de receta: %s para el usuario: %s", receta.NombreReceta, user.Codigo)

	if err := c.ShouldBindJSON(&receta); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	receta.ID = c.Param("id")
	resultado, err := handler.recetaService.ModificarReceta(&receta, user.Codigo)

	if err == nil {
		log.Printf("[handler:RecetaHandler] Modificacion de receta: %s exitosa para el usuario: %s", receta.NombreReceta, user.Codigo)
		c.JSON(http.StatusCreated, resultado)
	} else {
		log.Printf("[handler:RecetaHandler] Error al modificar la receta: %s para el usuario: %s", receta.NombreReceta, user.Codigo)
		c.JSON(http.StatusBadRequest, err)
	}
}

func (handler *RecetaHandler) EliminarReceta(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))
	id := c.Param("id")
	log.Printf("[handler:RecetaHandler] Iniciando eliminación de receta con ID: %s para el usuario: %s", id, user.ID)
	resultado, err := handler.recetaService.EliminarReceta(id)
	if err == nil {
		log.Printf("[handler:RecetaHandler] Receta eliminada con exito con ID: %s por el usuario: %s", id, user.ID)
		c.JSON(http.StatusCreated, resultado)
	} else {
		log.Printf("[handler:RecetaHandler] Error al eliminar receta con ID: %s por el usuario: %s", id, user.ID)
		c.JSON(http.StatusBadRequest, err)
	}
}

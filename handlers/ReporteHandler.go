package handlers

import (
	"TPFINAL-GINCITO/dto"
	"TPFINAL-GINCITO/services"
	"TPFINAL-GINCITO/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReporteHandler struct {
	recetaService   services.RecetaInterface
	alimentoService services.AlimentoInterface
	compraService   services.CompraInterface
}

func NewReporteHandler(recetaService services.RecetaInterface, alimentoService services.AlimentoInterface, compraService services.CompraInterface) *ReporteHandler {
	return &ReporteHandler{
		recetaService:   recetaService,
		alimentoService: alimentoService,
		compraService:   compraService,
	}
}

func (handler *ReporteHandler) ObtenerRecetasPorMomento(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))
	log.Printf("[handler:ReporteHandler] Obteniendo los reportes de recetas por momento para el usuario: " + user.Codigo)

	resultado, err := handler.recetaService.ObtenerRecetasPorMomento(user.Codigo)
	if resultado == nil {
		resultado = []dto.ElementoGrafico{}
	}
	if err == nil {
		log.Printf("[handler:ReporteHandler] Reportes de recetas por momento obtenidos para el usuario: %s", user.Codigo)
		c.JSON(http.StatusCreated, resultado)
	} else {
		log.Printf("[handler:ReporteHandler] Error al obtener los reportes de recetas por momento para el usuario: %s", user.Codigo)
		c.JSON(http.StatusBadRequest, err)
	}
}
func (handler *ReporteHandler) ObtenerRecetasPorTipoAlimento(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))
	log.Printf("[handler:ReporteHandler] Obteniendo los reportes de recetas por tipo de alimento para el usuario: " + user.Codigo)

	resultado, err := handler.recetaService.ObtenerRecetasPorTipoAlimento(user.Codigo)
	if resultado == nil {
		resultado = []dto.ElementoGrafico{}
	}
	if err == nil {
		log.Printf("[handler:ReporteHandler] Reportes de recetas por tipo de alimento obtenidos para el usuario: %s", user.Codigo)
		c.JSON(http.StatusCreated, resultado)
	} else {
		log.Printf("[handler:ReporteHandler] Error al obtener los reportes de recetas por tipo de alimento para el usuario: %s", user.Codigo)
		c.JSON(http.StatusBadRequest, err)
	}
}
func (handler *ReporteHandler) ObtenerCostosMensuales(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))
	log.Printf("[handler:ReporteHandler] Obteniendo los reportes de costos mensuales para el usuario: " + user.Codigo)

	resultado, err := handler.compraService.ObtenerDatosCompras(user.Codigo)
	if resultado == nil {
		resultado = []dto.ElementoGrafico{}
	}
	if err == nil {
		log.Printf("[handler:ReporteHandler] Reportes de costos mensuales obtenidos para el usuario: %s", user.Codigo)
		c.JSON(http.StatusCreated, resultado)
	} else {
		log.Printf("[handler:ReporteHandler] Error al obtener los reportes de costos mensuales para el usuario: %s", user.Codigo)
		c.JSON(http.StatusBadRequest, err)
	}
}

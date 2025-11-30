package handlers

import (
	"TPFINAL-GINCITO/dto"
	"TPFINAL-GINCITO/services"
	"TPFINAL-GINCITO/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CompraHandler struct {
	compraService services.CompraInterface
}

func NewCompraHandler(compraService services.CompraInterface) *CompraHandler {
	return &CompraHandler{
		compraService: compraService,
	}
}

func (handler *CompraHandler) ObtenerCompras(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))
	filtroTipoProducto := c.Query("tipo")
	filtroNombreProducto := c.Query("nombre")
	log.Printf("[handler:CompraHandler] Filtros - Nombre: %s, Tipo: %s", filtroNombreProducto, filtroTipoProducto)
	compras, err := handler.compraService.ObtenerAlimentosCompra(filtroNombreProducto, filtroTipoProducto, user.Codigo)
	if err != nil {
		log.Printf("[handler:CompraHandler] Error al obtener la compra para el usuario %s: %s", user.Codigo, err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Code,
			"message": err.Message,
		})
		return
	}
	log.Printf("[handler:CompraHandler] Cantidad de alimentos a comprar obtenidos para el usuario: %s, Cantidad: %d", user.Codigo, len(compras.AlimentosAComprar))

	c.JSON(http.StatusOK, compras)
}

func (handler *CompraHandler) GenerarNuevaCompra(c *gin.Context) {
	var requestBody struct {
		Productos []string `json:"productos"`
	}

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	user := dto.NewUser(utils.GetUserInfoFromContext(c))

	resultado, err := handler.compraService.GenerarCompra(requestBody.Productos, user.Codigo)

	if err == nil {
		c.JSON(http.StatusCreated, resultado)
	} else {
		log.Printf("Error [handler:CompraHandler] Compra: %s no generada para el usuario: %s", resultado.ID, user.Codigo)
		c.JSON(http.StatusBadRequest, err)
	}
}

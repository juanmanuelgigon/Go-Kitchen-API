package main

import (
	"TPFINAL-GINCITO/clients"
	"TPFINAL-GINCITO/handlers"
	"TPFINAL-GINCITO/middlewares"
	"TPFINAL-GINCITO/repositories"
	"TPFINAL-GINCITO/services"

	"github.com/gin-gonic/gin"
)

var (
	alimentoHandler *handlers.AlimentoHandler
	recetaHandler   *handlers.RecetaHandler
	compraHandler   *handlers.CompraHandler
	reporteHandler  *handlers.ReporteHandler
	router          *gin.Engine
)

func main() {
	router = gin.Default()
	//Iniciar objetos de handler
	dependencies()
	//Iniciar rutas
	mappingRoutes()

	router.Run(":8080")
}
func mappingRoutes() {
	//middleware para permitir peticiones del mismo server localhost

	//cliente para api externa
	var authClient clients.AuthClientInterface
	authClient = clients.NewAuthClient()
	//creacion de middleware de autenticacion
	authMiddleware := middlewares.NewAuthMiddleware(authClient)
	//Uso del middleware para todas las rutas del grupo
	router.Use(middlewares.CORSMiddleware())
	router.Use(authMiddleware.ValidateToken)

	//Listado de rutas
	groupAlimento := router.Group("/alimentos")

	groupAlimento.GET("", alimentoHandler.ObtenerAlimentos)
	groupAlimento.GET("/:id", alimentoHandler.ObtenerAlimentoPorID)
	groupAlimento.POST("", alimentoHandler.InsertarAlimento)
	groupAlimento.PUT("/:id", alimentoHandler.ModificarAlimento)
	groupAlimento.DELETE("/:id", alimentoHandler.EliminarAlimento)

	groupReceta := router.Group("/recetas")
	groupReceta.GET("", recetaHandler.ObtenerRecetas)
	groupReceta.GET("/:id", recetaHandler.ObtenerRecetaPorID)
	groupReceta.POST("", recetaHandler.InsertarReceta)
	groupReceta.DELETE("/:id", recetaHandler.EliminarReceta)
	groupReceta.PUT("/:id", recetaHandler.ModificarReceta)

	groupComida := router.Group("/compras")
	groupComida.GET("", compraHandler.ObtenerCompras)
	groupComida.POST("/", compraHandler.GenerarNuevaCompra)

	groupReportes := router.Group("/reportes")
	groupReportes.GET("/momento", reporteHandler.ObtenerRecetasPorMomento)
	groupReportes.GET("/tipo", reporteHandler.ObtenerRecetasPorTipoAlimento)
	groupReportes.GET("/costo", reporteHandler.ObtenerCostosMensuales)

}

func dependencies() {
	//Definicion de variables de interface
	var database repositories.DB
	var alimentoRepository repositories.AlimentoRepositoryInterface
	var recetaRepository repositories.RecetaRepositoryInterface
	var compraRepository repositories.CompraRepositoryInterface
	var alimentoService services.AlimentoInterface
	var recetaService services.RecetaInterface
	var compraService services.CompraInterface

	//Creamos los objetos reales y los pasamos como parametro
	database = repositories.NewMongoDB()
	alimentoRepository = repositories.NewAlimentoRepository(database)
	recetaRepository = repositories.NewRecetaRepository(database)
	compraRepository = repositories.NewCompraRepository(database)
	alimentoService = services.NewAlimentoService(alimentoRepository)
	recetaService = services.NewRecetaService(recetaRepository, alimentoRepository)
	compraService = services.NewCompraService(compraRepository, alimentoRepository)
	alimentoHandler = handlers.NewAlimentoHandler(alimentoService)
	recetaHandler = handlers.NewRecetaHandler(recetaService)
	compraHandler = handlers.NewCompraHandler(compraService)
	reporteHandler = handlers.NewReporteHandler(recetaService, alimentoService, compraService)

}

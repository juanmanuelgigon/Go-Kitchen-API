> **Trabajo Práctico Integrador | Programación II**
> Sistema Full Stack de gestión de inventario, recetas y control de costos desarrollado en **Golang**.

Este proyecto fue desarrollado en equipo como cierre de la materia **Programación II**. Es una solución integral para administrar los recursos culinarios de un hogar o negocio pequeño, permitiendo gestionar el stock de alimentos, crear recetas verificando la disponibilidad de ingredientes en tiempo real y generar reportes financieros.

Diseñado con una **Arquitectura en Capas** y totalmente contenerizado con **Docker**.

## 🚀 Key Features

* **Gestión Inteligente de Stock:** CRUD completo de alimentos con control de cantidad mínima y precios.
* **Motor de Recetas:** Creación de recetas que valida automáticamente si existe stock suficiente de los ingredientes necesarios.
* **Reportes y Métricas:**
    * Análisis de recetas por momento del día (Desayuno, Almuerzo, Cena).
    * Cálculo de costos mensuales.
    * Distribución por tipo de alimento.
* **Seguridad:** Middleware de autenticación personalizado integrado con API externa de usuarios.
* **Infraestructura:** Despliegue automatizado mediante Docker Compose (Backend + DB + Frontend Server).

## 🛠 Tech Stack

### Backend
* **Lenguaje:** Go (Golang)
* **Framework Web:** [Gin Gonic](https://github.com/gin-gonic/gin) (High performance HTTP web framework)
* **Arquitectura:** MVC / Clean Architecture (Handlers -> Services -> Repositories)

### Database
* **Motor:** MongoDB
* **Driver:** Mongo Go Driver

### Frontend
* **Tecnologías:** HTML5, CSS3, JavaScript (Vanilla).
* **Estilos:** Bootstrap.
* **Server:** Nginx (como Reverse Proxy y servidor de estáticos).

### DevOps & Tools
* **Docker & Docker Compose**

## 📂 Arquitectura del Proyecto

El código sigue una estructura modular para facilitar la escalabilidad y el mantenimiento:

```text
├── handlers/      # Controladores HTTP (Gin context)
├── services/      # Lógica de negocio y validaciones
├── repositories/  # Acceso a datos (MongoDB queries)
├── models/        # Definición de estructuras de datos
├── dto/           # Data Transfer Objects
├── middlewares/   # Auth y CORS
└── docker-compose.yml

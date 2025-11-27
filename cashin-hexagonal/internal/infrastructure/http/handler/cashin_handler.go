package handler

import (
	"cashin-hexagonal/internal/application/dto"
	"cashin-hexagonal/internal/application/usecase"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CashInHandler struct {
	createUseCase *usecase.CreateCashInUseCase
}

func HealthCheckHandler(c *gin.Context) {
	message := "Servicio OK"
	c.String(http.StatusOK, message)
}

func NewCashInHandler(createUseCase *usecase.CreateCashInUseCase) *CashInHandler {
	return &CashInHandler{createUseCase: createUseCase}
}

func RegisterCashInRoutes(rg *gin.RouterGroup, createUseCase *usecase.CreateCashInUseCase) {
	h := NewCashInHandler(createUseCase)
	rg.POST("/cashin", h.CreateCashIn)
}
func RegisterHealthCheckRoutes(rg *gin.RouterGroup) {
	rg.GET("/hc", HealthCheckHandler)
}

// CreateCashIn godoc
// @Summary      Registra un nuevo CashIn
// @Tags         cashin
// @Accept       json
// @Produce      json
// @Param        request body dto.CreateCashInRequest true "Datos del CashIn"
// @Success      201  {object}  map[string]string
// @Router       /cashin [post]
func (h *CashInHandler) CreateCashIn(c *gin.Context) {
	var req dto.CreateCashInRequest
	log.Println("[EXAGONAL]Received request to create CashIn")
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.createUseCase.Execute(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id, "message": "CashIn registrado correctamente"})
}

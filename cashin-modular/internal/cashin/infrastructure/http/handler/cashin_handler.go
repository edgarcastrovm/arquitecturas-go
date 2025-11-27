package http

import (
	"cashin-modular/internal/cashin/application/dto"
	"cashin-modular/internal/cashin/application/usecase"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CashInHandler struct {
	createUC *usecase.CreateCashInUseCase
	getUC    *usecase.GetCashInUseCase
}

func NewCashInHandler(createUC *usecase.CreateCashInUseCase, getUC *usecase.GetCashInUseCase) *CashInHandler {
	return &CashInHandler{createUC: createUC, getUC: getUC}
}

func (h *CashInHandler) Create(c *gin.Context) {
	var req dto.CreateCashInRequest
	log.Println("[MODULAR]Received request to create CashIn")
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	out, err := h.createUC.Execute(c.Request.Context(), req.AccountID, req.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, out)
}

func (h *CashInHandler) Get(c *gin.Context) {
	id := c.Param("id")
	cashin, err := h.getUC.Execute(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "CashIn not found"})
		return
	}

	c.JSON(http.StatusOK, cashin)
}

func (h *CashInHandler) GetHc(c *gin.Context) {
	message := "Servicio OK"
	c.String(http.StatusOK, message)
}

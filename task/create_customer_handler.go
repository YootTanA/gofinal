package task

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateCustomerHandler(c *gin.Context) {

	customer := Customer{}
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	row := h.DB.QueryRow("INSERT INTO customers (name, email, status) values ($1, $2, $3) RETURNING id", customer.Name, customer.Email, customer.Status)

	err := row.Scan(&customer.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, customer)
}

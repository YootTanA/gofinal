package task

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetCustomerById(c *gin.Context) {

	id := c.Param("id")
	stmt, err := h.DB.Prepare("SELECT * FROM customers where id=$1")

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	row := stmt.QueryRow(id)

	customer := &Customer{}

	err = row.Scan(&customer.ID, &customer.Name, &customer.Email, &customer.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, customer)

}

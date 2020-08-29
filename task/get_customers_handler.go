package task

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetCustomers(c *gin.Context) {

	status := c.Query("status")
	fmt.Println(status)

	stmt, err := h.DB.Prepare("SELECT * FROM customers")
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	rows, err := stmt.Query()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	customers := []Customer{}
	for rows.Next() {
		customer := Customer{}

		err := rows.Scan(&customer.ID, &customer.Name, &customer.Email, &customer.Status)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		customers = append(customers, customer)
	}

	c.JSON(http.StatusOK, customers)
}

package task

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) PutCustomerById(c *gin.Context) {

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

	if err := c.ShouldBindJSON(customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customer.ID, err = strconv.Atoi(id)
	if err != nil {
		fmt.Println("Can not prase String to Integer")
	}

	stmt, err = h.DB.Prepare("UPDATE customers SET name=$2, email=$3, status=$4 WHERE id=$1")
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if _, err := stmt.Exec(id, customer.Name, customer.Email, customer.Status); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, customer)

}

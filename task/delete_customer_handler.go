package task

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) DeleteCustomerById(c *gin.Context) {

	id := c.Param("id")
	stmt, err := h.DB.Prepare("DELETE FROM customers where id=$1")
	if err != nil {
		log.Fatal("can't execute delete statment", err)
	}

	if _, err := stmt.Exec(id); err != nil {
		log.Fatal("can't execute delete statment", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "customer deleted"})

}

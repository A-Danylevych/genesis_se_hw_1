package handler

import (
	"net/http"

	btcapi "github.com/A-Danylevych/btc-api"

	"github.com/gin-gonic/gin"
)

//Processing user creation requests. Need email and password, it returns the id of the created user or an error
func (h *Handler) create(c *gin.Context) {
	var input btcapi.User
	if err := c.BindJSON(&input); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

//Processing autorization requests. Need email and password, it returns the authorization token of the user or an error
func (h *Handler) logIn(c *gin.Context) {
	var input btcapi.User
	if err := c.BindJSON(&input); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

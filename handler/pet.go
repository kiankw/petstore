package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kiankw/petstore/model"
)

func (h *handlerStruct) GetAllPets(c *gin.Context) {
	pets, err := h.srv.GetAllPets(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusBadRequest, pets)
	} else {
		c.JSON(http.StatusOK, pets)
	}
}

func (h *handlerStruct) CreatePet(c *gin.Context) {
	pet := &model.Pet{}
	c.BindJSON(pet)
	pet, err := h.srv.CreatePet(c.Request.Context(), pet.Name)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	} else {
		c.JSON(http.StatusOK, pet)
	}
}

func (h *handlerStruct) DeletePet(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	err = h.srv.DeletePet(c.Request.Context(), id)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	} else {
		c.String(http.StatusOK, "")
	}
}

func (h *handlerStruct) GetPetByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	pet, err := h.srv.GetPetByID(c.Request.Context(), id)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	} else {
		c.JSON(http.StatusOK, pet)
	}
}

func (h *handlerStruct) UpdatePet(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	pet := &model.Pet{}
	c.BindJSON(pet)
	err = h.srv.UpdatePet(c.Request.Context(), id, pet.Name)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	} else {
		c.String(http.StatusOK, "")
	}
}

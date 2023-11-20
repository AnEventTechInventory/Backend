package api

import (
	"github.com/AnEventTechInventory/Backend/pkg/registry"
	"github.com/AnEventTechInventory/Backend/pkg/storageManager"
	"github.com/gin-gonic/gin"
	"net/http"
)

type manufacturerRequestHandler struct {
	requestInterface
	store *storageManager.ManufacturerStorageManager
}

func newManufacturerRequestHandler() *manufacturerRequestHandler {
	return &manufacturerRequestHandler{
		store: storageManager.NewManufacturerStorageManager(),
	}
}

func (handler *manufacturerRequestHandler) listManufacturer(context *gin.Context) {
	list, err := handler.store.List()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(list) == 0 {
		context.Status(http.StatusNoContent)
		return
	}
	context.JSON(http.StatusOK, list)
}

func (handler *manufacturerRequestHandler) getManufacturer(context *gin.Context) {
	id := context.Param("id")
	var manufacturer *registry.Manufacturer
	manufacturer, err := handler.store.Get(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if manufacturer == nil {
		context.String(http.StatusInternalServerError, "The database lookup returned nil.")
	}
	context.JSON(http.StatusOK, registry.JsonManufacturer{
		ID:          manufacturer.ID.String(),
		Name:        manufacturer.Name,
		Description: manufacturer.Description,
	})
}

func (handler *manufacturerRequestHandler) createManufacturer(context *gin.Context) {
	var newManufacturer registry.Manufacturer
	err := context.BindJSON(&newManufacturer)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = handler.store.Add(&newManufacturer)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"id": newManufacturer.ID})
}

func (handler *manufacturerRequestHandler) updateManufacturer(context *gin.Context) {
	var newManufacturer registry.Manufacturer
	err := context.BindJSON(&newManufacturer)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = handler.store.Update(&newManufacturer)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func (handler *manufacturerRequestHandler) deleteManufacturer(context *gin.Context) {
	id := context.Param("id")
	err := handler.store.Delete(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	context.Status(http.StatusAccepted)
}

var manufacturerRequestHandlerInstance = newManufacturerRequestHandler()

func RegisterManufacturers(context *gin.Engine) {
	manufacturerGroup := context.Group("/manufacturer")

	manufacturerGroup.GET("", manufacturerRequestHandlerInstance.listManufacturer)
	manufacturerGroup.GET("/:id", manufacturerRequestHandlerInstance.getManufacturer)
	manufacturerGroup.POST("", manufacturerRequestHandlerInstance.createManufacturer)
	manufacturerGroup.PUT("/:id", manufacturerRequestHandlerInstance.updateManufacturer)
	manufacturerGroup.DELETE("/:id", manufacturerRequestHandlerInstance.deleteManufacturer)
}
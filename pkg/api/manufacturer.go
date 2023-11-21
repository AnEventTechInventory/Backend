package api

import (
	"github.com/AnEventTechInventory/Backend/pkg/database"
	"github.com/AnEventTechInventory/Backend/pkg/registry"
	"github.com/AnEventTechInventory/Backend/pkg/storageManager"
	"github.com/gin-gonic/gin"
	"net/http"
)

type manufacturerRequestHandler struct {
	store *storageManager.ManufacturerStorageManager
}

func newManufacturerRequestHandler() *manufacturerRequestHandler {
	return &manufacturerRequestHandler{
		store: storageManager.NewManufacturerStorageManager(),
	}
}

func (handler *manufacturerRequestHandler) list(context *gin.Context) {
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

func (handler *manufacturerRequestHandler) get(context *gin.Context) {
	id := context.Param("id")
	var manufacturer *registry.Manufacturer
	manufacturer, err := handler.store.Get(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := manufacturer.Validate(database.Get()); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, registry.JsonManufacturer{
		ID:          manufacturer.ID.String(),
		Name:        manufacturer.Name,
		Description: manufacturer.Description,
	})
}

func (handler *manufacturerRequestHandler) create(context *gin.Context) {
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

func (handler *manufacturerRequestHandler) update(context *gin.Context) {
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

func (handler *manufacturerRequestHandler) delete(context *gin.Context) {
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

	manufacturerGroup.GET("", manufacturerRequestHandlerInstance.list)
	manufacturerGroup.GET("/:id", manufacturerRequestHandlerInstance.get)
	manufacturerGroup.POST("", manufacturerRequestHandlerInstance.create)
	manufacturerGroup.PUT("/:id", manufacturerRequestHandlerInstance.update)
	manufacturerGroup.DELETE("/:id", manufacturerRequestHandlerInstance.delete)
}

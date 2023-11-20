package api

import (
	"github.com/AnEventTechInventory/Backend/pkg/registry"
	"github.com/AnEventTechInventory/Backend/pkg/storageManager"
	"github.com/gin-gonic/gin"
	"net/http"
)

type locationRequestHandler struct {
	requestInterface
	store *storageManager.LocationStorageManager
}

func newLocationRequestHandler() *locationRequestHandler {
	return &locationRequestHandler{
		store: storageManager.NewLocationStorageManager(),
	}
}

func (handler *locationRequestHandler) listLocation(context *gin.Context) {
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

func (handler *locationRequestHandler) getLocation(context *gin.Context) {
	id := context.Param("id")
	var location *registry.Location
	location, err := handler.store.Get(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if location == nil {
		context.String(http.StatusInternalServerError, "The database lookup returned nil.")
		return
	}
	context.JSON(http.StatusOK, registry.JsonLocation{
		ID:          location.ID.String(),
		Name:        location.Name,
		Description: location.Description,
	})
}

func (handler *locationRequestHandler) createLocation(context *gin.Context) {
	var newLocation registry.Location
	err := context.BindJSON(&newLocation)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = handler.store.Add(&newLocation)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"id": newLocation.ID})
}

func (handler *locationRequestHandler) updateLocation(context *gin.Context) {
	var newLocation registry.Location
	err := context.BindJSON(&newLocation)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = handler.store.Update(&newLocation)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func (handler *locationRequestHandler) deleteLocation(context *gin.Context) {
	id := context.Param("id")
	err := handler.store.Delete(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	context.Status(http.StatusAccepted)
}

var locationRequestHandlerInstance = newLocationRequestHandler()

func RegisterLocations(context *gin.Engine) {
	locationGroup := context.Group("/location")

	locationGroup.GET("", locationRequestHandlerInstance.listLocation)
	locationGroup.GET("/:id", locationRequestHandlerInstance.getLocation)
	locationGroup.POST("", locationRequestHandlerInstance.createLocation)
	locationGroup.PUT("/:id", locationRequestHandlerInstance.updateLocation)
	locationGroup.DELETE("/:id", locationRequestHandlerInstance.deleteLocation)
}

package api

import (
	"github.com/AnEventTechInventory/Backend/pkg/registry"
	"github.com/AnEventTechInventory/Backend/pkg/storageManager"
	"github.com/gin-gonic/gin"
	"net/http"
)

type locationRequestHandler struct {
	store *storageManager.LocationStorageManager
}

func newLocationRequestHandler() *locationRequestHandler {
	return &locationRequestHandler{
		store: storageManager.NewLocationStorageManager(),
	}
}

func (handler *locationRequestHandler) list(context *gin.Context) {
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

func (handler *locationRequestHandler) get(context *gin.Context) {
	id := context.Param("id")
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

func (handler *locationRequestHandler) create(context *gin.Context) {
	newLocation := &registry.Location{}
	err := context.BindJSON(newLocation)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = handler.store.Add(newLocation)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"id": newLocation.ID})
}

func (handler *locationRequestHandler) update(context *gin.Context) {
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

func (handler *locationRequestHandler) delete(context *gin.Context) {
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

	locationGroup.GET("", locationRequestHandlerInstance.list)
	locationGroup.GET("/:id", locationRequestHandlerInstance.get)
	locationGroup.POST("", locationRequestHandlerInstance.create)
	locationGroup.PUT("/:id", locationRequestHandlerInstance.update)
	locationGroup.DELETE("/:id", locationRequestHandlerInstance.delete)
}

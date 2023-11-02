package api

import (
	"github.com/AnEventTechInventory/Backend/pkg/database"
	"github.com/AnEventTechInventory/Backend/pkg/registry"
	"github.com/AnEventTechInventory/Backend/pkg/storageManager"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type deviceRequestHandler struct {
	requestInterface
	store *storageManager.DeviceStorageManager
}

func newDeviceRequestHandler(db *gorm.DB) *deviceRequestHandler {
	return &deviceRequestHandler{
		store: storageManager.NewDeviceStorageManager(db),
	}
}

func (handler *deviceRequestHandler) list(context *gin.Context) {
	list, err := handler.store.List()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// if there are no devices present, return an empty list
	if len(list) == 0 {
		context.Status(http.StatusNoContent)
		return
	}
	context.JSON(http.StatusOK, list)
}

func (handler *deviceRequestHandler) get(context *gin.Context) {
	id := context.Param("id")
	device, err := handler.store.Get(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, device)
}

func (handler *deviceRequestHandler) create(context *gin.Context) {
	var newDevice registry.Device
	err := context.BindJSON(&newDevice)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = handler.store.Add(&newDevice)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// confirm and send the new device id
	context.JSON(http.StatusCreated, gin.H{"id": newDevice.Id})
}

func (handler *deviceRequestHandler) update(context *gin.Context) {
	var newDevice registry.Device
	err := context.BindJSON(&newDevice)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = handler.store.Update(&newDevice)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.Status(http.StatusAccepted)
}

func (handler *deviceRequestHandler) delete(context *gin.Context) {
	id := context.Param("id")
	err := handler.store.Delete(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.Status(http.StatusAccepted)
}

var devicesRequestHandler = newDeviceRequestHandler(database.Database)

func RegisterDevices(context *gin.Engine) {
	devicesGroup := context.Group("/devices")

	devicesGroup.GET("", devicesRequestHandler.list)
	devicesGroup.GET("/:id", devicesRequestHandler.get)
	devicesGroup.POST("", devicesRequestHandler.create)
	devicesGroup.PUT("/:id", devicesRequestHandler.update)
	devicesGroup.DELETE("/:id", devicesRequestHandler.delete)

}

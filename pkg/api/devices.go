package api

import (
	"github.com/AnEventTechInventory/Backend/pkg/database"
	"github.com/AnEventTechInventory/Backend/pkg/registry"
	"github.com/AnEventTechInventory/Backend/pkg/storageManager"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (handler *deviceHandler) listDevices(context *gin.Context) {
	list, err := handler.store.List()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// if there are no devices present, return an empty list
	if len(list) == 0 {
		context.JSON(http.StatusOK, []string{})
		return
	}
	context.JSON(http.StatusOK, list)
}

func (handler *deviceHandler) getDevice(context *gin.Context) {
	id := context.Param("id")
	device, err := handler.store.Get(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, device)
}

func (handler *deviceHandler) createDevice(context *gin.Context) {
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
	context.JSON(http.StatusOK, gin.H{"id": newDevice.Id})
}

func (handler *deviceHandler) updateDevice(context *gin.Context) {
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
}

func (handler *deviceHandler) deleteDevice(context *gin.Context) {
	id := context.Param("id")
	err := handler.store.Delete(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.Status(http.StatusOK)
}

type deviceHandler struct {
	store storageManager.DeviceStore
}

func newDevicesHandler() *deviceHandler {
	return &deviceHandler{
		store: storageManager.NewDeviceStorageManager(database.Database),
	}
}

var devicesHandler = newDevicesHandler()

func RegisterDevices(context *gin.Engine) {
	devicesGroup := context.Group("/devices")

	devicesGroup.GET("/", devicesHandler.listDevices)
	devicesGroup.GET("/:id", devicesHandler.getDevice)
	devicesGroup.POST("/", devicesHandler.createDevice)
	devicesGroup.PUT("/:id", devicesHandler.updateDevice)
	devicesGroup.DELETE("/:id", devicesHandler.deleteDevice)

}

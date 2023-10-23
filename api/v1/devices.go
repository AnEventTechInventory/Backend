package v1

import (
	"github.com/AnEventTechInventory/Backend/pkg/database"
	"github.com/AnEventTechInventory/Backend/pkg/devices"
	"github.com/AnEventTechInventory/Backend/pkg/registry"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (handler *deviceHandler) listDevices(context *gin.Context) {
	list, err := handler.store.List()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
	store devices.DeviceStore
}

func newDevicesHandler() *deviceHandler {
	return &deviceHandler{
		store: devices.NewDeviceStorageManager(database.Database),
	}
}

var devicesHandler = newDevicesHandler()

func Handler(context *gin.RouterGroup) {

	// Handle devices
	devicesGroup := context.Group("/devices")
	devicesGroup.GET("/", devicesHandler.listDevices)
	devicesGroup.GET("/:id", devicesHandler.getDevice)
	devicesGroup.POST("/", devicesHandler.createDevice)
	devicesGroup.PUT("/:id", devicesHandler.updateDevice)
	devicesGroup.DELETE("/:id", devicesHandler.deleteDevice)

}

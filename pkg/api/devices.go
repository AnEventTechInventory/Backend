package api

import (
	"github.com/AnEventTechInventory/Backend/pkg/database"
	"github.com/AnEventTechInventory/Backend/pkg/registry"
	"github.com/AnEventTechInventory/Backend/pkg/storageManager"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type deviceRequestHandler struct {
	store *storageManager.DeviceStorageManager
}

var _ requestInterface = &deviceRequestHandler{}

func newDeviceRequestHandler() *deviceRequestHandler {
	return &deviceRequestHandler{
		store: storageManager.NewDeviceStorageManager(),
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
	if id == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}
	device, err := handler.store.Get(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, registry.JsonDevice{
		ID:           device.ID.String(),
		Name:         device.Name,
		Description:  device.Description,
		Location:     device.LocationId,
		Manufacturer: device.ManufacturerId,
		Type:         device.TypeId,
		Quantity:     device.Quantity,
		Contents:     strings.Split(device.Contents, ","),
	})
}

func (handler *deviceRequestHandler) create(context *gin.Context) {
	var newDevice registry.JsonDevice
	if err := context.BindJSON(&newDevice); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	dev, err := registry.DeviceFromJson(newDevice, database.Get())
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := handler.store.Add(dev); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// confirm and send the new device id
	context.JSON(http.StatusCreated, gin.H{"id": newDevice.ID})
}

func (handler *deviceRequestHandler) update(context *gin.Context) {
	var newDevice registry.JsonDevice
	err := context.BindJSON(&newDevice)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	dev, err := registry.DeviceFromJson(newDevice, database.Get())
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	err = handler.store.Update(dev)
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

func (handler *deviceRequestHandler) updateContents(context *gin.Context) {
	id := context.Param("id")
	type contentsArray struct {
		Contents []string `json:"contents"`
	}
	var contents contentsArray
	err := context.BindJSON(&contents)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = handler.store.UpdateContents(id, contents.Contents)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.Status(http.StatusAccepted)
}

var devicesRequestHandler = newDeviceRequestHandler()

func RegisterDevices(context *gin.Engine) {
	devicesGroup := context.Group("/devices")

	devicesGroup.GET("", devicesRequestHandler.list)
	devicesGroup.GET("/:id", devicesRequestHandler.get)
	devicesGroup.POST("", devicesRequestHandler.create)
	devicesGroup.PUT("/:id", devicesRequestHandler.update)
	devicesGroup.DELETE("/:id", devicesRequestHandler.delete)

	devicesGroup.POST("/contents/:id", devicesRequestHandler.updateContents)
}

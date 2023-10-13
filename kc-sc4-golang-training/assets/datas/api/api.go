package api

import (
	"gitlab-techlab/techlab/training/golang/gin-samples/datas/items"
	"gitlab-techlab/techlab/training/golang/gin-samples/datas/stores"
	"net/http"
	"time"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// DatasEnvelope contains all information for output
type DatasEnvelope struct {
	Count      int           `json:"count"`
	Datas      []interface{} `json:"datas"`
	Limit      int           `json:"limit"`
	Offset     int           `json:"offset"`
	TotalCount int           `json:"totalCount"`
}

// DatasEnvelope contains all information for output
type ErrorEnvelope struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type MyApi struct {
	localStore stores.Store
}

func New(st stores.Store) *MyApi {
	lst := new(MyApi)
	lst.localStore = st
	return lst
}

// CreateData  adds a data item in the store
// returns HTTP 200 on success and HTTP 404 on failure
func (myapi *MyApi) CreateData(c *gin.Context) {

	var item items.Item
	if err := c.BindJSON(&item); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var timestamp = time.Now()
	item.Envelope.CreatedAt = timestamp
	item.Envelope.ModifiedAt = timestamp

	//set a new uuid
	item.ID = uuid.New().String()

	// add item in store
	if err := myapi.localStore.AddData(item.ID, item); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, item)

}

// GetDatas retrieves data items from data store
// returns HTTP 200 on success and HTTP 404 on failure
func (myapi *MyApi) GetDatas(c *gin.Context) {
	var offset int = 0
	var limit int = 10
	var err error

	stOffset := c.Query("offset")
	if stOffset != "" {
		offset, err = strconv.Atoi(stOffset)
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
	}

	stlimit := c.Query("limit")
	if stlimit != "" {
		limit, err = strconv.Atoi(stlimit)
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
	}

	if offset < 0 {
		var eEnv *ErrorEnvelope = new(ErrorEnvelope)
		eEnv.Code = http.StatusNotFound
		eEnv.Message = "Offset parameter cannot be negative"
		c.AbortWithStatusJSON(http.StatusNotFound, eEnv)
		return
	}
	if limit < 0 || limit > 20 {
		var eEnv *ErrorEnvelope = new(ErrorEnvelope)
		eEnv.Code = http.StatusNotFound
		eEnv.Message = "Limit parameter cannot be negative or exceed 20"
		c.AbortWithStatusJSON(http.StatusNotFound, eEnv)
		return
	}

	itemsInterfaceArray, size, storeErr := myapi.localStore.GetDatas(offset, limit)
	if storeErr != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	var de *DatasEnvelope = new(DatasEnvelope)
	de.Datas = itemsInterfaceArray
	de.Count = len(itemsInterfaceArray)
	de.Limit = limit
	de.Offset = offset
	de.TotalCount = size

	c.JSON(http.StatusOK, de)

}

// GetData returns a specific item from the store
// returns HTTP 200 on success and HTTP 404 on failure
func (myapi *MyApi) GetData(c *gin.Context) {
	id := c.Param("id")
	var dataUuid, err = uuid.Parse(id)
	if err != nil {
		var eEnv *ErrorEnvelope = new(ErrorEnvelope)
		eEnv.Code = http.StatusNotFound
		eEnv.Message = "invalid uid"
		c.AbortWithStatusJSON(http.StatusNotFound, eEnv)
		return
	}

	item, storeErr := myapi.localStore.GetData(dataUuid.String())
	if storeErr != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	if item == nil {
		c.JSON(http.StatusNotFound, item)
	}
	c.JSON(http.StatusOK, item)

}

// UpdateData updates a specific item in the store
// returns HTTP 200 on success and HTTP 404 on failure
func (myapi *MyApi) UpdateData(c *gin.Context) {
	id := c.Param("id")
	var dataUuid, uuidErr = uuid.Parse(id)
	if uuidErr != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	var sentData items.Data
	if bindErr := c.BindJSON(&sentData); bindErr != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	itemInterface, storeErr := myapi.localStore.GetData(dataUuid.String())
	if storeErr != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	var timestamp = time.Now()
	var item, assertOK = itemInterface.(items.Item)
	if !assertOK {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	item.Envelope.ModifiedAt = timestamp
	item.Data = sentData

	storeErr = myapi.localStore.UpdateData(dataUuid.String(), item)
	if storeErr != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, item)

}

// DeleteData deletes a specific item in the store
// returns HTTP 200 on success and HTTP 404 on failure
func (myapi *MyApi) DeleteData(c *gin.Context) {
	id := c.Param("id")
	var dataUuid, err = uuid.Parse(id)
	if err != nil {
		var eEnv *ErrorEnvelope = new(ErrorEnvelope)
		eEnv.Code = http.StatusNotFound
		eEnv.Message = "invalid uid"
		c.AbortWithStatusJSON(http.StatusNotFound, eEnv)
		return
	}

	storeErr := myapi.localStore.DeleteData(dataUuid.String())
	if storeErr != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.Status(http.StatusNoContent)
}

package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

const baseURL = "http://127.0.0.1:8080"

func createSegmentTestWithBadRequestStatus(t *testing.T, size int, unit string) {
	segment := Segment{Size: size, Unit: unit}
	jsonData, _ := json.Marshal(segment)

	resp, err := http.Post(baseURL+"/segments", "application/json", bytes.NewBuffer(jsonData))
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.StatusCode, http.StatusBadRequest)
}

func createSegmentTest(t *testing.T, size int, unit string) Segment {
	segment := Segment{Size: size, Unit: unit}
	jsonData, _ := json.Marshal(segment)

	resp, err := http.Post(baseURL+"/segments", "application/json", bytes.NewBuffer(jsonData))
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.StatusCode, http.StatusCreated)

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	var newSegment Segment
	err = json.Unmarshal(body, &newSegment)
	assert.Equal(t, err, nil)

	return newSegment
}

func TestCreateAndGetSegment(t *testing.T) {
	newSegment := createSegmentTest(t, 10, "bytes")

	assert.Equal(t, newSegment.Size, 10)
	assert.Equal(t, newSegment.Unit, "BYTES")
}

func TestCreateOverSizedSegment(t *testing.T) {
	createSegmentTestWithBadRequestStatus(t, 10^10, "gigabytes")
}

func TestCreateSegmentwithBadUnit(t *testing.T) {
	createSegmentTestWithBadRequestStatus(t, 10, "giga")
}

func TestGetSegment(t *testing.T) {
	newSegment := createSegmentTest(t, 10, "bytes")

	// Test getSegment
	resp, err := http.Get(baseURL + "/segments/" + newSegment.ID.String())
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.StatusCode, http.StatusOK)

	body, _ := ioutil.ReadAll(resp.Body)
	var fetchedSegment Segment
	err = json.Unmarshal(body, &fetchedSegment)
	assert.Equal(t, err, nil)

	assert.Equal(t, fetchedSegment.ID, newSegment.ID)
	//fmt.Printf("fetchedSegment : %s)", fetchedSegment.String())
}

func TestGetNonExistingSegment(t *testing.T) {

	// Test getSegment
	resp, err := http.Get(baseURL + "/segments/" + uuid.New().String())
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.StatusCode, http.StatusNotFound)
}

func TestUpdateSegment(t *testing.T) {
	newSegment := createSegmentTest(t, 10, "bytes")

	// Test updateSegment
	newSegment.Size = 20
	jsonData, _ := json.Marshal(newSegment)
	req, err := http.NewRequest(http.MethodPut, baseURL+"/segments/"+newSegment.ID.String(), bytes.NewBuffer(jsonData))
	resp, err := http.DefaultClient.Do(req)
	assert.Equal(t, err, nil)

	assert.Equal(t, resp.StatusCode, http.StatusOK)

	body, _ := ioutil.ReadAll(resp.Body)
	var updatedSegment Segment
	err = json.Unmarshal(body, &updatedSegment)
	assert.Equal(t, err, nil)
	assert.Equal(t, updatedSegment.Size, newSegment.Size)
}

func TestNonExistingSegment(t *testing.T) {
	newSegment := createSegmentTest(t, 10, "bytes")

	// Test updateSegment
	newSegment.Size = 20
	jsonData, _ := json.Marshal(newSegment)
	req, err := http.NewRequest(http.MethodPut, baseURL+"/segments/"+uuid.New().String(), bytes.NewBuffer(jsonData))
	resp, err := http.DefaultClient.Do(req)
	assert.Equal(t, err, nil)

	assert.Equal(t, resp.StatusCode, http.StatusNotFound)

}

func TestDeleteSegment(t *testing.T) {

	newSegment := createSegmentTest(t, 10, "bytes")

	// Test deleteSegment
	req, err := http.NewRequest(http.MethodDelete, baseURL+"/segments/"+newSegment.ID.String(), nil)
	assert.Equal(t, err, nil)

	resp, err := http.DefaultClient.Do(req)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.StatusCode, http.StatusNoContent)
}

func TestDeleteNonExistingSegment(t *testing.T) {

	createSegmentTest(t, 10, "bytes")

	// Test deleteSegment
	req, err := http.NewRequest(http.MethodDelete, baseURL+"/segments/"+uuid.New().String(), nil)
	assert.Equal(t, err, nil)

	resp, err := http.DefaultClient.Do(req)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.StatusCode, http.StatusNotFound)
}

func TestDeleteMiddleSegment(t *testing.T) {

	createSegmentTest(t, 10, "bytes")
	newSegment2 := createSegmentTest(t, 10, "bytes")
	createSegmentTest(t, 10, "bytes")

	// Test deleteSegment
	req, err := http.NewRequest(http.MethodDelete, baseURL+"/segments/"+newSegment2.ID.String(), nil)
	assert.Equal(t, err, nil)

	resp, err := http.DefaultClient.Do(req)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.StatusCode, http.StatusNoContent)
}

func TestListSegment(t *testing.T) {
	// Test listSegments
	resp, err := http.Get(baseURL + "/segments")
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.StatusCode, http.StatusOK)

	body, err := ioutil.ReadAll(resp.Body)
	assert.Equal(t, err, nil)

	var segmentList []Segment
	err = json.Unmarshal(body, &segmentList)
	assert.Equal(t, err, nil)

	previousSegmentListSize := len(segmentList)

	// create a segment
	createSegmentTest(t, 10, "bytes")

	//get segment list
	resp, err = http.Get(baseURL + "/segments")
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.StatusCode, http.StatusOK)

	body, err = ioutil.ReadAll(resp.Body)
	assert.Equal(t, err, nil)

	err = json.Unmarshal(body, &segmentList)
	assert.Equal(t, err, nil)

	newSegmentListSize := len(segmentList)

	assert.Equal(t, previousSegmentListSize+1, newSegmentListSize)
}

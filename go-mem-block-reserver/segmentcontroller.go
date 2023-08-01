package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func createSegment(c *gin.Context, repo *SegmentRepository) {
	var segment Segment
	if err := c.ShouldBindJSON(&segment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newSegment, err := NewSegment(segment.Size, segment.Unit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	repo.Create(*newSegment)
	c.JSON(http.StatusCreated, newSegment)
}

func getSegment(c *gin.Context, repo *SegmentRepository) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid segment ID"})
		return
	}

	segment, err := repo.Get(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Segment not found"})
		return
	}

	c.JSON(http.StatusOK, segment)
}

func updateSegment(c *gin.Context, repo *SegmentRepository) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid segment ID"})
	}

	var segment Segment
	if err := c.ShouldBindJSON(&segment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	segment.ID = id
	err = repo.Update(segment)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Segment not found"})
		return
	}

	c.JSON(http.StatusOK, segment)
}

func deleteSegment(c *gin.Context, repo *SegmentRepository) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid segment ID"})
		return
	}

	err = repo.Delete(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Segment not found"})
		return
	}

	c.Status(http.StatusNoContent)
}

func listSegments(c *gin.Context, repo *SegmentRepository) {
	segments := repo.List()
	c.JSON(http.StatusOK, segments)
}

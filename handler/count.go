package handler

import (
	"firebase.google.com/go/v4/db"
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
	"ynufesCounterBackend/pkg/firebase"
)

const (
	FBCountPath = "stat/count"
	trialTimes  = 5
)

type CountHandler struct {
	countRef *db.Ref
}

func NewCountHandler(db firebase.Firebase) *CountHandler {
	return &CountHandler{
		countRef: db.Client(FBCountPath),
	}
}

func (h CountHandler) retryFunc(process func() error) (times int, err error) {
	for trialTimes > times {
		times++
		if err = process(); err == nil {
			return times, nil
		}
		// sleep for random time
		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
	}
	return times, err
}

func (h CountHandler) HandleEntry(c *gin.Context) {
	times, err := h.retryFunc(func() error {
		return h.countRef.Transaction(c, func(tx db.TransactionNode) (interface{}, error) {
			var count int
			if err := tx.Unmarshal(&count); err != nil {
				return nil, err
			}
			count++
			return count, nil
		})
	})
	if err != nil {
		c.JSON(500, gin.H{
			"message": "failed",
			"trial":   times,
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
		"trial":   times,
	})
}

func (h CountHandler) HandleExit(c *gin.Context) {
	times, err := h.retryFunc(func() error {
		return h.countRef.Transaction(c, func(tx db.TransactionNode) (interface{}, error) {
			var count int
			if err := tx.Unmarshal(&count); err != nil {
				return nil, err
			}
			count--
			return count, nil
		})
	})
	if err != nil {
		c.JSON(500, gin.H{
			"message": "failed",
			"trial":   times,
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
		"trial":   times,
	})
}

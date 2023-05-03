package v1

import (
	"firebase.google.com/go/v4/db"
	"github.com/gin-gonic/gin"
	"ynufesCounterBackend/pkg/firebase"
	"ynufesCounterBackend/util"
)

const (
	FBCountPath   = "stat/count"
	trialTimes    = 5
	trialInterval = 1000
)

type CountHandler struct {
	countRef *db.Ref
}

func NewCountHandler(db firebase.Firebase) *CountHandler {
	return &CountHandler{
		countRef: db.Client(FBCountPath),
	}
}

func (h CountHandler) HandleEntry(c *gin.Context) {
	times, err := util.RetryFunc(trialTimes, trialInterval, func() error {
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
	times, err := util.RetryFunc(trialTimes, trialInterval, func() error {
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

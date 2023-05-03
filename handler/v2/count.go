package v2

import (
	"firebase.google.com/go/v4/db"
	"github.com/gin-gonic/gin"
	"time"
	"ynufesCounterBackend/domain"
	"ynufesCounterBackend/pkg/firebase"
	"ynufesCounterBackend/util"
)

const (
	FBCountPath   = "stat/count"
	FBRecordPath  = "stat/record"
	FBTimePath    = "stat/time"
	trialTimes    = 5
	trialInterval = 1000
)

type CountHandler struct {
	countRef  *db.Ref
	recordRef *db.Ref
	timeRef   *db.Ref
}

func NewCountHandler(db firebase.Firebase) *CountHandler {
	return &CountHandler{
		countRef:  db.Client(FBCountPath),
		recordRef: db.Client(FBRecordPath),
		timeRef:   db.Client(FBTimePath),
	}
}

func (h CountHandler) HandleEntry(c *gin.Context) {
	WithAuth(h.EntryFunc, c)
}

func (h CountHandler) EntryFunc(id domain.UserID, c *gin.Context) {
	gate, err := domain.NewGate(c.GetInt("gate"))
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	recordTime := time.Now()
	record := domain.ClickRecord{
		UserID: id,
		Gate:   gate,
		Time:   recordTime.UnixMilli(),
	}
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
	if _, err := h.recordRef.Push(c, record); err != nil {
		c.JSON(500, gin.H{
			"message": "failed to update total record",
			"trial":   times,
		})
		return
	}
	timeLabel := recordTime.Format("200601021504")
	tTries, err := util.RetryFunc(trialTimes, trialInterval, func() error {
		return h.timeRef.Child(timeLabel).Transaction(c, func(tx db.TransactionNode) (interface{}, error) {
			var t int
			if err := tx.Unmarshal(&t); err != nil {
				return nil, err
			}
			t++
			return t, nil
		})
	})
	if err != nil {
		c.JSON(500, gin.H{
			"message":   "failed to update time record",
			"statTrial": tTries,
			"timeTrial": times,
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
		"trial":   times,
	})
}

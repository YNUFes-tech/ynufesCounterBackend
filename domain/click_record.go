package domain

type ClickRecord struct {
	UserID int   `json:"id"`
	Gate   Gate  `json:"gate"`
	Time   int64 `json:"time"`
}

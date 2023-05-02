package domain

type ClickRecord struct {
	UserID UserID `json:"id"`
	Gate   Gate   `json:"gate"`
	Time   int64  `json:"time"`
}

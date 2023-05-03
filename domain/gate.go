package domain

import "errors"

type Gate int

const (
	Main        Gate = 1
	North       Gate = 2
	West        Gate = 3
	South       Gate = 4
	SouthCommon Gate = 5
	EastCommon  Gate = 6
)

var ErrInvalidGate = errors.New("invalid gate")

func NewGate(i int) (Gate, error) {
	// validate
	if i < 1 || 6 < i {
		return 0, ErrInvalidGate
	}
	return Gate(i), nil
}

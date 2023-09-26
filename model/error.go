package model

import (
	"go/types"
)

type ErrNotFound struct {
	error types.Interface
}
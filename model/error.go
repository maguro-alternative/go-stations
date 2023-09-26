package model

import (
	"go/types"
)

type ErrNotFound struct {
	Error *types.Interface
}
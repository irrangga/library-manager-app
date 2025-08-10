package entity

import "backend/pkg/constant"

type URL struct {
	Operation    constant.URLOperation
	InitialURL   string
	ProcessedURL string
}

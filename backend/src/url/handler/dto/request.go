package dto

import "backend/pkg/constant"

type URLRequest struct {
	URL       string                `json:"url" binding:"required"`
	Operation constant.URLOperation `json:"operation" binding:"required"`
}

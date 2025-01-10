package utils

import "your_project/entity"

// NewResultEntity 快速构造 ResultEntity
func NewResultEntity(message string, success bool, data interface{}) *entity.ResultEntity {
	return &entity.ResultEntity{
		Message: message,
		Success: success,
		Data:    data,
	}
}

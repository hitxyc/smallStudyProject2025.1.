package entity

// ResultEntity 定义通用的 API 响应结构体
type ResultEntity struct {
	Message string      `json:"message"` // 响应消息
	Success bool        `json:"success"` // 是否成功
	Data    interface{} `json:"data"`    // 数据内容
}

// Construct 构造一个 ResultEntity 对象
func (re *ResultEntity) Construct(message string, success bool, data interface{}) *ResultEntity {
	return &ResultEntity{
		Message: message,
		Success: success,
		Data:    data,
	}
}

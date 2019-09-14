package serializer

import "DuckyGo/model"

// Base 主要序列化器
type Base struct {
	Base string `json:"data"`
}

// BaseOne 单个数据序列化器
type BaseOne struct {
	Result Base `json:"result"`
}

// BaseAll 多个数据序列化器
type BaseAll struct {
	Results []Base `json:"results"`
	Count   int    `json:"count"`
}

// BaseResponse 主要序列化响应
func BaseResponse(db model.Comment) Base {
	return Base{}
}

// BaseOneResponse 单个数据序列化响应
func BaseOneResponse(db model.Comment) BaseOne {
	return BaseOne{}
}

// BaseAllResponse 多个数据序列化响应
func BaseAllResponse(db []model.Comment, count int) BaseAll {
	var result []Base
	for _, i := range db {
		result = append(result, BaseResponse(i))
	}
	return BaseAll{}
}

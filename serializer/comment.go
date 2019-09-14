package serializer

import "DuckyGo/model"

// Comment 用户序列化器
type Comment struct {
	ID       uint   `json:"id"`
	URL      string `json:"Url"`
	name     string `json:"name"`
	location string `json:"location"`
	Rate     uint   `json:"rate"`
	title    string `json:"title"`
	keyword  string `json:"keyword"`
	cid      uint   `json:"c_id"`
	vid      uint   `json:"v_id"`
}

// UserResponse 单个用户序列化
type CommentResponse struct {
	Data Comment `json:"user"`
}

func BuildComment(c model.Comment) Comment {
	return Comment{
		ID:   c.ID,
		URL:  c.URL,
		Rate: c.Rate,
	}
}

// BuildUserResponse 序列化用户响应
func BuildCommentResponse(comment model.Comment) CommentResponse {
	return CommentResponse{
		Data: BuildComment(comment),
	}
}

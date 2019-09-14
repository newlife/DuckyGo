package api

import (
	"DuckyGo/model"
	"DuckyGo/serializer"
	//"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	//"gopkg.in/go-playground/validator.v8"
)

// Ping 状态检查页面
func Ping(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Msg: "Pong",
	}.Result())
}

// GetComment for test
func GetComment(c *gin.Context) {
	var comment model.Comment
	fmt.Println(model.DB)
	model.DB.First(&comment)
	fmt.Println(comment.ID)
	res := serializer.Response{Data: serializer.BuildCommentResponse(comment)}
	c.JSON(200, res.Result())
}

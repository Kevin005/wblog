package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/wangsongyan/wblog/models"
	"net/http"
	"strconv"
)

func CommentPost(c *gin.Context) {
	s := sessions.Default(c)
	sessionUserID := s.Get("UserID")
	postId := c.PostForm("postId")
	pid, _ := strconv.ParseUint(postId, 10, 64)
	if sessionUserID != nil {
		userId, _ := sessionUserID.(uint)
		content := c.PostForm("content")

		comment := new(models.Comment)
		comment.UserID = userId
		comment.Content = content
		comment.PostID = uint(pid)
		comment.Insert()

	}
	c.Redirect(http.StatusMovedPermanently, "/post/"+postId)
}

func CommentDelete(c *gin.Context) {
	s := sessions.Default(c)
	sessionUserID := s.Get("UserID")
	commentId := c.Param("id")
	cid, _ := strconv.ParseUint(commentId, 10, 64)
	var err error
	if sessionUserID != nil {
		var comment *models.Comment
		comment, err = models.GetComment(commentId)
		if err == nil && comment.ID == uint(cid) {
			err = comment.Delete()
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"succeed": err == nil,
	})
}
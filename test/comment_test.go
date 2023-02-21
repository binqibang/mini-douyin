package test

import (
	"fmt"
	"github.com/binqibang/mini-douyin/business"
	"github.com/binqibang/mini-douyin/controller"
	"testing"
)

func TestGetCommentsByVideoId(t *testing.T) {
	fmt.Println(business.GetCommentsByVideoId("5"))
}

func TestConvComment(t *testing.T) {
	comments, err := business.GetCommentsByVideoId("5")
	if err != nil {
		t.Error("Query error")
		t.FailNow()
	}
	fmt.Println(controller.ConvComment(comments))
}

package handler

import (
	proto "article_service/proto"
	"article_service/db"
	"context"
	//"log"
)

type ArticleService struct {}

func (art *ArticleService) ArticleList(ctx context.Context, req *proto.ArticleListRequest, rsp *proto.ArticleListResponse) error {
	rsp.Article, _ = db.GetArticleList()

	return nil
}
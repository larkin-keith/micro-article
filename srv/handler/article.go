package handler

import (
	proto "article_service/proto"
	"article_service/db"
	"context"
	//"log"
)

type ArticleService struct {}

func (art *ArticleService) ArticleList(ctx context.Context, req *proto.ArticleListRequest, rsp *proto.ArticleListResponse) error {

	//for _, v := range articles {
	//	result := []*proto.Article{Id: v.ID}
	//}
	rsp.Article, _ = db.GetArticleList()
	//rsp.Article = []*proto.Article{
	//	{
	//		Id: 1,
	//		Title: "test",
	//		Description: "xxxxx",
	//		CreatAt: "2018-03-01 12:00:00",
	//	},
	//}
	//rsp.Article = result
	return nil
}
package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/micro/go-micro"
	proto "article_service/proto"
	"github.com/micro/go-micro/errors"
	api "github.com/micro/micro/api/proto"
	//"greeter.srv/proto"
	"fmt"
)

type Article struct {
	Client proto.ArticleServiceClient
}

func (article *Article) List(ctx context.Context, req *api.Request, rsp *api.Response) error {
	//log.Print("Received Say.Hello API request")

	page, ok := req.Get["page"]
	pagesize, ok := req.Get["pagesize"]
	if !ok || len(page.Values) == 0 || len(pagesize.Values) == 0 {
		return errors.BadRequest("blog.com.api.article", "Page and Pagesize cannot be blank")
	}

	response, err := article.Client.ArticleList(ctx, &proto.ArticleListRequest{
		Page: 0,
		Pagesize: 10,
	})
	fmt.Println(response)
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(response)

	rsp.Body = string(b)

	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("blog.com.api.article"),
	)

	// parse command line flags
	service.Init()

	service.Server().Handle(
		service.Server().NewHandler(
			&Article{Client: proto.NewArticleServiceClient("blog.com.srv.aritcle", service.Client())},
		),
	)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
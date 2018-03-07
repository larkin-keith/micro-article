package main

import (
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/cli"
	proto "article_service/proto"
	"article_service/srv/handler"
	"article_service/db"
	"context"
	"os"
	//"github.com/micro/db-srv/db"
)

// Setup and the client
func runClient(service micro.Service) {
	// Create new greeter client
	article := proto.NewArticleServiceClient("blog.com.srv.article", service.Client())

	// Call the greeter
	rsp, err := article.ArticleList(context.TODO(), &proto.ArticleListRequest{Page: 1, Pagesize:10})
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print response
	fmt.Println(rsp.Article)
}

func main() {
	//db.GetArticleList()
	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("blog.com.srv.article"),
		micro.Version("latest"),
		micro.Flags(cli.BoolFlag{
			Name:  "run_client",
			Usage: "Launch the client",
		}),
	)

	service.Init(
		micro.Action(func(c *cli.Context) {
			db.Init()
			//db.GetArticleList()
			if c.Bool("run_client") {
				runClient(service)
				os.Exit(0)
			}
		}))

	// Register handler
	proto.RegisterArticleServiceHandler(service.Server(), new(handler.ArticleService))

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
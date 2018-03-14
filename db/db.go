package db

import (
	//"log"
	"time"
	//"fmt"
	
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	proto "article_service/proto"
)

var (
	DB *gorm.DB
	//err error
	//article []Article
)

type Article struct {
	ID        	int    `gorm:"primary_key"`
	Title     	string `gorm:"type:varchar(128);not null;index:title_idx"`
	Description string
	CreatedAt 	time.Time
}

/*
 * Init sql
 */
func Init() (*gorm.DB, error){

	db, err := gorm.Open("mysql", "root:root@/article?charset=utf8&parseTime=True&loc=Local")
	if err == nil {
		DB = db
		db.AutoMigrate(&Article{})
		return db, nil
	}

	return nil, err
}

/*
 * Create articles table from gorm migrate
 *
 * return error
 */
//func migrate() {
//	if ! db.HasTable(&Article{}) {
//		if err := db.AutoMigrate(&Article{}); err != nil {
//			log.Fatal(err)
//		}
//		log.Println("Create table success !")
//	}
//}

//func GetArticleInfo(id int32) *proto.Article {
//	pArticle := &proto.Article{}
//	//r := db.First(&article, id)
//	//pArticle = r
//	//var article []*proto.Article
//
//	return pArticle
//
//}

func GetArticleList() ([]*proto.Article, error) {
	var article []*proto.Article
	DB.Find(&article)
	return article, nil
}




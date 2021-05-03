package repository

import (
	"github.com/zaenalarifin12/golang_article/entity"
	"gorm.io/gorm"
)

type ArticleRepository interface {
	AllArticle() []entity.Article
	Insert(article entity.Article) entity.Article
	Update(article entity.Article) entity.Article
	Show(articleID uint64) entity.Article
	Delete(article entity.Article)
}

type articleConnection struct {
	connection *gorm.DB
}

func NewArticleRepository(db *gorm.DB) *articleConnection {
	return &articleConnection{
		connection: db,
	}
}

func (db *articleConnection) AllArticle() []entity.Article {
	var allArticle []entity.Article
	db.connection.Find(&allArticle)
	return allArticle
}

func (db *articleConnection) Insert(article entity.Article) entity.Article  {
	db.connection.Save(&article)
	db.connection.Preload("User").First(&article)
	return article
}

func (db *articleConnection) Update(article entity.Article) entity.Article  {
	db.connection.Updates(&article)
	db.connection.Preload("User").First(&article)
	return article
}

func (db *articleConnection) Show(articleID uint64) entity.Article {
	var article entity.Article
	db.connection.Preload("User").Find(&article, articleID)
	return article
}

func (db *articleConnection) Delete(article entity.Article)  {
	db.connection.Delete(&article)
}

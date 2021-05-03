package service

import (
	"fmt"
	"github.com/mashingan/smapping"
	"github.com/zaenalarifin12/golang_article/dto"
	"github.com/zaenalarifin12/golang_article/entity"
	"github.com/zaenalarifin12/golang_article/repository"
	"log"
)

type ArticleService interface {
	GetAll() []entity.Article
	Insert(dto dto.ArticleCreateDTO) entity.Article
	Update(updateDTO dto.ArticleUpdateDTO) entity.Article
	FindById(articleID uint64) entity.Article
	Delete(article entity.Article)
	IsAllowedToEdit(userID string, articleID uint64) bool
}

type articleService struct {
	articleRepository repository.ArticleRepository
}

func NewArticleService(repository repository.ArticleRepository) *articleService {
	return &articleService{
		articleRepository: repository,
	}
}

func (services *articleService) GetAll() []entity.Article {
	return services.articleRepository.AllArticle()
}

func (services *articleService) Insert(createDTO dto.ArticleCreateDTO) entity.Article {
	var article entity.Article
	err := smapping.FillStruct(&article, smapping.MapFields(&createDTO))
	if err != nil {
		log.Fatalf("failed to map %v", err)
	}
	res := services.articleRepository.Insert(article)
	return res
}

func (services *articleService) Update(updateDTO dto.ArticleUpdateDTO) entity.Article {
	var article entity.Article
	err := smapping.FillStruct(&article, smapping.MapFields(&updateDTO))
	if err != nil {
		log.Fatalf("Failed to map %v", err)
	}

	res := services.articleRepository.Update(article)
	return res
}

func (services *articleService) FindById(articleID uint64) entity.Article {
	return services.articleRepository.Show(articleID)
}

func (services *articleService) Delete(article entity.Article)  {
	services.articleRepository.Delete(article)
}

func (services *articleService) IsAllowedToEdit(userID string, articleID uint64) bool {
	article := services.articleRepository.Show(articleID)
	id := fmt.Sprintf("%v", article.UserID)
	return userID == id
}

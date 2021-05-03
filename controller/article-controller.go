package controller

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"github.com/zaenalarifin12/golang_article/dto"
	"github.com/zaenalarifin12/golang_article/entity"
	"github.com/zaenalarifin12/golang_article/helper"
	"github.com/zaenalarifin12/golang_article/service"
	"net/http"
	"strconv"
	"strings"
)

type ArticleController interface {
	GetAll(ctx *gin.Context)
	Insert(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	FindById(ctx *gin.Context)
}

type articleController struct {
	articleService service.ArticleService
	jwtService     service.JWTService
}

func NewArticleController(articleService service.ArticleService, jwtService service.JWTService) *articleController {
	return &articleController{
		articleService: articleService,
		jwtService:     jwtService,
	}
}

func (controller *articleController) GetAll(ctx *gin.Context) {
	var articles []entity.Article = controller.articleService.GetAll()
	ctx.JSON(http.StatusOK, articles)
}

func (controller *articleController) Insert(ctx *gin.Context) {
	var articleDTO dto.ArticleCreateDTO

	errDTO := ctx.ShouldBind(&articleDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	userID := controller.getUserByToken(authHeader)
	convertedUserID, err := strconv.ParseUint(userID, 10, 64)
	if err == nil {
		articleDTO.Slug = slug.Make(articleDTO.Title)
		articleDTO.UserID = convertedUserID
	}
	result := controller.articleService.Insert(articleDTO)
	response := helper.BuildResponse(true, "OK", result)
	ctx.JSON(http.StatusOK, response)
}

func (controller *articleController) Update(ctx *gin.Context) {
	var articleUpdate dto.ArticleUpdateDTO
	articleUpdate.ID, _ = strconv.ParseUint(ctx.Param("id"), 10, 64)

	article := controller.articleService.FindById(articleUpdate.ID)
	// check not exist article
	if (article == entity.Article{}) {
		res := helper.BuildErrorResponse("Article not found", "Article not found", helper.EmptyObj{})
		ctx.JSON(http.StatusNotFound, res)
		return
	}

	errDTO := ctx.ShouldBind(&articleUpdate)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	userID := controller.getUserByToken(authHeader)

	if controller.articleService.IsAllowedToEdit(userID, article.UserID) {

		articleUpdate.UserID = article.UserID
		// if image null
		if articleUpdate.Image == "" {
			articleUpdate.Image = article.Image
		}

		result := controller.articleService.Update(articleUpdate)
		res := helper.BuildResponse(true, "article updated", result)
		ctx.JSON(http.StatusOK, res)

		return
	} else {
		res := helper.BuildErrorResponse("you have not permission", "you are not owner", helper.EmptyObj{})
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}

}

func (controller *articleController) Delete(ctx *gin.Context) {
	var id, _ = strconv.ParseUint(ctx.Param("id"), 10, 64)

	article := controller.articleService.FindById(id)
	if (article == entity.Article{}) {
		res := helper.BuildErrorResponse("Article not found", "Article not found", helper.EmptyObj{})
		ctx.JSON(http.StatusNotFound, res)
		return
	}

	controller.articleService.Delete(article)
	res := helper.BuildResponse(true, "OK", 1)
	ctx.JSON(http.StatusOK, res)
}
func (controller *articleController) FindById(ctx *gin.Context) {

	var id, _ = strconv.ParseUint(ctx.Param("id"), 10, 64)

	article := controller.articleService.FindById(id)
	if (article == entity.Article{}) {
		res := helper.BuildErrorResponse("Article not found", "Article not found", helper.EmptyObj{})
		ctx.JSON(http.StatusNotFound, res)
		return
	}
	res := helper.BuildResponse(true, "OK", article)
	ctx.JSON(http.StatusOK, res)

}

/***
get id user from token
*/
func (controller *articleController) getUserByToken(header string) string {
	splitAuthHeader := strings.Fields(header)
	token := strings.Join(splitAuthHeader[1:], "")
	aToken, err := controller.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}

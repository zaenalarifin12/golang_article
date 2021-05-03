package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"net/http"
	"time"
)

type ImageController interface {
	UploadImage(ctx *gin.Context)
}

type imageController struct {
}

func NewImageController() *imageController {
	return &imageController{}
}

func (controller *imageController) UploadImage(ctx *gin.Context) {

	//	input type file
	file, err := ctx.FormFile("image")


	if err != nil {
		ctx.String(http.StatusBadRequest, fmt.Sprintf("err :%", err.Error()))
		return
	}

	// set folder for save file

	path := "public/images/" + slug.Make(time.Now().String()) + file.Filename
	if err := ctx.SaveUploadedFile(file, path); err != nil {
		ctx.String(http.StatusBadRequest, fmt.Sprintf("err: %s", err.Error()))
		return
	}

	// response
	ctx.JSON(http.StatusOK, gin.H{
		"data": path,
	})

}

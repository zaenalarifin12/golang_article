package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zaenalarifin12/golang_article/dto"
	"github.com/zaenalarifin12/golang_article/entity"
	"github.com/zaenalarifin12/golang_article/helper"
	"github.com/zaenalarifin12/golang_article/service"
	"net/http"
	"strconv"
)

type AuthController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type authController struct {
	authService service.AuthService
	jwtService service.JWTService
}

func NewAuthController(authService service.AuthService, jwtService service.JWTService) *authController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (controller *authController) Register(ctx *gin.Context) {
	var registerDTO dto.AuthRegisterDTO
	errDTO := ctx.ShouldBind(&registerDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if !controller.authService.IsDuplicateEmail(registerDTO.Email){
		response := helper.BuildErrorResponse("Failed to process request", "email already exists", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}else{
		createdUser := controller.authService.Register(registerDTO)
		token := controller.jwtService.GenerateToken(strconv.FormatUint(createdUser.ID, 10))
		createdUser.Token = token
		response := helper.BuildResponse(true, "OK", createdUser)
		ctx.JSON(http.StatusOK, response)
		return
	}
}

func (controller *authController) Login(ctx *gin.Context)  {
	var loginDTO dto.AuthLoginDTO
	errDTO := ctx.ShouldBind(&loginDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	authResult := controller.authService.VerifyCredential(loginDTO.Email, loginDTO.Password)
	if v, ok := authResult.(entity.User); ok {
		generatedToken := controller.jwtService.GenerateToken(strconv.FormatUint(v.ID, 10))
		v.Token = generatedToken
		res := helper.BuildResponse(true, "success", v)
		ctx.JSON(http.StatusOK, res)
		return
	}

	res := helper.BuildErrorResponse("check again your email and password", "Invalid credential", helper.EmptyObj{})
	ctx.JSON(http.StatusUnauthorized, res)


//	jika ada maka akan saya cocokan
// apakah email dan password sama
//jika ya maka sukses
//jika tidak ada maka gagal

}

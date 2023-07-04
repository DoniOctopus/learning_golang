package handler

import (
	"goclean/model"
	"goclean/usecase"
	"goclean/utils/authutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	userUc usecase.UserUsecase
}

func (l LoginHandler) login(ctx *gin.Context) {
	loginUserName := &model.LoginModel{}
	err := ctx.ShouldBindJSON(&loginUserName)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}
	usr, errGetName := l.userUc.GetUserByName(loginUserName.Username)
	if usr == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Name is Invalid",
		})
		return
	}
	if errGetName != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": err.Error(),
		})
		return
	}
	temp, err := authutil.GenerateToken(loginUserName.Username)
	if err != nil {
		log.Println("Token Invalid")
	}
	ctx.JSON(http.StatusOK, gin.H{
		"token": temp,
	})
}

func NewLoginHandler(lg *gin.Engine, loginusecase usecase.UserUsecase) {
	LoginHandler := &LoginHandler{
		userUc: loginusecase,
	}
	lg.POST("/login", LoginHandler.login)
}

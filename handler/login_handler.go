package handler

import (
	"goclean/usecase"
	"goclean/utils/authutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	loginUsecase usecase.LoginUsecase
}

func (loginHandler LoginHandler) login(lg *gin.Context){
	temp,err := authutil.GenerateToken("doni")
	if err != nil {
		lg.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})	
	}
	lg.JSON(http.StatusOK, gin.H{
        "token": temp,
    })
}



func NewLoginHandler(lg *gin.Engine, loginusecase usecase.LoginUsecase){
	LoginHandler := &LoginHandler{
		loginUsecase: loginusecase ,
	}
		lg.POST("/login", LoginHandler.login)
}
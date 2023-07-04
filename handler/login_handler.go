package handler

import (
	"goclean/usecase"
	"goclean/utils/authutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	loginUsecase usecase.LoginUsecase
}

func (loginHandler LoginHandler) login(lg *gin.Context){
	temp,err := authutil.GenerateToken("doni")
	if err != nil {
		log.Println("Token Invalid")
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

func LoginUser(){
	
}
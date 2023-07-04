package handler

import (
	"errors"
	"fmt"
	"goclean/apperror"
	"goclean/model"
	"goclean/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	usrUsecase usecase.UserUsecase
}

func (usrHandler UserHandler) InsertUser(ctx *gin.Context) {
	usr := &model.UserModel{}
	err := ctx.ShouldBindJSON(&usr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}

	if len(usr.Name) > 15 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Panjang Nama tidak boleh lebih dari 15 karakter",
		})
		return
	}

	err = usrHandler.usrUsecase.InsertUser(usr)
	if err != nil {
		appError := apperror.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("UserHandler.InsertUser() 1 : %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
		} else {
			fmt.Printf("UserHandler.InsertUser() 2 : %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "Terjadi kesalahan ketika menyimpan data User",
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})

}

func (usrHandler UserHandler) GetUserByName(ctx *gin.Context) {
	name := ctx.Param("name")
	if name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Nama tidak boleh kosong",
		})
		return
	}

	usr, err := usrHandler.usrUsecase.GetUserByName(name)
	if err != nil {
		fmt.Printf("UserHandler.GetUserByName() : %v ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Terjadi kesalahan ketika mengambil data User",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    usr,
	})
}


func NewUserHandler(srv *gin.Engine, usrUsecase usecase.UserUsecase) *UserHandler {
	usrHandler := &UserHandler{
		usrUsecase: usrUsecase,
	}
	srv.POST("/User", usrHandler.InsertUser)
	srv.GET("/User/:name", usrHandler.GetUserByName)
	return usrHandler
}

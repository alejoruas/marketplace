package controller

import (
	"encoding/json"
	"fmt"
	"marketplace/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateProjectAction struct {
	usecase usecases.CreateProjectUseCase
	//TODO: add logger and validator
}

func NewCreateProjectAction(usecase usecases.CreateProjectUseCase) CreateProjectAction {
	return CreateProjectAction{usecase: usecase}
}

func (cpa CreateProjectAction) Execute(ctx *gin.Context) {

	var input usecases.ProjectInput

	fmt.Println(ctx.Request.Body)

	err1 := json.NewDecoder(ctx.Request.Body).Decode(&input)

	if err1 != nil {
		ctx.JSON(http.StatusBadRequest, err1)
	}

	//TODO: validate input

	output, err2 := cpa.usecase.Execute(ctx, input)

	if err2 != nil {
		ctx.JSON(http.StatusInternalServerError, err2)
	}

	fmt.Println(output)

	ctx.JSON(http.StatusOK, output)
}

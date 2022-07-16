package approuter

import (
	controller "marketplace/adapter/controllers"
	presenter "marketplace/adapter/presenters"
	"marketplace/adapter/repository"
	"marketplace/usecases"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func StartRouter(ginEngine *gin.Engine, sqldb repository.SQL) {
	ginEngine.POST("/projects", buildCreateProjectAction(ginEngine, sqldb))
	ginEngine.GET("/hello", helloWorld())
}

func buildCreateProjectAction(g *gin.Engine, sqldb repository.SQL) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			usecase = usecases.NewProjectInteractor(
				repository.NewProjectSql(sqldb),
				repository.NewImputationUnitySql(sqldb),
				presenter.NewCreateProjectPresenter(),
				10*time.Second)
			action = controller.NewCreateProjectAction(usecase)
		)

		action.Execute(ctx)
	}
}

func helloWorld() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "holaaa -- PROBANDO CAMBIOS")
	}
}

package approuter

import (
	controller "marketplace/adapter/controllers"
	presenter "marketplace/adapter/presenters"
	"marketplace/usecases"
	"time"

	"github.com/gin-gonic/gin"
)

func StartRouter(ginEngine *gin.Engine) {
	ginEngine.POST("/projects", buildCreateProjectAction(ginEngine))
}

func buildCreateProjectAction(g *gin.Engine) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			usecase = usecases.NewProjectInteractor(nil, presenter.NewCreCreateProjectPresenter(), 10*time.Second)
			action  = controller.NewCreateProjectAction(usecase)
		)

		action.Execute(ctx)
	}
}

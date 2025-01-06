package factory

import (
	"chapter1/config"

	b_hnd "chapter1/internal/features/books/handler"
	b_rep "chapter1/internal/features/books/repository"
	b_srv "chapter1/internal/features/books/service"

	u_hnd "chapter1/internal/features/users/handler"
	u_rep "chapter1/internal/features/users/repository"
	u_srv "chapter1/internal/features/users/service"

	f_hnd "chapter1/internal/features/feedbacks/handler"
	f_rep "chapter1/internal/features/feedbacks/repository"
	f_srv "chapter1/internal/features/feedbacks/service"

	q_hnd "chapter1/internal/features/quotes/handler"
	q_rep "chapter1/internal/features/quotes/repository"
	q_srv "chapter1/internal/features/quotes/service"

	"chapter1/internal/routes"

	"chapter1/internal/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitFactory(e *echo.Echo) {
	db, _ := config.ConnectDB()

	pu := utils.NewPassUtil()
	jwt := utils.NewJwtUtility()
	cloud := utils.NewCloudinaryUtility()

	uq := u_rep.NewUserQuery(db)
	us := u_srv.NewUserServices(uq, pu, jwt, cloud)
	uh := u_hnd.NewUserHandler(us, jwt)

	bq := b_rep.NewBookQuery(db)
	bs := b_srv.NewBookServices(bq, jwt, cloud, us)
	bh := b_hnd.NewBookHandler(bs, jwt)

	fq := f_rep.NewFeedbackQuery(db)
	fs := f_srv.NewFeedbackServices(fq, us)
	fh := f_hnd.NewFeedbackHandler(fs, jwt)

	qq := q_rep.NewQuoteQuery(db)
	qs := q_srv.NewQuoteServices(qq, us)
	qh := q_hnd.NewQuoteHandler(qs, jwt)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	routes.InitRoute(e, uh, bh, fh, qh)
}

package routes

import (
	"chapter1/config"
	"chapter1/internal/features/books"
	"chapter1/internal/features/feedbacks"
	"chapter1/internal/features/quotes"
	"chapter1/internal/features/users"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo, uh users.UHandler, bh books.BHandler, fh feedbacks.FHandler, qh quotes.QHandler) {
	e.POST("/login", uh.Login())
	e.POST("/register", uh.Register())

	e.GET("/books/:id", bh.GetBookByID())
	e.GET("/books", bh.GetAllBooks())

	e.GET("/feedbacks", fh.GetAllFeedbacks())

	e.GET("/quotes", qh.AddQuote())

	MemberRoute(e, uh)
	BookRoute(e, bh)
	FeedbackRoute(e, fh)
	QuoteRoute(e, qh)
}

func MemberRoute(e *echo.Echo, uh users.UHandler) {
	u := e.Group("/users")
	u.Use(JWTConfig())
	u.PUT("", uh.UpdateUser())
	u.DELETE("/:id", uh.DeleteUser())
	u.GET("", uh.GetUserByID())
	u.GET("/admin", uh.GetAllUsers())

}

func BookRoute(e *echo.Echo, bh books.BHandler) {
	b := e.Group("/books")
	b.Use(JWTConfig())
	b.PUT("/:id", bh.UpdateBook())
	b.DELETE("/:id", bh.DeleteBook())
	b.POST("", bh.AddBook())
}

func FeedbackRoute(e *echo.Echo, fh feedbacks.FHandler) {
	f := e.Group("/feedbacks")
	f.Use(JWTConfig())
	f.POST("", fh.AddFeedback())
	f.DELETE("/:id", fh.DeleteFeedback())
}

func QuoteRoute(e *echo.Echo, qh quotes.QHandler) {
	q := e.Group("quotes")
	q.Use(JWTConfig())
	q.POST("", qh.AddQuote())
}

func JWTConfig() echo.MiddlewareFunc {
	return echojwt.WithConfig(
		echojwt.Config{
			SigningKey:    []byte(config.ImportSetting().JWTSecret),
			SigningMethod: jwt.SigningMethodHS256.Name,
		},
	)
}

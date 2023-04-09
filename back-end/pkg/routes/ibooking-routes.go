package routes

import (
	"github.com/Nrdzz/iBooking/pkg/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

var RegisterBookingRoutes = func(router *gin.Engine) {
	// administrator
	//adminRouter := router.Group("/admin")
	//{
	//	adminRouter.POST("/login", controllers.AdminLogin)
	//}

	// room management
	roomRouter := router.Group("/room")
	{
		roomRouter.POST("/", controllers.CreateRoom)
		roomRouter.GET("/", controllers.GetRoom)
		roomRouter.PUT("/", controllers.UpdateRoom)
		roomRouter.DELETE("/:roomId", controllers.DeleteRoom)
		roomRouter.GET("/:roomId", controllers.GetRoomById)
	}

	// seat management
	seatRouter := router.Group("/seat")
	//{
	seatRouter.POST("/", controllers.CreateSeat)
	//	seatRouter.GET("/", controllers.GetSeat)
	//	seatRouter.PUT("/", controllers.UpdateSeat)
	//	seatRouter.DELETE("/:seatId", controllers.DeleteSeat)
	//	seatRouter.GET("/:seatId", controllers.GetSeatById)
	//}

	// user management
	//userRouter := router.Group("/user")
	//{
	//	userRouter.POST("/", controllers.CreateUser)
	//	userRouter.POST("/login", controllers.UserLogin)
	//	userRouter.PUT("/:userId", controllers.UpdateUser)
	//	// userRouter.DELETE("/:userId", controllers.DeleteUser)
	//	userRouter.GET("/:userId", controllers.GetUserById)
	//}

	// booking management
	//bookingRouter := router.Group("/book")
	//{
	//	bookingRouter.POST("/", controllers.BookSeat)
	//	bookingRouter.GET("/", controllers.GetBook)
	//	bookingRouter.PUT("/", controllers.UpdateBook) // update or attend
	//	bookingRouter.DELETE("/:bookingId", controllers.DeleteBook) // cancel
	//}

	// notification management
	//notificationRouter := router.Group("/notification")
	//{
	//	notificationRouter.POST("/", controllers.CreateNotification)
	//}

	// default 404
	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "templates/404.html", nil)
	})
}

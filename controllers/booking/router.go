package booking

import "github.com/labstack/echo"

func Private(g *echo.Group) {
	g.GET("/countBooking", CountBooking)
	g.POST("/booking/create_booking",Booking)
}

package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Booking struct {
	CheckIn  time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`
}

func main() {
	r := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookabledate", bookableDate)
	}

	r.GET("/bookable", func(c *gin.Context) {
		var b Booking
		if err := c.ShouldBind(&b); err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
			"booking": b,
		})
	})

	r.Run()
}

func bookableDate(fl validator.FieldLevel) bool {

	if data, ok := fl.Field().Interface().(time.Time); ok {
		today := time.Now()
		if data.Unix() > today.Unix() {
			return true
		}
	}
	return false
}

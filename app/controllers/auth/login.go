package auth

import (
	"gin-api/app/models"
	"gin-api/app/services"
	"gin-api/db"
	"net/http"

	"github.com/gin-gonic/gin"
)


func Login(c *gin.Context) {
  var u models.User
	var user models.User

  if err := c.ShouldBindJSON(&u); err != nil {
     c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
     return
  }

	db.GetDB().Find(&user, "name = ?", u.Name)

  if user.Name != u.Name || !services.IsHashEqual(u.Password, user.Password) {
     c.JSON(http.StatusUnauthorized, "Please provide valid login details")
     return
  }

  token, err := services.CreateToken(uint64(user.ID))
  if err != nil {
     c.JSON(http.StatusUnprocessableEntity, err.Error())
     return
  }
  c.JSON(http.StatusOK, token)
}

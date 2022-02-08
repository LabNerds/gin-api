package auth

import (
	"gin-api/app/models"
	"gin-api/app/services"
	"gin-api/db"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/jsonapi"
)

func SignUp(c *gin.Context) {
	r := c.Request
	w := c.Writer

	user := new(models.User)

	if err := jsonapi.UnmarshalPayload(r.Body, user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.Password, _ = services.Hash(user.Password)

	db.GetDB().Create(&user)

	w.Header().Set("Content-Type", jsonapi.MediaType)
	w.WriteHeader(http.StatusCreated)

	if err := jsonapi.MarshalPayload(w, user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	return
}

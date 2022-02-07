package services

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(userid uint64) (string, error) {
  var err error
  //Creating Access Token
	// NOTE: this should be in an env file
  os.Setenv("ACCESS_SECRET", "jdnfksdmfksd")

  atClaims := jwt.MapClaims{}
  atClaims["authorized"] = true
  atClaims["user_id"] = userid
  atClaims["exp"] = time.Now().Add(time.Minute * 30).Unix()

  at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
  token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
  if err != nil {
     return "", err
  }
  return token, nil
}

package controllers

import (
	"golang_backend/initializers"
	"golang_backend/models"
	"golang_backend/schemas"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func UserCreate(c * gin.Context) {
	if c.Bind(&schemas.SignupBody) !=nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error": "Failed to read request body",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(schemas.SignupBody.Password),10)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error": "Failed to hash password",
		})
		return		
	}

	user := models.UserModel{
		Email: schemas.SignupBody.Email,
		Password: string(hash),
	}
	result:= initializers.DB.Create(&user)
	if result.Error!=nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error": "Failed to create user",
		})
		return	
	}

c.JSON(http.StatusOK,gin.H{})
}

func UserLogin(c *gin.Context) {
	var user models.UserModel

	if c.Bind(&schemas.SignupBody) !=nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error": "Failed to read request body",
		})
		return
	}

	initializers.DB.First(&user,"email = ?", schemas.SignupBody.Email)
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest,gin.H{
			"error": "Incorrect email",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(schemas.SignupBody.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error": "Incorrect Password",
		})
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Email,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error": "Failed to create token",
		})
		return
	}
	c.SetSameSite(http.SameSiteDefaultMode)
	c.SetCookie("Authorization",tokenString,3600*24*30,"","",false,true)
	c.JSON(http.StatusOK,gin.H{})


}
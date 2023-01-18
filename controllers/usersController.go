package controllers

import (
	"golang_backend/initializers"
	"golang_backend/models"
	"golang_backend/schemas"
	"golang_backend/validators"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/datatypes"
)

func UserCreate(c * gin.Context) {
	roles := []string {0:"Doctor",1:"Patient",3:"TechStaff"}
	if c.Bind(&schemas.SignupBody) !=nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error": "Failed to read request body",
		})
		return
	}

	newDOB,err:= time.Parse(time.RFC3339,schemas.SignupBody.DOB)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error": "bad date format",
		})
		return		
	}

	if !validators.OptionValidator(roles,schemas.SignupBody.Role) {
		c.JSON(http.StatusBadRequest,gin.H{
			"error": "wrong value for role format",
		})
		return	
	}

	isAdminBool,err := strconv.ParseBool(schemas.SignupBody.IsAdmin)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error": "Wrong value for isadmin",
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
		FirstName: schemas.SignupBody.FirstName,
		LastName: schemas.SignupBody.LastName,
		DOB: datatypes.Date(newDOB) ,
		IsAdmin: isAdminBool,
		Role: schemas.SignupBody.Role,
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

	if c.Bind(&schemas.LoginBody) !=nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error": "Failed to read request body",
		})
		return
	}

	initializers.DB.First(&user,"email = ?", schemas.LoginBody.Email)
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest,gin.H{
			"error": "Incorrect email",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(schemas.LoginBody.Password))
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


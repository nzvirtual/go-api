package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	log "github.com/dhawton/log4g"
	"github.com/gin-gonic/gin"
	"github.com/nzvirtual/go-api/lib/database/models"
	"github.com/nzvirtual/go-api/lib/utils"
	"golang.org/x/crypto/bcrypt"
)

type RegisterDTO struct {
	Email     string `json:"email" form:"email" binding:"required"`
	Firstname string `json:"firstname" form:"firstname" binding:"required"`
	Lastname  string `json:"lastname" form:"lastname" binding:"required"`
}

type LoginDTO struct {
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

func PostRegister(c *gin.Context) {
	var j RegisterDTO

	log.Category("api/auth/register").Debug("Unmarshalling request")
	if err := c.ShouldBindJSON(&j); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	log.Category("api/auth/register").Debug("Unmarshalled DTO, generating password")
	temppass := utils.GeneratePassword(10)
	hpass, err := bcrypt.GenerateFromPassword([]byte(temppass), 10)
	if err != nil {
		log.Category("api/auth/register").Error("Error generating random password and hashing " + err.Error())
	}
	log.Category("api/auth/register").Debug("Password generated, got: " + string(hpass))

	var airport models.Airport
	models.DB.Where("icao = ?", "KDEN").First(&airport)

	var rank models.Rank
	models.DB.Where("name = ?", "Second Officer").First(&rank)

	user := models.User{
		Email:             j.Email,
		Firstname:         j.Firstname,
		Lastname:          j.Lastname,
		Password:          string(hpass),
		Verified:          false,
		VerificationToken: "123123",
		LastAirportID:     airport.ID,
		LastAirport:       airport,
		RankID:            rank.ID,
		Rank:              rank,
	}

	log.Category("api/auth/register").Debug("Model built, creating entry")

	if err = models.DB.Create(&user).Error; err != nil {
		log.Category("api/auth/register").Error("Failed to add new user to database " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	log.Category("api/auth/register").Debug("Calling Email Server")

	jsonData := map[string]string{
		"to":              fmt.Sprintf("%s %s <%s>", user.Firstname, user.Lastname, user.Email),
		"name":            fmt.Sprintf("%s %s", user.Firstname, user.Lastname),
		"TempPassword":    temppass,
		"verificationURL": "http://localhost/test",
	}
	jsonValue, _ := json.Marshal(jsonData)
	request, _ := http.NewRequest("POST", os.Getenv("EMAIL_API")+"/send/registration", bytes.NewBuffer(jsonValue))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("APIKey %s", os.Getenv("EMAIL_API_KEY")))
	client := &http.Client{}
	client.Do(request)

	log.Category("api/auth/register").Debug("Created, returning response 201")
	c.JSON(http.StatusCreated, gin.H{"message": "Created"})
}

func PostLogin(c *gin.Context) {
	var data LoginDTO

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	user, err := utils.AuthenticateUser(data.Email, data.Password)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"message": "Forbidden"})
	}

	token, err := utils.CreateJWTToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		log.Category("api/login").Error("Error while creating token: " + err.Error())
		c.Abort()
	}

	c.JSON(http.StatusCreated, gin.H{"message": "OK", "token": token})
}

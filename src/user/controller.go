package user

import (
	"log"
	"net/http"
	"regexp"

	"mydj_server/config"
	authToken "mydj_server/src/authToken"
	"mydj_server/src/entity"
	"mydj_server/src/inputs"
	"mydj_server/src/response"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetIDFromUUID(uuid uuid.UUID) (uint, error) {
	var err error
	var users []entity.User

	db, _ := config.ReturnDB()
	err = db.Where("uuid = ?", uuid).Find(&users).Error
	if len(users) == 0 {
		log.Println("Error", err.Error())
		return 0, err
	}
	if err != nil {
		log.Println("Error", err.Error())
	}
	return users[0].ID, nil
}

func UserLoginController(c *gin.Context) {
	var credential inputs.UserCredentialInput
	if err := c.ShouldBindJSON(&credential); err != nil {
		response.RespondWithError(c, http.StatusBadRequest, err)
		return
	}

	if credential.Email == "" || credential.Password == "" {
		response.RespondWithError(c, http.StatusBadRequest, nil)
		return
	}

	uuid, find, err := UserLoginService(credential.Email, credential.Password)
	if err != nil {
		response.RespondWithError(c, http.StatusBadRequest, err)
		return
	}

	if find {
		token, err := authToken.GenerateToken(uint64(uuid.ID()))
		if err != nil {
			response.RespondWithError(c, http.StatusInternalServerError, err)
			return
		}

		response.RespondWithSuccess(c, "connected", gin.H{"token": token})
		return
	}

	response.RespondWithNotFound(c)
}

func IsValidPassword(password string) bool {
	uppercase := regexp.MustCompile("[A-Z]")
	lowercase := regexp.MustCompile("[a-z]")
	digit := regexp.MustCompile(`\d`)
	specialChar := regexp.MustCompile("[!@$*?.%]")

	return uppercase.MatchString(password) &&
		lowercase.MatchString(password) &&
		digit.MatchString(password) &&
		specialChar.MatchString(password)
}

func RegisterController(c *gin.Context) {
	var err error
	db, _ := config.ReturnDB()
	var input inputs.UserInitInput
	if err = c.ShouldBindJSON(&input); err != nil {
		response.RespondWithError(c, 400, err)
		return
	}

	if !IsValidPassword(input.Password) {
		response.RespondWithInvalid(c, 400, "Bad password")
		return
	}
	reg, err := regexp.MatchString(`^(?P<name>[a-z][\w\-_.]*[^.])(?P<domain>@\w+)(?P<ext>\.[a-z]+(\.[a-z]+)?[^.\W])$`, input.Email)
	if err != nil {
		response.RespondWithError(c, 400, err)
		return
	}
	if !reg {
		response.RespondWithInvalid(c, 400, "Bad email")
		return
	}
	u := entity.User{}
	u.Firstname = input.Firstname
	u.Lastname = input.Lastname
	u.Email = input.Email
	u.Password = input.Password
	u.UUID, _ = uuid.NewRandom()
	err = BeforeSaveService(&u, db)
	if err != nil {
		response.RespondWithError(c, 400, err)
		return
	}
	_, err = SaveUserService(&u, db)
	if err != nil {
		response.RespondWithError(c, 400, err)
		return
	}

	response.RespondWithSuccessCreation(c, "User created", nil)
}

package controllers

import (
	"strconv"
	"time"

	"github.com/badoux/checkmail"
	"github.com/gofiber/fiber/v2"
	"github.com/musukwamoshi/golang-test-api/databaseconfig"
	"github.com/musukwamoshi/golang-test-api/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User models.User
type Session models.Session

func GetUser(sessionid int) (User, error) {
	db := databaseconfig.DB
	query := Session{ID: sessionid}
	found := Session{}
	err := db.First(&found, &query).Error
	if err == gorm.ErrRecordNotFound {
		return User{}, err
	}
	user := User{}
	usrQuery := User{ID: int(found.UserId)}
	err = db.First(&user, &usrQuery).Error
	if err == gorm.ErrRecordNotFound {
		return User{}, err
	}
	return user, nil
}

func Login(c *fiber.Ctx) error {
	type LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	db := databaseconfig.DB
	json := new(LoginRequest)
	if err := c.BodyParser(json); err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
		})
	}

	found := User{}
	query := User{Email: json.Email}
	err := db.First(&found, &query).Error
	if err == gorm.ErrRecordNotFound {
		return c.JSON(fiber.Map{
			"code":    404,
			"message": "User not found",
		})
	}
	if !comparePasswords(found.Hash, []byte(json.Password)) {
		return c.JSON(fiber.Map{
			"code":    401,
			"message": "Invalid Password",
		})
	}
	session := Session{UserId: found.ID}
	db.Create(&session)
	c.Cookie(&fiber.Cookie{
		Name:     "sessionid",
		Expires:  SessionExpires(),
		Value:    strconv.Itoa(session.ID),
		HTTPOnly: true,
	})
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "sucess",
		"data":    session,
	})
}

func Logout(c *fiber.Ctx) error {
	db := databaseconfig.DB
	json := new(Session)
	if err := c.BodyParser(json); err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
		})
	}
	session := Session{}
	query := Session{ID: json.ID}
	err := db.First(&session, &query).Error
	if err == gorm.ErrRecordNotFound {
		return c.JSON(fiber.Map{
			"code":    404,
			"message": "Session not found",
		})
	}
	db.Delete(&session)
	c.ClearCookie("sessionid")
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "sucess",
	})
}

func CreateUser(c *fiber.Ctx) error {
	type CreateUserRequest struct {
		Password string `json:"password"`
		Username string `json:"username"`
		Email    string `json:"email"`
	}

	db := databaseconfig.DB
	json := new(CreateUserRequest)
	if err := c.BodyParser(json); err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
		})
	}
	password := hashAndSalt([]byte(json.Password))
	err := checkmail.ValidateFormat(json.Email)
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid Email Address",
		})
	}
	new := User{
		Email: json.Email,
		Hash:  password,
		Salt:  json.Password,
	}
	found := User{}
	query := User{Email: json.Email}
	err = db.First(&found, &query).Error
	if err != gorm.ErrRecordNotFound {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "User already exists",
		})
	}
	db.Create(&new)
	session := Session{UserEmail: new.Email, UserId: new.ID}
	err = db.Create(&session).Error
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    500,
			"message": "Creation Error",
		})
	}
	c.Cookie(&fiber.Cookie{
		Name:     "sessionid",
		Expires:  time.Now().Add(5 * 24 * time.Hour),
		Value:    strconv.Itoa(session.ID),
		HTTPOnly: true,
	})
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "sucess",
		"data":    session,
	})
}

func GetUserInfo(c *fiber.Ctx) error {
	user := c.Locals("user").(User)
	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	type DeleteUserRequest struct {
		password string
	}
	db := databaseconfig.DB
	json := new(DeleteUserRequest)
	user := c.Locals("user").(User)
	if err := c.BodyParser(json); err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
		})
	}
	if !comparePasswords(user.Hash, []byte(json.password)) {
		return c.JSON(fiber.Map{
			"code":    401,
			"message": "Invalid Password",
		})
	}
	db.Model(&user).Association("Sessions").Delete()
	//db.Model(&user).Association("Products").Delete()
	db.Delete(&user)
	c.ClearCookie("sessionid")
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "sucess",
	})
}

func ChangePassword(c *fiber.Ctx) error {
	type ChangePasswordRequest struct {
		Password    string `json:"password"`
		NewPassword string `json:"newPassword"`
	}
	db := databaseconfig.DB
	user := c.Locals("user").(User)
	json := new(ChangePasswordRequest)
	if err := c.BodyParser(json); err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
		})
	}
	if !comparePasswords(user.Hash, []byte(json.Password)) {
		return c.JSON(fiber.Map{
			"code":    401,
			"message": "Invalid Password",
		})
	}
	user.Hash = hashAndSalt([]byte(json.NewPassword))
	user.Salt = hashAndSalt([]byte(json.NewPassword))
	db.Save(&user)
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "sucess",
	})
}

func hashAndSalt(pwd []byte) string {
	hash, _ := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	return string(hash)
}
func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	return err == nil
}

// Universal date the Session Will Expire
func SessionExpires() time.Time {
	return time.Now().Add(5 * 24 * time.Hour)
}

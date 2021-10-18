package helpers

import (
	"errors"
	"github.com/gin-gonic/gin"
)

func CheckUserType(c *gin.Context, role string) (err error) {
	userType := c.GetString("user_type")
	err = nil

	if userType != role {
		err = errors.New("Unauthorized to access this resource.")
		return err
	}
	return err
}

func MatchUserTypeToUid(c *gin.Context, userID string) (err error) {
	userType := c.GetString("user_type")
	uID := c.GetString("user_id")
	err = nil

	if userType == "USER" && uID != userID {
		err = errors.New("Unauthorized to access this resource.")
		return err
	}
	err = CheckUserType(c, userType)
	return err
}

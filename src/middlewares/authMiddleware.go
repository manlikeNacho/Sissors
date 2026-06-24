package middlewares

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		authorization := c.Request.Header.Get("Authorization")
		if authorization == "" {
			errDetails := errors.New("no Authorization header provided")
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errDetails.Error(),
			})
			c.Abort()
			return
		}

		accessToken := ""
		sp := strings.Split(authorization, " ")
		if len(sp) > 1 {
			accessToken = sp[1]

		}

		if accessToken == "" {
			c.JSON(403, gin.H{
				"error": "Provide access token",
			})
			c.Abort()
			return
		}

		//...Validates user token and assigns role from the token
		// user_Id, user_type, err := service.TokenServ.ValidateAccessToken(accessToken)
		// if err != nil {
		// 	logger.AppLogger.Error("error validating token ", zap.Error(err))
		// 	errCode := err.(*rerrors.Err).Status()
		// 	c.JSON(errCode, gin.H{
		// 		"error": err,
		// 	})

		// 	c.Abort()
		// 	return
		// }

		// c.Set("userID", user_Id)
		// c.Set("userType", user_type)
		c.Next()
	}
}

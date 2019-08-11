package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	funk "github.com/thoas/go-funk"
)

//AuthenticationRequired ...
func AuthenticationRequired(auths ...string) gin.HandlerFunc {
	fmt.Println("************************** Authenticating *******************************")
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")
		fmt.Println("#####################", user)
		if user == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user needs to be signed in to access this service"})
			c.Abort()
			return
		}
		if len(auths) != 0 {
			authType := session.Get("authType")
			if authType == nil || !funk.ContainsString(auths, authType.(string)) {
				c.JSON(http.StatusForbidden, gin.H{"error": "invalid request, restricted endpoint"})
				c.Abort()
				return
			}
		}
		// add session verification here, like checking if the user and authType
		// combination actually exists if necessary. Try adding caching this (redis)
		// since this middleware might be called a lot
		c.Next()
	}
}

package sub

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSimpleString(c *gin.Context) {
	println("🌟")
	fmt.Println("🐮")
	c.IndentedJSON(http.StatusOK, "albums")

}

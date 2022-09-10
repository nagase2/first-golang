package sub

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetSimpleString(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetSimpleString(tt.args.c)
		})
	}
}
func TestNewCompanyHandler(t *testing.T) {

}

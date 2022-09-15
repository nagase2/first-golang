package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type AuthResult struct {
	Allowed       bool   // `json:"allowed"`
	Reason        string //`json:"reason"`
	EventMetadata struct {
		DeviceId string
	}
}

// @Tags        Authentication
// Weあああああああ
// @Summary SoraWebの認証エンドポイントです。
// @Description Webhookで通知された接続認証情報と、本ソフトで管理する認証情報を照合し
// @Description チャネルIDとシグナリングキーが合致し、設定されたデバイスIDに含まれる場合には 認証成功を示す情報を返す。
// @Produce     json
// @Param request body model.AuthRequestJson true "認証リクエスト"
// @Accept      json
// @Success     200 {object} model.AuthResult
// @Router      /sora/auth/webhook [post]
func (c *Controller) SoraWebhookHandler(ctx *gin.Context) {
	fmt.Println("🈲")
}

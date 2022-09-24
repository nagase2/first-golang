package main

import "fmt"

type RequestHandler interface {
	PeerAuthProvider
}
type SoraWebhookHandler struct {
}

func (s *SoraWebhookHandler) PeerAuth() {
	fmt.Println("Yeah")
}

type PeerAuthProvider interface {
	iAuthUserService
	iAuditLogService
	// Create() (AuthUserService, AuditLogService)
	// PeerAuth()
}

type SoraPeerAuthProvider struct {
	PeerAuthProvider
}

// AuthUser implements PeerAuthProvider
func (*SoraPeerAuthProvider) AuthUser() bool {
	//panic("unimplemented")
}

func (s *SoraPeerAuthProvider) Create() {
	fmt.Println("Yeah")

}

type iAuthUserService interface {
	AuthUser() bool
}
type AuthUserYamlService struct {
}

func (a *AuthUserYamlService) AuthUser() {

}

type iAuditLogService interface {
	//WriteLog()
}
type AuditLogFileService struct {
}

func newSoraPeerAuthProvider() PeerAuthProvider {
	return &SoraPeerAuthProvider{
		AuthUserService: &iAuthUserService{},
	}
}

// // factory method
// func getPeerAuthProvider(providerName string) (PeerAuthProvider, error) {
// 	if providerName == "sora" {
// 		return newSoraPeerAuthProvider()
// 	}
// 	if providerName == "sky" {
// 		return newSkywayPeerAuthProvider()
// 	}
// }

func main() {
	fmt.Println("aa")
	controllerMethod()

}
func controllerMethod() {
	//コントローラの中

	// factoryを使ってサービスを作る

	//必要なAuthとLogを作成
	// create

	//Auth実施
	// PeerAuth

	//log実施
	// writeLog

	//結果を返す
	// gin write request

	//peerAuthProvider := SoraPeerAuthProvider
	//peerAuthProvider.PeerAuth()
}

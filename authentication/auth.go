package authentication

import (
	"context"

	auth "github.com/zgiber/playground/grpc-experiment/pkg/authentication"
)

// Persistence provides data storage for the Authentication service
type Persistence interface {
	ClientIDByToken(ctx context.Context, token string) (string, bool, error)
	SetAccessToken(ctx context.Context, clientID, token string) error
	DeleteAccessTokens(ctx context.Context, clientID string) error
}

type service struct {
	p Persistence
}

func (s *service) Authenticate(ctx context.Context, req *auth.AuthenticationRequest) (res *auth.AuthenticationResponse, err error) {
	var clientID string
	var found bool
	clientID, found, err = s.p.ClientIDByToken(ctx, req.AccessToken)
	if !found {
		res.Error = &auth.Error{
			ErrorCode:   "CLIENT_NOT_FOUND",
			Description: "no client for provided token",
		}
		return
	}

	res.ClientID = clientID
	res.Success = true
	return
}

func (s *service) CreateAccessToken(ctx context.Context, req *auth.CreateAccessTokenRequest) (res *auth.CreateAccessTokenResponse, err error) {
	accessToken := generateRandomToken()
	err = s.p.SetAccessToken(ctx, req.ClientID, accessToken)
	if err != nil {
		res.Error = &auth.Error{
			ErrorCode:   "CREATE_TOKEN_FAILED",
			Description: err.Error(),
		}
		return
	}
	res.AccessToken = accessToken
	res.Success = true
	return
}

func (s *service) DeleteAccessTokens(ctx context.Context, req *auth.DeleteAccessTokensRequest) (res *auth.DeleteAccessTokensResponse, err error) {
	err = s.p.DeleteAccessTokens(ctx, req.ClientID)
	if err != nil {
		res.Error = &auth.Error{
			ErrorCode:   "DELETE_TOKEN_FAILED",
			Description: err.Error(),
		}
		return
	}
	res.Success = true
	return
}

func generateRandomToken() string {
	return "TODO!!" // TODO
}

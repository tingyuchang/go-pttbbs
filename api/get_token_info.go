package api

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"
)

const GET_TOKEN_INFO_R = "/token/info"

type GetTokenInfoParams struct {
	Jwt string `json:"token" form:"token" url:"token"`
}

type GetTokenInfoResult struct {
	ClientInfo string      `json:"client_info"`
	UserID     bbs.UUserID `json:"user_id"`
}

func GetTokenInfoWrapper(c *gin.Context) {
	params := &GetTokenInfoParams{}

	LoginRequiredJSON(GetTokenInfo, params, c)
}

func GetTokenInfo(remoteAddr string, uuserID bbs.UUserID, params interface{}) (result interface{}, err error) {
	theParams, ok := params.(*GetTokenInfoParams)
	if !ok {
		return nil, ErrInvalidParams
	}

	userID, clientInfo, err := VerifyJwt(theParams.Jwt)
	if err != nil {
		return nil, err
	}
	if userID != uuserID {
		return nil, ErrInvalidUser
	}

	result = &GetTokenInfoResult{
		ClientInfo: clientInfo,
		UserID:     userID,
	}

	return result, nil
}

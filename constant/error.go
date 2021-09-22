package constant

import "errors"

var (
	ERR_TOKEN_IS_REQUIRED = "Token is not empty"
	ERR_TOKEN_IS_INVALID  = "Token is invalid"
	ERR_GEN_TOKEN         = "Generate token failed"
	ERR_LIST_USER         = "Fetch data failed"
	ERR_INSERT_USER       = "Insert failed"
	ERR_USER_IS_EXIST     = "Email is exist"
	ERR_USER_NOT_FOUND    = errors.New("User not found")
)

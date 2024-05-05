package helpers

import (
	"github.com/phongtran11/go-project/pkg/constants"
	"github.com/phongtran11/go-project/pkg/validations"
)

type TBaseHttpResponse struct {
	Result           any                             `json:"result"`
	Success          bool                            `json:"success"`
	ResultCode       constants.TResultCode           `json:"resultCode"`
	ValidationErrors *[]validations.TValidationError `json:"validationErrors"`
	Error            any                             `json:"error"`
}

func GenerateBaseResponse(result any, success bool) *TBaseHttpResponse {
	return &TBaseHttpResponse{Result: result,
		Success: success,
	}
}

func GenerateBaseResponseWithError(result any, success bool, resultCode constants.TResultCode, err error) *TBaseHttpResponse {
	return &TBaseHttpResponse{
		Result:     result,
		ResultCode: resultCode,
		Success:    success,
		Error:      err.Error(),
	}

}

func GenerateBaseResponseWithAnyError(result any, success bool, err any) *TBaseHttpResponse {
	return &TBaseHttpResponse{
		Result:  result,
		Success: success,
		Error:   err,
	}
}

func GenerateBaseResponseWithValidationError(result any, success bool, resultCode constants.TResultCode, err error) *TBaseHttpResponse {
	return &TBaseHttpResponse{
		Result:           result,
		Success:          success,
		ResultCode:       resultCode,
		ValidationErrors: validations.GetValidationErrors(err),
	}
}

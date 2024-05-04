package helpers

type TBaseHttpResponse struct {
	Result  any  `json:"result"`
	Success bool `json:"success"`
	Error   any  `json:"error"`
}

func GenerateBaseResponse(result any, success bool) *TBaseHttpResponse {
	return &TBaseHttpResponse{Result: result,
		Success: success,
	}
}

func GenerateBaseResponseWithError(result any, success bool, err error) *TBaseHttpResponse {
	return &TBaseHttpResponse{Result: result,
		Success: success,
		Error:   err.Error(),
	}

}

func GenerateBaseResponseWithAnyError(result any, success bool, err any) *TBaseHttpResponse {
	return &TBaseHttpResponse{Result: result,
		Success: success,
		Error:   err,
	}
}

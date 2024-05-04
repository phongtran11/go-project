package helpers

type BaseHttpResponse struct {
	Result  any  `json:"result"`
	Success bool `json:"success"`
	Error   any  `json:"error"`
}

func GenerateBaseResponse(result any, success bool) *BaseHttpResponse {
	return &BaseHttpResponse{Result: result,
		Success: success,
	}
}

func GenerateBaseResponseWithError(result any, success bool, err error) *BaseHttpResponse {
	return &BaseHttpResponse{Result: result,
		Success: success,
		Error:   err.Error(),
	}

}

func GenerateBaseResponseWithAnyError(result any, success bool, err any) *BaseHttpResponse {
	return &BaseHttpResponse{Result: result,
		Success: success,
		Error:   err,
	}
}

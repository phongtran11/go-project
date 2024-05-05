package constants

const (
	// User
	EmailExists      = "user.001"
	UserNotFound     = "user.002"
	PasswordNotMatch = "user.003"

	// JWT
	GenerateTokenFailed = "jwt.001"
)

type TErrorMap struct {
	Message string `json:"message"`
}

func (e *TErrorMap) Error() string {
	return e.Message
}

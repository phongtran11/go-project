package constants

type TResultCode int

const (
	Success         TResultCode = 200
	ValidationError TResultCode = 400
	AuthError       TResultCode = 401
	ForbiddenError  TResultCode = 403
	NotFoundError   TResultCode = 404
	InternalError   TResultCode = 502
)

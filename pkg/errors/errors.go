package errors

import "errors"

var (
	ErrorRepositoryUserAlreadyExsist = errors.New("user already exsist")
	ErrorServiceEmailInvalid         = errors.New("invalid email")
	ErrorRepositoryEmailNotExsist    = errors.New("email not exsist")
	ErrorInvalidToken                = errors.New("invalid token")
	ErrorKeyIdempotencyAlreadyUsed   = errors.New("key idempotency already used")
)

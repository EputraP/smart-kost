package errs

import "errors"

var (
	InvalidBearerFormat = errors.New("Invalid Authorization Bearer Format")
	InvalidToken        = errors.New("Invalid Token")
	InvalidIssuer       = errors.New("Invalid Token Issuer")
	InvalidIDParam      = errors.New("Invalid ID Parameter")

	ForbiddenAccess = errors.New("user is forbidden to access this resource")

	InvalidRequestBody  = errors.New("invalid request body")
	InvalidRequestParam = errors.New("invalid request param")

	EmailAlreadyUsed          = errors.New("email already used")
	UsernameAlreadyUsed       = errors.New("username already used")
	PasswordDoesntMatch       = errors.New("password doesn't match")
	PasswordContainUsername   = errors.New("password must not contain username")
	PasswordSameAsBefore      = errors.New("Password cannot be same as before")
	UsernamePasswordIncorrect = errors.New("username or password incorrect")

	ReferralCodeDoesnExist = errors.New("referral code doesnt exist")

	ImageNotFound = errors.New("ImageNotFound")
)

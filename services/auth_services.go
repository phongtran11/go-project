package services

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/phongtran11/go-project/database"
	"github.com/phongtran11/go-project/models"
	"github.com/phongtran11/go-project/pkg/constants"
	"github.com/phongtran11/go-project/pkg/dto/request"
	"github.com/phongtran11/go-project/pkg/dto/response"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type TAuthServices struct {
	Database     *gorm.DB
	UserServices *TUserServices
}

func AuthServices() *TAuthServices {
	return &TAuthServices{
		Database:     database.GetDB(),
		UserServices: UserServices(),
	}
}

func (authServices *TAuthServices) Register(request request.TRegisterRequest) error {
	// process register
	databaseError := authServices.UserServices.Create(&request)

	if databaseError != nil {
		return databaseError
	}

	return nil
}

func (authServices *TAuthServices) Login(request request.TLoginRequest) (*response.TTokenResponse, error) {
	user, err := authServices.UserServices.FindByEmail(request.Email)

	// error database
	if err != nil {
		return nil, &constants.TErrorMap{Message: constants.UserNotFound}
	}

	// compare password hash
	encryptErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if encryptErr != nil {
		return nil, &constants.TErrorMap{Message: constants.PasswordNotMatch}
	}

	// generate token
	tokenResult, tokenErr := authServices.generateToken(*user)

	// error generate token
	if tokenErr != nil {
		return nil, &constants.TErrorMap{Message: constants.GenerateTokenFailed}
	}

	return tokenResult, nil
}

func (authServices *TAuthServices) generateToken(user models.User) (*response.TTokenResponse, error) {
	tokenResult := &response.TTokenResponse{}

	tokenResult.AccessTokenExpireTime = time.Now().Add(1440 * time.Minute).Unix()
	tokenResult.RefreshTokenExpireTime = time.Now().Add(60 * time.Minute).Unix()

	// Generate access token
	accessTokenClaim := jwt.MapClaims{}

	accessTokenClaim["user_id"] = user.Id
	accessTokenClaim["email"] = user.Email
	accessTokenClaim["first_name"] = user.FirstName
	accessTokenClaim["last_name"] = user.LastName

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaim)

	accessTokenSigned, err := accessToken.SignedString([]byte("heWhy"))

	tokenResult.AccessToken = accessTokenSigned

	if err != nil {
		return nil, err
	}

	// Generate refresh token
	refreshTokenClaim := jwt.MapClaims{}

	refreshTokenClaim["user_id"] = user.Id
	refreshTokenClaim["expire_time"] = tokenResult.RefreshTokenExpireTime

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaim)

	refreshTokenSigned, err := refreshToken.SignedString([]byte("heWhyRefresh"))

	if err != nil {
		return nil, err
	}

	tokenResult.RefreshToken = refreshTokenSigned

	return tokenResult, nil
}

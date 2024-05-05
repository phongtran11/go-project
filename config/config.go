package config

import (
	"os"
	"strconv"
)

type TAppConfig struct {
	JTW_SECRET_KEY              string
	JTW_REFRESH_SECRET_KEY      string
	JTW_EXPIRATION_TIME         int
	JTW_REFRESH_EXPIRATION_TIME int
}

func GetAppConfig() *TAppConfig {
	jtwExpirationTime, _ := strconv.Atoi(os.Getenv("JWT_EXPIRATION_TIME"))
	jtwRefreshExpirationTime, _ := strconv.Atoi(os.Getenv("JWT_REFRESH_EXPIRATION_TIME"))

	return &TAppConfig{
		JTW_SECRET_KEY:              os.Getenv("JWT_SECRET_KEY"),
		JTW_REFRESH_SECRET_KEY:      os.Getenv("JWT_REFRESH_SECRET_KEY"),
		JTW_EXPIRATION_TIME:         jtwExpirationTime,
		JTW_REFRESH_EXPIRATION_TIME: jtwRefreshExpirationTime,
	}

}

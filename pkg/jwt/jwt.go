package jwt

import (
	"gin-blog/global"
	jwt "github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(global.Config.AppConfig.JwtSecret)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// GenerateToken 生成Token
func GenerateToken(username, password string) (string, error) {
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: global.Config.AppConfig.JwtExpiresat, //过期时间
			Issuer:    global.Config.AppConfig.JwtIssuer,    //签发人
		},
	})
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// ParseToken 解析token
func ParseToken(token string) (*Claims, error) {
	// 解析和校验token
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		// tokenClaims.Valid 返回token有效
		// 断言用户信息
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

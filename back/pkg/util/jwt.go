package util

import (
    "time"

    "Api/pkg/e"
    "Api/pkg/setting"
    "github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(setting.JwtSecret)

type Claims struct {
    Username string `json:"username"`
    jwt.StandardClaims
}

/**
 * @Description: 生成Token
 * @param username
 * @return string
 * @return error
 */
func GenerateToken(username string) (string, error) {
    nowTime := time.Now()
    expireTime := nowTime.Add(24 * time.Hour)
    claims := Claims{
        username,
        jwt.StandardClaims{
            IssuedAt:  nowTime.Unix(),
            NotBefore: nowTime.Unix(),
            ExpiresAt: expireTime.Unix(),
            Subject:   e.ParamsToken,
            Issuer:    "gin-blog",
            Audience:  username,
            Id:        username,
        },
    }
    tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    token, err := tokenClaims.SignedString(jwtSecret)

    return token, err
}

func ParseToken(token string) (*Claims, error) {
    tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        return jwtSecret, nil
    })
    if tokenClaims != nil {
        if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
            return claims, nil
        }
    }
    return nil, err
}

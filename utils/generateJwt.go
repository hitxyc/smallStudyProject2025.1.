package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

// 定义 JWT 签名密钥
var jwtSecret = []byte("your_secret_key") // 替换为你自己的密钥

// GenerateJwt 生成 JWT Token
func GenerateJwt(claims jwt.MapClaims) (string, error) {
	// 创建 Token 对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":  time.Now().Add(24 * time.Hour).Unix(), // 过期时间
		"iat":  time.Now().Unix(),                     // 签发时间
		"nbf":  time.Now().Unix(),                     // 生效时间
		"data": claims,                                // 自定义声明
	})

	// 使用密钥签名并生成 Token 字符串
	return token.SignedString(jwtSecret)
}

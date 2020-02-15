package token

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	Name string
	jwt.StandardClaims
}

const SecretKey = "secret key"

func MakeJWT(name string) (string, error) {
	// 사용자 클라임 정의
	claims := &Claims{
		Name: name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
		},
	}

	// 토큰 생성 에 필요한 Header 정보와 Payload에 들어갈 Cliam 전달
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Header와 Payload 값 기반으로 Signing
	// Signing 된 결과도 base64로 인코딩 되어 JWT에 dot(.) 구분자로 포함됨
	return token.SignedString([]byte(SecretKey)) // Header.Payload.Signature 형태로 반환
}

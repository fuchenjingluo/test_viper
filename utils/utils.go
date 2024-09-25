package utils

import (
	"fmt"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_trans "github.com/go-playground/validator/v10/translations/zh"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// 密码加密
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(hash), err
}

// 生成jwt令牌
type claims struct {
	User string `json:"user"`
	jwt.RegisteredClaims
}

func GenerateJwt(user string) (string, error) {
	claims := claims{
		User: user,
		RegisteredClaims: jwt.RegisteredClaims{
			// ExpiresAt 表示 JWT 的过期时间，设置为当前时间 3 小时后，过期后令牌将失效。
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 3)),
			// IssuedAt 表示 JWT 的签发时间，设置为当前时间，用于表示令牌什么时候被创建。
			IssuedAt: jwt.NewNumericDate(time.Now()),
			// NotBefore 表示 JWT 不能早于这个时间使用，设置为当前时间，意味着立刻可用。
			NotBefore: jwt.NewNumericDate(time.Now()),
			// Issuer 表示令牌的签发者，通常用来标识该 JWT 是由谁生成的。在这里是 "test_viper"。
			Issuer: "test_viper",
		},
	}
	fmt.Println(claims)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("secret"))
	return signedToken, err
}

// 解析token
func ParseJwt(tokenString string) (*claims, error) {
	// 解析 JWT 并使用指定的解析方法
	token, err := jwt.ParseWithClaims(tokenString, &claims{}, func(token *jwt.Token) (interface{}, error) {
		// 返回 JWT 的密钥（这里是 "secret"）
		return []byte("secret"), nil
	}, jwt.WithValidMethods([]string{"HS256"}))

	// 如果解析过程中出现错误，直接返回错误
	if err != nil {
		fmt.Println("Error parsing JWT:", err)
		return nil, err
	}

	// 检查 token 是否非空，且是有效的
	if token == nil || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// 提取并验证 claims
	if claims, ok := token.Claims.(*claims); ok {
		fmt.Println(claims)
		return claims, nil
	} else {
		return nil, fmt.Errorf("failed to parse claims")
	}
}

// 验证器语言

// 定义一个全局翻译器T
var Trans ut.Translator
var (
	uni      *ut.UniversalTranslator
	Validate *validator.Validate
)

func InitTrans() (err error) {
	zhT := zh.New() // 中文翻译器
	uni = ut.New(zhT, zhT)
	Trans, _ = uni.GetTranslator("zh")
	Validate = validator.New()
	err = zh_trans.RegisterDefaultTranslations(Validate, Trans)
	return
}

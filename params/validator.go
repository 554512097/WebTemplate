package params

import (
	"log"
	"regexp"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func InitCustomValidator() bool {
	// 注册验证
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("phone", phoneValid)
		if err != nil {
			log.Println("success")
			return false
		}
		return true
	}
	return false
}

func phoneValid(fl validator.FieldLevel) bool {
	// 生成手机号校验的正则表达式
	phoneRegexp := `^1[3456789]\d{9}$`
	// 使用正则表达式验证手机号
	return regexp.MustCompile(phoneRegexp).MatchString(fl.Field().String())
}

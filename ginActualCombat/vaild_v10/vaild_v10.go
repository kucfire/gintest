package main

import (
	"github.com/gin-gonic/gin"
	en2 "github.com/go-playground/locales/en"
	zh2 "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

type Person struct {
	Age     int    `form:"age" validate:"required,gt=10"`
	Name    string `form:"name" validate:"required"`
	Address string `form:"address"`
}

var (
	Uni      *ut.UniversalTranslator //
	Validate *validator.Validate     //
)

func main() {
	// 定义一个翻译器
	Validate = validator.New()
	zh := zh2.New() // 中文翻译器
	en := en2.New() // 英文翻译器
	Uni = ut.New(zh, en)

	r := gin.Default()
	r.GET("testing", func(c *gin.Context) {
		// 定义验证器
		locale := c.DefaultQuery("locale", "zh") // 默认设置locale为zh
		trans, _ := Uni.GetTranslator(locale)
		switch locale {
		case "zh":
			zh_translations.RegisterDefaultTranslations(Validate, trans)
			break
		case "en":
			en_translations.RegisterDefaultTranslations(Validate, trans)
			break
		default:
			zh_translations.RegisterDefaultTranslations(Validate, trans)
			break
		}

		// 实际逻辑
		var person Person
		if err := c.ShouldBind(&person); err != nil {
			c.String(500, "%v", err)
			c.Abort()
			return
		}

		if err := Validate.Struct(person); err != nil {
			errs := err.(validator.ValidationErrors)
			sliceErrs := []string{}
			for _, e := range errs {
				sliceErrs = append(sliceErrs, e.Translate(trans))
			}
			c.String(500, "%v", sliceErrs)
			c.Abort()
			return
		}

		c.String(200, "%v", person)
	})

	r.Run()
}

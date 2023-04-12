package common

import (
	"fmt"
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
)

const (
	StrRegexFuncName                    = "strregex"
	RequiredValidation                  = "required"
	OmitemptyValidation                 = "omitempty"
	AlphaTagValidation                  = "alpha"
	NumericTagValidation                = "numeric"
	AlphanumericTagValidation           = "alphanum"
	AlphanumericUnderScoreTagValidation = "alpha_num"
	AlphanumericHyphenTagValidation     = "alpha-num"
	LocationTagValidation               = "location"
	ValidCharactersTagValidation        = "charaters"
)

func CreateValidator() *validator.Validate {
	validate := validator.New()

	_ = validate.RegisterValidation(StrRegexFuncName, Strregex)

	return validate
}

func GetStrRegexTag(checkType string, isRequired bool) string {
	require := OmitemptyValidation
	if isRequired {
		require = RequiredValidation
	}
	return fmt.Sprintf("%s,%s=%s", require, StrRegexFuncName, checkType)
}

func Strregex(fl validator.FieldLevel) bool {
	var checkExpr string

	checkType := fl.Param()
	toCheckString := fl.Field().String()

	switch checkType {
	case AlphanumericUnderScoreTagValidation:
		checkExpr = `^[a-zA-Z0-9_]*$`
	case AlphanumericHyphenTagValidation:
		checkExpr = `^[a-zA-Z0-9-]*$`
	case LocationTagValidation:
		return validateLocation(toCheckString)
	case ValidCharactersTagValidation:
		checkExpr = `^[a-zA-Z0-9-+*/_]*$`
	case "date":
		checkExpr = `^(\d{4})-(\d{2})-(\d{2})`
	case "time":
		checkExpr = `^(\d{2}):(\d{2}):(\d{2}(?:\.\d*|Z)?)((-(\d{2}):(\d{2})|Z)?)$`
	case "datetime":
		fallthrough
	default:
		checkExpr = `^(\d{4})-(\d{2})-(\d{2})T(\d{2}):(\d{2}):(\d{2}(?:\.\d*|Z)?)((-(\d{2}):(\d{2})|Z)?)$`
	}

	reg, _ := regexp.Compile(checkExpr)

	return reg.MatchString(toCheckString)
}

func validateLocation(locationName string) bool {
	location, err := time.LoadLocation(locationName)
	if err != nil || location == nil {
		return false
	}
	return true
}

package validators

import (
	"reflect"
	"strings"

	"github.com/pkg/errors"
)

type Validator struct {
	ValidatorType string
	Value         string
}

type ValidationError struct {
	Field string `json:"field"`
	Err   error  `json:"error"`
}

type ValidationErrorMessage struct {
	Field        string `json:"field"`
	ErrorMessage string `json:"errorMessage"`
}

var (
	ErrValidationTagParse    = errors.New("failed to parse validator tag")
	ErrUnknownValidatorType  = errors.New("unknown validator type")
	ErrUnknownTypeToValidate = errors.New("unknown type to validate")

	supportedValidators = map[string]interface{}{
		"in":     nil,
		"min":    nil,
		"max":    nil,
		"len":    nil,
		"regexp": nil,
	}
)

func ParseValidateTag(tag reflect.StructTag) ([]Validator, error) {
	result := []Validator{}
	structValidationTag := tag.Get("validate")
	if structValidationTag == "" {
		return []Validator{}, nil
	}
	for _, s := range strings.Split(tag.Get("validate"), "|") {
		parsedValidator := strings.Split(s, ":")
		if len(parsedValidator) != 2 {
			return []Validator{}, errors.Wrap(ErrValidationTagParse, tag.Get("validate"))
		}
		validatorType := parsedValidator[0]
		if _, ok := supportedValidators[validatorType]; !ok {
			return []Validator{}, errors.Wrap(ErrUnknownValidatorType, tag.Get("validate"))
		}
		result = append(result, Validator{
			ValidatorType: parsedValidator[0],
			Value:         parsedValidator[1],
		})
	}
	return result, nil
}

func (v ValidationError) Error() string {
	return v.Err.Error()
}

package validators

import (
	"github.com/pkg/errors"
	"reflect"
	"strings"
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
	ValidationParseTagError = errors.New("failed to parse validator tag")
	UnknownValidatorType    = errors.New("unknown validator type")

	supportedValidators = []string{"len", "regexp", "in", "min", "max", "in"}
)

func ParseValidateTag(tag reflect.StructTag) ([]Validator, error) {
	var result []Validator
	structValidationTag := tag.Get("validate")
	if structValidationTag == "" {
		return []Validator{}, nil
	}
	for _, s := range strings.Split(tag.Get("validate"), "|") {
		parsedValidator := strings.Split(s, ":")
		if len(parsedValidator) != 2 {
			return []Validator{}, errors.Wrap(ValidationParseTagError, tag.Get("validate"))
		}
		ValidatorType := parsedValidator[0]
		if !Contains(supportedValidators, ValidatorType) {
			return []Validator{}, errors.Wrap(UnknownValidatorType, ValidatorType)
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

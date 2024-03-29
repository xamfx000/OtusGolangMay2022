package validators

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

var (
	ErrLengthValidation = errors.New("invalid length")
	ErrStringNotInSet   = errors.New("string not in set")
	ErrRegexMismatch    = errors.New("value not match with regex")
	ErrRegexCompile     = errors.New("failed to compile regex")
)

func ValidateStringField(validator Validator, val string, name string) error {
	switch validator.ValidatorType {
	case "len":
		expectedLength, err := strconv.ParseInt(validator.Value, 10, 64)
		if err != nil {
			return errors.Wrap(ErrIntParsing, validator.Value)
		}
		return validateStringLen(expectedLength, val, name)
	case "regexp":
		return validateStringRegexpMatch(validator.Value, val, name)
	case "in":
		return validateStringInSet(validator.Value, val, name)
	}
	return nil
}

func validateStringRegexpMatch(validatorValue string, val string, name string) error {
	re, err := regexp.Compile(validatorValue)
	if err != nil {
		return ErrRegexCompile
	}
	if !re.MatchString(val) {
		return ValidationError{
			Field: name,
			Err:   ErrRegexMismatch,
		}
	}
	return nil
}

func validateStringLen(validatorValue int64, val string, name string) error {
	length := len(strings.Split(val, ""))
	if int64(length) != validatorValue {
		return ValidationError{
			Field: name,
			Err:   ErrLengthValidation,
		}
	}
	return nil
}

func validateStringInSet(validatorValue string, val string, name string) error {
	allowedVals := strings.Split(validatorValue, ",")
	for _, v := range allowedVals {
		if v == val {
			return nil
		}
	}
	return ValidationError{
		Field: name,
		Err:   ErrStringNotInSet,
	}
}

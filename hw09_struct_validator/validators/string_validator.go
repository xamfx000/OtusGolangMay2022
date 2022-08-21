package validators

import (
	"fmt"
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
			return errors.Wrap(ErrIntParsing, fmt.Sprintf("%s", validator.Value))
		}
		return validateStringLen(expectedLength, val, name)
	case "regexp":
		return validateStringRegexpMatch(validator.Value, val, name)
	case "in":
		return validateStringInSet(validator.Value, val, name)
	}
	return nil
}

func validateStringRegexpMatch(value string, val string, name string) error {
	re, err := regexp.Compile(value)
	if err != nil {
		return errors.Wrap(ErrRegexCompile, err.Error())
	}
	if !re.MatchString(val) {
		return ValidationError{
			Field: name,
			Err:   errors.Wrap(ErrRegexMismatch, "field validation failed"),
		}
	}
	return nil
}

func validateStringLen(validatorValue int64, val string, name string) error {
	length := len(val)
	if int64(length) != validatorValue {
		return ValidationError{
			Field: name,
			Err:   errors.Wrap(ErrLengthValidation, "field validation failed"),
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
		Err:   errors.Wrap(ErrStringNotInSet, "field validation failed"),
	}
}

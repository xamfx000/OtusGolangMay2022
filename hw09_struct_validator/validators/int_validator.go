package validators

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

var (
	ErrIntExceedsMax = errors.New("max int")
	ErrIntBelowMin   = errors.New("min int")
	ErrIntNotInSet   = errors.New("not in set int")
	ErrIntParsing    = errors.New("failed to parse int")
)

func ValidateIntField(validator Validator, val int64, name string) error {
	switch validator.ValidatorType {
	case "min":
		return ValidateMinInt(validator.Value, val, name)
	case "max":
		return ValidateMaxInt(validator.Value, val, name)
	case "in":
		return ValidateInIntSet(validator.Value, val, name)
	}
	return nil
}

func ValidateMinInt(validatorValue string, val int64, name string) error {
	min, err := strconv.ParseInt(validatorValue, 10, 64)
	if err != nil {
		return errors.Wrap(ErrIntParsing, validatorValue)
	}
	if val < min {
		return ValidationError{
			Field: name,
			Err:   errors.Wrap(ErrIntBelowMin, "field validation failed"),
		}
	}
	return nil
}

func ValidateMaxInt(validatorValue string, val int64, name string) error {
	max, err := strconv.ParseInt(validatorValue, 10, 64)
	if err != nil {
		return errors.Wrap(ErrIntParsing, validatorValue)
	}
	if val > max {
		return ValidationError{
			Field: name,
			Err:   errors.Wrap(ErrIntExceedsMax, "field validation failed"),
		}
	}
	return nil
}

func ValidateInIntSet(validatorValue string, val int64, name string) error {
	stringVals := strings.Split(validatorValue, ",")
	for _, allowedValue := range stringVals {
		allowedValueParsed, err := strconv.ParseInt(allowedValue, 10, 64)
		if err != nil {
			return errors.Wrap(ErrIntParsing, validatorValue)
		}
		if allowedValueParsed == val {
			return nil
		}
	}
	return ValidationError{
		Field: name,
		Err:   errors.Wrap(ErrIntNotInSet, "field validation failed"),
	}
}

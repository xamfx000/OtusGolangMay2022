package validators

import (
	"fmt"
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

var (
	MaxIntError      = errors.New("max int")
	MinIntError      = errors.New("min int")
	IntNotInSetError = errors.New("not in set int")
	IntParsingError  = errors.New("failed to parse int")
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
		return errors.Wrap(IntParsingError, fmt.Sprintf("%s", validatorValue))
	}
	if val < min {
		return ValidationError{
			Field: name,
			Err:   errors.Wrap(MinIntError, "field validation failed"),
		}
	}
	return nil
}

func ValidateMaxInt(validatorValue string, val int64, name string) error {
	max, err := strconv.ParseInt(validatorValue, 10, 64)
	if err != nil {
		return errors.Wrap(IntParsingError, fmt.Sprintf("%s", validatorValue))
	}
	if val > max {
		return ValidationError{
			Field: name,
			Err:   errors.Wrap(MaxIntError, "field validation failed"),
		}
	}
	return nil
}

func ValidateInIntSet(validatorValue string, val int64, name string) error {
	stringVals := strings.Split(validatorValue, ",")
	for _, allowedValue := range stringVals {
		allowedValueParsed, err := strconv.ParseInt(allowedValue, 10, 64)
		if err != nil {
			return errors.Wrap(IntParsingError, fmt.Sprintf("%s", validatorValue))
		}
		if allowedValueParsed == val {
			return nil
		}
	}
	return ValidationError{
		Field: name,
		Err:   errors.Wrap(IntNotInSetError, "field validation failed"),
	}
}

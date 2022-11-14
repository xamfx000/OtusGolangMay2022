package hw09structvalidator

import (
	"encoding/json"
	"reflect"

	"github.com/pkg/errors"
	"github.com/xamfx000/OtusGolangMay2022/hw09_struct_validator/validators"
)

type ValidationErrors []validators.ValidationError

var ErrNonStructInputGiven = errors.New("non-struct given")

func (v ValidationErrors) Error() string {
	errorMessages := []validators.ValidationErrorMessage{}
	for _, e := range v {
		errorMessages = append(errorMessages, validators.ValidationErrorMessage{
			Field:        e.Field,
			ErrorMessage: e.Error(),
		})
	}
	result, _ := json.Marshal(errorMessages)
	return string(result)
}

func Validate(v interface{}) error {
	overallValidationResult := ValidationErrors{}
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Struct {
		return ErrNonStructInputGiven
	}
	rvt := rv.Type()
	for i := 0; i < rvt.NumField(); i++ {
		validatorsFromTag, err := validators.ParseValidateTag(rvt.Field(i).Tag)
		if err != nil {
			return err
		}
		for _, validator := range validatorsFromTag {
			if rv.Field(i).Kind() == reflect.Slice {
				sliceLen := rv.Field(i).Len()
				slice := rv.Field(i).Slice(0, sliceLen)
				for k := 0; k < sliceLen; k++ {
					overallValidationResult, err = processSingleField(
						slice.Index(k),
						rvt.Field(i),
						validator,
						overallValidationResult,
					)
					if err != nil {
						return err
					}
				}
				continue
			}
			overallValidationResult, err = processSingleField(
				rv.Field(i),
				rvt.Field(i),
				validator,
				overallValidationResult,
			)
			if err != nil {
				return err
			}
		}
	}
	return overallValidationResult
}

func processSingleField(
	fieldVal reflect.Value,
	structField reflect.StructField,
	validator validators.Validator,
	overallValidationResult ValidationErrors,
) (ValidationErrors, error) {
	err := validateField(fieldVal, structField, validator)
	if err != nil {
		var ve validators.ValidationError
		if !errors.As(err, &ve) {
			return nil, err
		}
		overallValidationResult = append(overallValidationResult, ve)
	}
	return overallValidationResult, nil
}

func validateField(field reflect.Value, structField reflect.StructField, validator validators.Validator) error {
	fieldName := structField.Name
	if structField.Tag == "" {
		return nil
	}
	switch field.Kind() { //nolint
	case reflect.Int:
		return validators.ValidateIntField(validator, field.Int(), fieldName)
	case reflect.String:
		return validators.ValidateStringField(validator, field.String(), fieldName)
	default:
		return validators.ErrUnknownTypeToValidate
	}
}

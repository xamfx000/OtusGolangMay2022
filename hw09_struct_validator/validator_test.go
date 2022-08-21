package hw09structvalidator

import (
	"fmt"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"github.com/xamfx000/OtusGolangMay2022/hw09_struct_validator/validators"
)

type UserRole string

// Test the function on different structures and other types.
type (
	User struct {
		ID     string `json:"id" validate:"len:36"`
		Name   string
		Age    int      `validate:"min:18|max:50"`
		Email  string   `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   UserRole `validate:"in:admin,stuff"`
		Phones []string `validate:"len:11"`
	}

	StructWithInvalidRegex struct {
		ID string `validate:"regexp:^../("`
	}

	Token struct {
		Header    []byte
		Payload   []byte
		Signature []byte
	}

	Response struct {
		Code int    `validate:"in:200,404,500"`
		Body string `json:"omitempty"`
	}

	ResponseCodes struct {
		Codes []int `validate:"min:100|max:500"`
	}

	StructWithInvalidTag struct {
		ID int `validate:"in:202,404,500,kek"`
	}

	StructWithCombinedValidator struct {
		ID int `validate:"in:202,404,500|max:500"`
	}

	StructWithUnsupportedValidator struct {
		ID int `validate:"whoami:13"`
	}
)

var tests = []struct {
	in          interface{}
	expectedErr error
}{
	{
		in:          "non-struct",
		expectedErr: ErrNonStructInputGiven,
	},
	{
		in: User{
			Age:   16,
			Role:  "admin",
			Email: "example@example..com",
		},
		expectedErr: ValidationErrors{
			validators.ValidationError{
				Field: "ID",
				Err:   validators.ErrLengthValidation,
			},
			validators.ValidationError{
				Field: "Age",
				Err:   validators.ErrIntBelowMin,
			},
			validators.ValidationError{
				Field: "Email",
				Err:   validators.ErrRegexMismatch,
			},
		},
	},
	{
		in: User{
			ID:    "123",
			Age:   52,
			Role:  "stuff",
			Email: "example@example.com",
		},
		expectedErr: ValidationErrors{
			validators.ValidationError{
				Field: "ID",
				Err:   validators.ErrLengthValidation,
			},
			validators.ValidationError{
				Field: "Age",
				Err:   validators.ErrIntExceedsMax,
			},
		},
	},
	{
		in: User{
			ID:    "123456123456123456123456123456123456",
			Age:   52,
			Email: "example@example.com",
		},
		expectedErr: ValidationErrors{
			validators.ValidationError{
				Field: "Age",
				Err:   validators.ErrIntExceedsMax,
			},
			validators.ValidationError{
				Field: "Role",
				Err:   validators.ErrStringNotInSet,
			},
		},
	},
	{
		in: User{
			ID:     "123456123456123456123456123456123456",
			Age:    50,
			Email:  "example@example.com",
			Phones: []string{"8999123456", "8999123455"},
			Role:   "stuff",
		},
		expectedErr: ValidationErrors{
			validators.ValidationError{
				Field: "Phones",
				Err:   validators.ErrLengthValidation,
			},
			validators.ValidationError{
				Field: "Phones",
				Err:   validators.ErrLengthValidation,
			},
		},
	},
	{
		in: Response{
			Code: 499,
		},
		expectedErr: ValidationErrors{
			validators.ValidationError{
				Field: "Code",
				Err:   validators.ErrIntNotInSet,
			},
		},
	},
	{
		in: Response{
			Code: 404,
		},
		expectedErr: nil,
	},
	{
		in: ResponseCodes{
			Codes: []int{123, 455, 666, 111},
		},
		expectedErr: ValidationErrors{
			validators.ValidationError{
				Field: "Codes",
				Err:   validators.ErrIntExceedsMax,
			},
		},
	},
	{
		in: Token{
			Header:    []byte("header"),
			Payload:   []byte("payload"),
			Signature: []byte("signature"),
		},
		expectedErr: nil,
	},
	{
		in: StructWithInvalidRegex{
			ID: "dgdf",
		},
		expectedErr: validators.ErrRegexCompile,
	},
	{
		in: StructWithCombinedValidator{
			ID: 600,
		},
		expectedErr: ValidationErrors{
			validators.ValidationError{
				Field: "ID",
				Err:   validators.ErrIntNotInSet,
			},
			validators.ValidationError{
				Field: "ID",
				Err:   validators.ErrIntExceedsMax,
			},
		},
	},
	{
		in: StructWithInvalidTag{
			ID: 124125,
		},
		expectedErr: validators.ErrIntParsing,
	},
	{
		in: StructWithUnsupportedValidator{
			ID: 124125,
		},
		expectedErr: validators.UnknownValidatorType,
	},
}

func TestValidate(t *testing.T) {
	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tt := tt
			t.Parallel()
			validationResult := Validate(tt.in)
			var ve ValidationErrors
			if errors.As(validationResult, &ve) {
				errorsSlice := prepareErrSliceFromRootErr(ve)
				for i, err := range errorsSlice {
					require.Equal(t, tt.expectedErr.(ValidationErrors)[i].Field, validationResult.(ValidationErrors)[i].Field, i)
					require.ErrorAs(t, err.Err, &tt.expectedErr.(ValidationErrors)[i].Err)
				}
			} else {
				require.ErrorIs(t, validationResult, tt.expectedErr)
			}
		})
	}
}

func prepareErrSliceFromRootErr(errors ValidationErrors) []validators.ValidationError {
	result := []validators.ValidationError{}
	for _, err := range errors {
		result = append(result, err)
	}
	return result
}

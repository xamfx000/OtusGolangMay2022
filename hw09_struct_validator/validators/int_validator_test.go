package validators

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateMinInt(t *testing.T) {
	type args struct {
		validatorValue string
		val            int64
		name           string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "validator passed",
			args: args{
				validatorValue: "18",
				val:            int64(20),
				name:           "test",
			},
			wantErr: false,
		},
		{
			name: "validator failed",
			args: args{
				validatorValue: "18",
				val:            int64(15),
				name:           "test",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateMinInt(tt.args.validatorValue, tt.args.val, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateMinInt() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				require.ErrorAs(t, err, &ErrIntBelowMin)
			}
		})
	}
}

func TestValidateMaxInt(t *testing.T) {
	type args struct {
		validatorValue string
		val            int64
		name           string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "validator passed",
			args: args{
				validatorValue: "18",
				val:            int64(15),
				name:           "test",
			},
			wantErr: false,
		},
		{
			name: "validator failed",
			args: args{
				validatorValue: "18",
				val:            int64(20),
				name:           "test",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateMaxInt(tt.args.validatorValue, tt.args.val, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateMaxInt() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				require.ErrorAs(t, err, &ErrIntExceedsMax)
			}
		})
	}
}

func TestValidateInIntSet(t *testing.T) {
	type args struct {
		validatorValue string
		val            int64
		name           string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "validator passed",
			args: args{
				validatorValue: "18,13,22,33,55,15,66,23,26",
				val:            int64(15),
				name:           "test",
			},
			wantErr: false,
		},
		{
			name: "validator failed",
			args: args{
				validatorValue: "18,13,22,33,55,15,66,23,26",
				val:            int64(20),
				name:           "test",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateInIntSet(tt.args.validatorValue, tt.args.val, tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("ValidateInIntSet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

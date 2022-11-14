package validators

import "testing"

func Test_validateStringLen(t *testing.T) {
	type args struct {
		validatorValue int64
		val            string
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
				validatorValue: 4,
				val:            "test",
				name:           "test",
			},
			wantErr: false,
		},
		{
			name: "validator passed with cyrillic val",
			args: args{
				validatorValue: 4,
				val:            "тест",
				name:           "test",
			},
			wantErr: false,
		},
		{
			name: "validator failed",
			args: args{
				validatorValue: 5,
				val:            "test",
				name:           "test",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateStringLen(tt.args.validatorValue, tt.args.val, tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("validateStringLen() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_validateStringInSet(t *testing.T) {
	type args struct {
		validatorValue string
		val            string
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
				validatorValue: "т,е,с,т,о,в,ы,й,с,е,т",
				val:            "т",
				name:           "test",
			},
			wantErr: false,
		},
		{
			name: "validator failed",
			args: args{
				validatorValue: "т,е,с,т,о,в,ы,й,с,е,т",
				val:            "test",
				name:           "test",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateStringInSet(tt.args.validatorValue, tt.args.val, tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("validateStringInSet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_validateStringRegexpMatch(t *testing.T) {
	type args struct {
		validatorValue string
		val            string
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
				validatorValue: `^\w+@\w+\.\w+$`,
				val:            "valid@email.test",
				name:           "test",
			},
			wantErr: false,
		},
		{
			name: "validator failed",
			args: args{
				validatorValue: `^\w+@\w+\.\w+$`,
				val:            "invalid_email.test",
				name:           "test",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateStringRegexpMatch(tt.args.validatorValue, tt.args.val, tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("validateStringRegexpMatch() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

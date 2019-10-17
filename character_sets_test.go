package spinner

import (
	"testing"
)

func Test_checkCharSet(t *testing.T) {
	type args struct {
		c []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"check empty char set",
			args{[]string{}},
			false,
		},
		{
			"char set is too big",
			args{returnBigCharSet(maxCharSetSize)},
			true,
		},
		{
			"numbers char set",
			args{[]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}},
			false,
		},
		{
			"ambiguous widths char set",
			args{[]string{"0", "  ", "0",}},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := checkCharSet(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("checkCharSet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func returnBigCharSet(size int) []string {
	big := make([]string, size + 5)
	for i:= range big {
		big[i] = "-"
	}
	return big
}
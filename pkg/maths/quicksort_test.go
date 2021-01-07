package maths

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestDecimalQuickSort(t *testing.T) {
	type args struct {
		v    []decimal.Decimal
		low  int
		high int
	}
	tests := []struct {
		name    string
		args    args
		want    []decimal.Decimal
		wantErr bool
	}{
		{
			name: "sort values",
			args: args{
				v: []decimal.Decimal{
					decimal.NewFromInt(10),
					decimal.NewFromInt(7),
					decimal.NewFromInt(8),
					decimal.NewFromInt(9),
					decimal.NewFromInt(1),
					decimal.NewFromInt(5),
				},
				low:  0,
				high: 5,
			},
			want: []decimal.Decimal{
				decimal.NewFromInt(1),
				decimal.NewFromInt(5),
				decimal.NewFromInt(7),
				decimal.NewFromInt(8),
				decimal.NewFromInt(9),
				decimal.NewFromInt(10),
			},
			wantErr: false,
		},
		{
			name: "empty list",
			args: args{
				v:    nil,
				low:  0,
				high: 0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "fractional values",
			args: args{
				v: []decimal.Decimal{
					decimal.NewFromFloat(1.0001),
					decimal.NewFromFloat(1.00001),
					decimal.NewFromFloat(1.0000000001),
					decimal.NewFromFloat(1.00000001),
					decimal.NewFromFloat(1.0000000000001),
					decimal.NewFromFloat(1.0000001),
				},
				low:  0,
				high: 5,
			},
			want: []decimal.Decimal{
				decimal.NewFromFloat(1.0000000000001),
				decimal.NewFromFloat(1.0000000001),
				decimal.NewFromFloat(1.00000001),
				decimal.NewFromFloat(1.0000001),
				decimal.NewFromFloat(1.00001),
				decimal.NewFromFloat(1.0001),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DecimalQuickSort(tt.args.v, tt.args.low, tt.args.high)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecimalQuickSort() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			for i, v := range got {
				if !tt.want[i].Equal(v) {
					t.Errorf("DecimalQuickSort() i = %d, got = %s, want = %s", i, v.String(), tt.want[i].String())
				}
			}
		})
	}
}

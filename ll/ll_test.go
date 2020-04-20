package ll

import (
	"reflect"
	"testing"
)

func TestGenericLL_Get(t *testing.T) {

	type args struct {
		index int
	}
	tests := []struct {
		name      string
		ll        *GenericLL
		args      args
		want      Generic
		wantPanic bool
	}{
		{
			name:      "Getting a value",
			ll:        NewGenericLL(1, 2, 3, 4),
			args:      args{index: 1},
			want:      2,
			wantPanic: false,
		},
		{
			name:      "Getting first value",
			ll:        NewGenericLL(1, 2, 3, 4),
			args:      args{index: 0},
			want:      1,
			wantPanic: false,
		},
		{
			name:      "Getting last value",
			ll:        NewGenericLL(1, 2, 3, 4),
			args:      args{index: 3},
			want:      4,
			wantPanic: false,
		},
		{
			name:      "Negative index panic",
			ll:        NewGenericLL(1, 2, 3, 4),
			args:      args{index: -1},
			want:      0,
			wantPanic: true,
		},
		{
			name:      "Out of range panic",
			ll:        NewGenericLL(1, 2, 3, 4),
			args:      args{index: 4},
			want:      0,
			wantPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tt.wantPanic {
					t.Errorf("GenericLL.Get() recover = %v, wantPanic = %v", r, tt.wantPanic)
				}
			}()
			if got := tt.ll.Get(tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenericLL.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

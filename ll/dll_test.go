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
		ll        *GenericDLL
		args      args
		want      Generic
		wantPanic bool
	}{
		{
			name:      "Getting a value",
			ll:        NewGenericDLL(1, 2, 3, 4),
			args:      args{index: 1},
			want:      2,
			wantPanic: false,
		},
		{
			name:      "Getting first value",
			ll:        NewGenericDLL(1, 2, 3, 4),
			args:      args{index: 0},
			want:      1,
			wantPanic: false,
		},
		{
			name:      "Getting last value",
			ll:        NewGenericDLL(1, 2, 3, 4),
			args:      args{index: 3},
			want:      4,
			wantPanic: false,
		},
		{
			name:      "Negative index panic",
			ll:        NewGenericDLL(1, 2, 3, 4),
			args:      args{index: -1},
			want:      0,
			wantPanic: true,
		},
		{
			name:      "Out of range panic",
			ll:        NewGenericDLL(1, 2, 3, 4),
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

func TestGenericDLL_InsertAfter(t *testing.T) {
	type args struct {
		index int
		data  Generic
	}
	tests := []struct {
		name      string
		ll        *GenericDLL
		args      args
		wantPanic bool
	}{
		{
			name:      "Insert in first half",
			ll:        NewGenericDLL(1, 2, 3, 4, 5),
			args:      args{index: 1, data: 6},
			wantPanic: false,
		},
		{
			name:      "Insert in second half",
			ll:        NewGenericDLL(1, 2, 3, 4, 5),
			args:      args{index: 3, data: 6},
			wantPanic: false,
		},
		{
			name:      "Insert at end",
			ll:        NewGenericDLL(1, 2, 3, 4, 5),
			args:      args{index: 4, data: 6},
			wantPanic: false,
		},
		{
			name:      "Negative index",
			ll:        NewGenericDLL(1, 2, 3, 4, 5),
			args:      args{index: -1, data: 6},
			wantPanic: true,
		},
		{
			name:      "Out of range index",
			ll:        NewGenericDLL(1, 2, 3, 4, 5),
			args:      args{index: 10, data: 6},
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
			tt.ll.InsertAfter(tt.args.index, tt.args.data)
			if v := tt.ll.Get(tt.args.index + 1); v != tt.args.data {
				t.Errorf("GenericLL.InsertAfter() inserted %v, want %v", v, tt.args.data)
			}
		})
	}
}

func TestGenericDLL_String(t *testing.T) {
	tests := []struct {
		name string
		ll   *GenericDLL
		want string
	}{
		{
			name: "Empty list",
			ll:   NewGenericDLL(),
			want: "[]",
		},
		{
			name: "One element",
			ll:   NewGenericDLL(1),
			want: "[1]",
		},
		{
			name: "Multiple elements element",
			ll:   NewGenericDLL(1, 2, 3),
			want: "[1 2 3]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ll.String(); got != tt.want {
				t.Errorf("GenericDLL.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenericDLL_Prepend(t *testing.T) {
	type args struct {
		data Generic
	}
	tests := []struct {
		name string
		ll   *GenericDLL
		args args
		want *GenericDLL
	}{
		{
			name: "Empty list",
			ll:   NewGenericDLL(),
			args: args{data: 1},
			want: NewGenericDLL(1),
		},
		{
			name: "One Element",
			ll:   NewGenericDLL(2),
			args: args{data: 1},
			want: NewGenericDLL(1, 2),
		},
		{
			name: "Multiple elements",
			ll:   NewGenericDLL(2, 3),
			args: args{data: 1},
			want: NewGenericDLL(1, 2, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ll.Prepend(tt.args.data)
			if !reflect.DeepEqual(tt.ll, tt.want) {
				t.Errorf("GenericLL.Prepend() = %v, want %v", tt.ll, tt.want)
			}
		})
	}
}

func TestGenericDLL_Enumerate(t *testing.T) {
	tests := []struct {
		name string
		ll   *GenericDLL
		want []Generic
	}{
		{
			name: "Empty list",
			ll:   NewGenericDLL(),
			want: []Generic{},
		},
		{
			name: "One Element",
			ll:   NewGenericDLL(1),
			want: []Generic{1},
		},
		{
			name: "Multiple elements",
			ll:   NewGenericDLL(1, 2, 3),
			want: []Generic{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ll.Enumerate(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenericDLL.Enumerate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenericDLL_Remove(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name      string
		ll        *GenericDLL
		args      args
		want      Generic
		wantPanic bool
		wantLL    *GenericDLL
	}{
		{
			name:      "Empty list",
			ll:        NewGenericDLL(),
			args:      args{index: 0},
			wantPanic: true,
			want:      nil,
			wantLL:    nil,
		},
		{
			name:      "One Element",
			ll:        NewGenericDLL(1),
			args:      args{index: 0},
			wantPanic: false,
			want:      1,
			wantLL:    NewGenericDLL(),
		},
		{
			name:      "Multiple elements - head",
			ll:        NewGenericDLL(1, 2, 3),
			args:      args{index: 0},
			wantPanic: false,
			want:      1,
			wantLL:    NewGenericDLL(2, 3),
		},
		{
			name:      "Multiple elements - middle",
			ll:        NewGenericDLL(1, 2, 3),
			args:      args{index: 1},
			wantPanic: false,
			want:      2,
			wantLL:    NewGenericDLL(1, 3),
		},
		{
			name:      "Multiple elements - tail",
			ll:        NewGenericDLL(1, 2, 3),
			args:      args{index: 2},
			wantPanic: false,
			want:      3,
			wantLL:    NewGenericDLL(1, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tt.wantPanic {
					t.Errorf("GenericLL.Remove() recover = %v, wantPanic = %v", r, tt.wantPanic)
				}
			}()
			if got := tt.ll.Remove(tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenericDLL.Remove() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(tt.ll, tt.wantLL) {
				t.Errorf("GenericDLL.Remove() LL is %v, want %v", tt.ll, tt.wantLL)
			}
		})
	}
}

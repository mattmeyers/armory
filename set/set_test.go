package set

import (
	"reflect"
	"testing"
)

func TestNewGenericSet(t *testing.T) {
	type args struct {
		vals []Generic
	}
	tests := []struct {
		name string
		args args
		want *GenericSet
	}{
		{
			name: "New empty set",
			args: args{[]Generic{}},
			want: &GenericSet{vals: map[Generic]bool{}},
		},
		{
			name: "New set with values",
			args: args{[]Generic{1, 1, 2, 3}},
			want: &GenericSet{vals: map[Generic]bool{1: true, 2: true, 3: true}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGenericSet(tt.args.vals...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGenericSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenericSet_Contains(t *testing.T) {
	type fields struct {
		vals map[Generic]bool
		cap  int
	}
	type args struct {
		val Generic
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "Does contain",
			fields: fields{vals: map[Generic]bool{1: true}},
			args:   args{1},
			want:   true,
		},
		{
			name:   "Does not contain",
			fields: fields{vals: map[Generic]bool{1: true}},
			args:   args{2},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &GenericSet{
				vals: tt.fields.vals,
				cap:  tt.fields.cap,
			}
			if got := s.Contains(tt.args.val); got != tt.want {
				t.Errorf("GenericSet.Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenericSet_IsSubset(t *testing.T) {
	type fields struct {
		vals map[Generic]bool
		cap  int
	}
	type args struct {
		b GenericSet
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "Proper subset",
			fields: fields{vals: map[Generic]bool{1: true, 3: true}},
			args:   args{b: GenericSet{vals: map[Generic]bool{1: true, 2: true, 3: true}}},
			want:   true,
		},
		{
			name:   "Equal set",
			fields: fields{vals: map[Generic]bool{1: true, 2: true, 3: true}},
			args:   args{b: GenericSet{vals: map[Generic]bool{1: true, 2: true, 3: true}}},
			want:   true,
		},
		{
			name:   "Not subset",
			fields: fields{vals: map[Generic]bool{1: true, 3: true}},
			args:   args{b: GenericSet{vals: map[Generic]bool{1: true, 2: true, 4: true}}},
			want:   false,
		},
		{
			name:   "Empty sets",
			fields: fields{vals: map[Generic]bool{}},
			args:   args{b: GenericSet{vals: map[Generic]bool{}}},
			want:   true,
		},
		{
			name:   "More elements",
			fields: fields{vals: map[Generic]bool{1: true, 3: true}},
			args:   args{b: GenericSet{vals: map[Generic]bool{1: true}}},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &GenericSet{
				vals: tt.fields.vals,
				cap:  tt.fields.cap,
			}
			if got := s.IsSubset(tt.args.b); got != tt.want {
				t.Errorf("GenericSet.IsSubset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenericSet_Enumerate(t *testing.T) {
	type fields struct {
		vals map[Generic]bool
		cap  int
	}
	tests := []struct {
		name   string
		fields fields
		want   []Generic
	}{
		{
			name:   "Empty set",
			fields: fields{vals: map[Generic]bool{}},
			want:   []Generic{},
		},
		{
			name:   "Enumerate keys",
			fields: fields{vals: map[Generic]bool{2: true, 1: true}},
			want:   []Generic{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &GenericSet{
				vals: tt.fields.vals,
				cap:  tt.fields.cap,
			}

			got := s.Enumerate()

			if len(tt.fields.vals) != len(got) {
				t.Errorf("GenericSet.Enumerate() = %v, want %v", got, tt.want)
			}

			for _, i := range got {
				if _, ok := tt.fields.vals[i]; !ok {
					t.Errorf("GenericSet.Enumerate() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

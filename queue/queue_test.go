package queue

import (
	"reflect"
	"testing"
)

func TestNewGenericQueue(t *testing.T) {
	type args struct {
		vals []Generic
	}
	tests := []struct {
		name string
		args args
		want *GenericQueue
	}{
		{
			name: "New empty queue",
			args: args{[]Generic{}},
			want: &GenericQueue{vals: []Generic{}},
		},
		{
			name: "New queue with values",
			args: args{[]Generic{1, 2, 3}},
			want: &GenericQueue{vals: []Generic{1, 2, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGenericQueue(tt.args.vals...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGenericQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenericQueue_IsEmpty(t *testing.T) {
	type fields struct {
		vals []Generic
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "Queue is empty",
			fields: fields{},
			want:   true,
		},
		{
			name:   "Queue is not empty",
			fields: fields{vals: []Generic{1}},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &GenericQueue{
				vals: tt.fields.vals,
			}
			if got := q.IsEmpty(); got != tt.want {
				t.Errorf("GenericQueue.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenericQueue_String(t *testing.T) {
	type fields struct {
		vals []Generic
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "Empty queue string",
			fields: fields{},
			want:   "[]",
		},
		{
			name:   "Queue is not empty",
			fields: fields{vals: []Generic{1, 2, 3}},
			want:   "[ 1 2 3 ]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &GenericQueue{
				vals: tt.fields.vals,
			}
			if got := q.String(); got != tt.want {
				t.Errorf("GenericQueue.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenericQueue_Push(t *testing.T) {
	type fields struct {
		vals []Generic
	}
	type args struct {
		val Generic
	}
	tests := []struct {
		name   string
		fields fields
		args   []args
		want   *GenericQueue
	}{
		{
			name:   "Push one value",
			fields: fields{},
			args: []args{
				{val: 1},
			},
			want: &GenericQueue{vals: []Generic{1}},
		},
		{
			name:   "Push multiple value",
			fields: fields{},
			args: []args{
				{val: 1},
				{val: 2},
			},
			want: &GenericQueue{vals: []Generic{1, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &GenericQueue{
				vals: tt.fields.vals,
			}

			for _, a := range tt.args {
				q.Push(a.val)
			}

			if !reflect.DeepEqual(q, tt.want) {
				t.Errorf("GenericQueue.Push() = %v, want %v", q, tt.want)
			}
		})
	}
}

func TestGenericQueue_Pop(t *testing.T) {
	type fields struct {
		vals []Generic
	}
	tests := []struct {
		name   string
		fields fields
		want   Generic
		after  *GenericQueue
	}{
		{
			name:   "Pop empty queue",
			fields: fields{},
			want:   nil,
			after:  &GenericQueue{},
		},
		{
			name:   "Pop single value",
			fields: fields{vals: []Generic{1, 2, 3}},
			want:   1,
			after:  &GenericQueue{vals: []Generic{2, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &GenericQueue{
				vals: tt.fields.vals,
			}

			got := q.Pop()

			if got != tt.want {
				t.Errorf("GenericQueue.Pop() = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(q, tt.after) {
				t.Errorf("GenericQueue.vals = %v, after %v", q, tt.after)
			}
		})
	}
}

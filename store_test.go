package main

import (
	"reflect"
	"sync"
	"testing"
)

func TestNewStore(t *testing.T) {
	tests := []struct {
		name string
		want *Store
	}{
		{
			name: "crete new store",
			want: NewStore(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStore(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStore_Set(t *testing.T) {
	type fields struct {
		mu    sync.RWMutex
		Items map[string]string
	}
	type args struct {
		key string
		val string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "set new key",
			fields: fields{sync.RWMutex{}, make(map[string]string)},
			args:   args{"test", "11"},
			want:   "11",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Store{
				mu:    tt.fields.mu,
				Items: tt.fields.Items,
			}
			s.Set(tt.args.key, tt.args.val)

			if got := s.Get(tt.args.key); got != tt.want {
				t.Errorf("Store.Set() = %v, want %v", got, tt.want)
			}

		})
	}
}

func TestStore_Get(t *testing.T) {
	type fields struct {
		mu    sync.RWMutex
		Items map[string]string
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "get test key",
			fields: fields{sync.RWMutex{}, map[string]string{"test": "11"}},
			args:   args{"test"},
			want:   "11",
		},
		{
			name:   "get testsecond key",
			fields: fields{sync.RWMutex{}, map[string]string{"testsecond": "12"}},
			args:   args{"testsecond"},
			want:   "12",
		},
		{
			name:   "get not found key",
			fields: fields{sync.RWMutex{}, map[string]string{"testsecond": "12"}},
			args:   args{"test"},
			want:   "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Store{
				mu:    tt.fields.mu,
				Items: tt.fields.Items,
			}

			if got := s.Get(tt.args.key); got != tt.want {
				t.Errorf("Store.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStore_Del(t *testing.T) {
	type fields struct {
		mu    sync.RWMutex
		Items map[string]string
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "del test key",
			fields: fields{sync.RWMutex{}, map[string]string{"test": "11", "testsecond": "12"}},
			args:   args{"test"},
			want:   "",
		},
		{
			name:   "del testsecond key",
			fields: fields{sync.RWMutex{}, map[string]string{"test": "11", "testsecond": "12"}},
			args:   args{"testsecond"},
			want:   "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Store{
				mu:    tt.fields.mu,
				Items: tt.fields.Items,
			}
			s.Del(tt.args.key)

			if got := s.Get(tt.args.key); got != tt.want {
				t.Errorf("Store.Del() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStore_GetAll(t *testing.T) {
	type fields struct {
		mu    sync.RWMutex
		Items map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]string
	}{
		{
			name:   "get all items",
			fields: fields{sync.RWMutex{}, map[string]string{"test": "11", "testsecond": "12"}},
			want:   map[string]string{"test": "11", "testsecond": "12"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Store{
				mu:    tt.fields.mu,
				Items: tt.fields.Items,
			}
			if got := s.GetAll(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Store.GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStore_Stats(t *testing.T) {
	type fields struct {
		mu    sync.RWMutex
		Items map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "count store items",
			fields: fields{sync.RWMutex{}, map[string]string{"test": "11", "testsecond": "12"}},
			want:   2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Store{
				mu:    tt.fields.mu,
				Items: tt.fields.Items,
			}
			if got := s.Stats(); got != tt.want {
				t.Errorf("Store.Stats() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStore_Flush(t *testing.T) {
	type fields struct {
		mu    sync.RWMutex
		Items map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "flush items",
			fields: fields{sync.RWMutex{}, map[string]string{"test": "11", "testsecond": "12"}},
			want:   0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Store{
				mu:    tt.fields.mu,
				Items: tt.fields.Items,
			}
			s.Flush()

			if got := s.Stats(); got != tt.want {
				t.Errorf("Store.Flush() = %v, want %v", got, tt.want)
			}
		})
	}
}

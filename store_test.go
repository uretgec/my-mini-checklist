package main

import (
	"os"
	"reflect"
	"sync"
	"testing"
)

func TestNewStore(t *testing.T) {
	tests := []struct {
		name string
		want *Store
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStore(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStore_Sync(t *testing.T) {
	type fields struct {
		mu    sync.RWMutex
		Items map[string]string
	}
	type args struct {
		db        *os.File
		firstTime bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Store{
				mu:    tt.fields.mu,
				Items: tt.fields.Items,
			}
			if err := s.Sync(tt.args.db, tt.args.firstTime); (err != nil) != tt.wantErr {
				t.Errorf("Store.Sync() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStore_Load(t *testing.T) {
	type fields struct {
		mu    sync.RWMutex
		Items map[string]string
	}
	type args struct {
		db *os.File
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Store{
				mu:    tt.fields.mu,
				Items: tt.fields.Items,
			}
			s.Load(tt.args.db)
		})
	}
}

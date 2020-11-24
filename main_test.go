package main

import (
	"net/http"
	"testing"
)

func TestHelloServer(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HelloServer(tt.args.w, tt.args.r)
		})
	}
}

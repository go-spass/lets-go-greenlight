package main

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestReadIDParam(t *testing.T) {
	tests := []struct {
		name    string
		id      string
		want    int64
		wantErr bool
	}{
		{name: "valid", id: "1", want: 1},
		{name: "large valid", id: "9223372036854775807", want: 9223372036854775807},
		{name: "zero", id: "0", wantErr: true},
		{name: "negative", id: "-1", wantErr: true},
		{name: "not a number", id: "abc", wantErr: true},
		{name: "empty", id: "", wantErr: true},
		{name: "overflow", id: "99999999999999999999", wantErr: true},
	}

	app := &application{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest("GET", "/v1/movies/"+tt.id, nil)
			ctx := context.WithValue(r.Context(), httprouter.ParamsKey, httprouter.Params{{Key: "id", Value: tt.id}})
			r = r.WithContext(ctx)

			got, err := app.readIDParam(r)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("readIDParam(%q) = %d, want error", tt.id, got)
				}
				return
			}
			if err != nil {
				t.Fatalf("readIDParam(%q) unexpected error: %v", tt.id, err)
			}
			if got != tt.want {
				t.Errorf("readIDParam(%q) = %d, want %d", tt.id, got, tt.want)
			}
		})
	}
}

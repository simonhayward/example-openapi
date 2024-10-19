package api

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

func TestServeUser(t *testing.T) {

	tests := []struct {
		id         openapi_types.UUID
		link       string
		wantStatus int
		wantData   User
	}{
		{
			id:         uuid.MustParse("62c29903-df2f-49da-9c52-d07b92e5a840"),
			link:       "/user/62c29903-df2f-49da-9c52-d07b92e5a840",
			wantStatus: http.StatusOK,
			wantData:   User{Id: uuid.MustParse("62c29903-df2f-49da-9c52-d07b92e5a840")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.link, func(t *testing.T) {
			s := NewServer()
			r := httptest.NewRequest("GET", tt.link, nil)
			w := httptest.NewRecorder()
			s.GetUserById(w, r, tt.id)

			if w.Code != tt.wantStatus {
				t.Errorf("TestServeUser(%q) got %d; want %d", tt.link, w.Code, tt.wantStatus)
			}
			res := w.Result()
			defer res.Body.Close()
			data, err := io.ReadAll(res.Body)
			if err != nil {
				t.Errorf("expected error to be nil got %v", err)
			}
			var u User
			err = json.Unmarshal(data, &u)
			if err != nil {
				t.Errorf("expected error to be nil got %v", err)
			}
			if u != tt.wantData {
				t.Errorf("TestServeUser(%q) got %v; want %v", tt.link, u, tt.wantData)
			}
		})
	}
}

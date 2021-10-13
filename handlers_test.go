package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_home(t *testing.T) {
	tests := []struct {
		name       string
		method     string
		statusCode int
		key        string
		value      string
		want       string
	}{
		{
			name:       "welcome home",
			method:     http.MethodGet,
			want:       `{"status":"success","message":"Welcome Home","result":null}`,
			statusCode: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(tt.method, "/", nil)
			responseRecorder := httptest.NewRecorder()

			home().ServeHTTP(responseRecorder, request)

			if responseRecorder.Code != tt.statusCode {
				t.Errorf("Want status '%d', got '%d'", tt.statusCode, responseRecorder.Code)
			}

			if strings.TrimSpace(responseRecorder.Body.String()) != tt.want {
				t.Errorf("Want '%s', got '%s'", tt.want, responseRecorder.Body)
			}
		})
	}
}

func Test_apiHome(t *testing.T) {
	tests := []struct {
		name       string
		method     string
		statusCode int
		key        string
		value      string
		want       string
	}{
		{
			name:       "api home",
			method:     http.MethodGet,
			want:       `{"status":"success","message":"Api Home","result":null}`,
			statusCode: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(tt.method, "/api", nil)
			responseRecorder := httptest.NewRecorder()

			apiHome().ServeHTTP(responseRecorder, request)

			if responseRecorder.Code != tt.statusCode {
				t.Errorf("Want status '%d', got '%d'", tt.statusCode, responseRecorder.Code)
			}

			if strings.TrimSpace(responseRecorder.Body.String()) != tt.want {
				t.Errorf("Want '%s', got '%s'", tt.want, responseRecorder.Body)
			}
		})
	}
}

func Test_apiStoreSet(t *testing.T) {
	tests := []struct {
		name       string
		method     string
		statusCode int
		key        string
		value      string
		want       string
	}{
		{
			name:       "api store set action get request",
			method:     http.MethodGet,
			key:        "test",
			value:      "12",
			want:       `{"status":"success","message":"test key updated","result":"12"}`,
			statusCode: http.StatusOK,
		},
		{
			name:       "api store set action post request",
			method:     http.MethodPost,
			key:        "testsecond",
			value:      "11",
			want:       `{"status":"success","message":"testsecond key updated","result":"11"}`,
			statusCode: http.StatusOK,
		},
		{
			name:       "api store set action empty request",
			method:     http.MethodPost,
			key:        "",
			value:      "",
			want:       `{"status":"error","message":" Not found!","result":null}`,
			statusCode: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			store = NewStore()

			request := httptest.NewRequest(tt.method, "/api/set?key="+tt.key+"&value="+tt.value, nil)
			if tt.method == "POST" {
				request := httptest.NewRequest(tt.method, "/api/set", strings.NewReader("key="+tt.key+"&value="+tt.value))
				request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
			}
			responseRecorder := httptest.NewRecorder()

			apiStoreSet().ServeHTTP(responseRecorder, request)

			if responseRecorder.Code != tt.statusCode {
				t.Errorf("Want status '%d', got '%d'", tt.statusCode, responseRecorder.Code)
			}

			if strings.TrimSpace(responseRecorder.Body.String()) != tt.want {
				t.Errorf("Want '%s', got '%s'", tt.want, responseRecorder.Body)
			}
		})
	}
}

func Test_apiStoreGet(t *testing.T) {
	tests := []struct {
		name       string
		method     string
		statusCode int
		key        string
		value      string
		want       string
	}{
		{
			name:       "api store get action get request",
			method:     http.MethodGet,
			key:        "test",
			value:      "12",
			want:       `{"status":"success","message":"test key found","result":"12"}`,
			statusCode: http.StatusOK,
		},
		{
			name:       "api store get action post request",
			method:     http.MethodPost,
			key:        "testsecond",
			value:      "11",
			want:       `{"status":"success","message":"testsecond key found","result":"11"}`,
			statusCode: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			store = NewStore()
			store.Set(tt.key, tt.value)

			request := httptest.NewRequest(tt.method, "/api/get?key="+tt.key, nil)
			if tt.method == "POST" {
				request := httptest.NewRequest(tt.method, "/api/get", strings.NewReader("key="+tt.key))
				request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
			}
			responseRecorder := httptest.NewRecorder()

			apiStoreGet().ServeHTTP(responseRecorder, request)

			if responseRecorder.Code != tt.statusCode {
				t.Errorf("Want status '%d', got '%d'", tt.statusCode, responseRecorder.Code)
			}

			if strings.TrimSpace(responseRecorder.Body.String()) != tt.want {
				t.Errorf("Want '%s', got '%s'", tt.want, responseRecorder.Body)
			}
		})
	}
}

func Test_apiStoreDel(t *testing.T) {
	tests := []struct {
		name       string
		method     string
		statusCode int
		key        string
		value      string
		want       string
	}{
		{
			name:       "api store del action get request",
			method:     http.MethodGet,
			key:        "test",
			value:      "12",
			want:       `{"status":"success","message":"test key deleted","result":null}`,
			statusCode: http.StatusOK,
		},
		{
			name:       "api store del action post request",
			method:     http.MethodPost,
			key:        "testsecond",
			value:      "11",
			want:       `{"status":"success","message":"testsecond key deleted","result":null}`,
			statusCode: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			store = NewStore()
			store.Set(tt.key, tt.value)

			request := httptest.NewRequest(tt.method, "/api/del?key="+tt.key, nil)
			if tt.method == "POST" {
				request := httptest.NewRequest(tt.method, "/api/del", strings.NewReader("key="+tt.key))
				request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
			}
			responseRecorder := httptest.NewRecorder()

			apiStoreDel().ServeHTTP(responseRecorder, request)

			if responseRecorder.Code != tt.statusCode {
				t.Errorf("Want status '%d', got '%d'", tt.statusCode, responseRecorder.Code)
			}

			if strings.TrimSpace(responseRecorder.Body.String()) != tt.want {
				t.Errorf("Want '%s', got '%s'", tt.want, responseRecorder.Body)
			}
		})
	}
}

func Test_apiStoreList(t *testing.T) {
	tests := []struct {
		name       string
		method     string
		statusCode int
		key        string
		value      string
		want       string
	}{
		{
			name:       "api store list action get request",
			method:     http.MethodGet,
			key:        "test",
			value:      "12",
			want:       `{"status":"success","message":"1 items found","result":{"test":"12"}}`,
			statusCode: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			store = NewStore()
			store.Set(tt.key, tt.value)

			request := httptest.NewRequest(tt.method, "/api/list", nil)
			responseRecorder := httptest.NewRecorder()

			apiStoreList().ServeHTTP(responseRecorder, request)

			if responseRecorder.Code != tt.statusCode {
				t.Errorf("Want status '%d', got '%d'", tt.statusCode, responseRecorder.Code)
			}

			if strings.TrimSpace(responseRecorder.Body.String()) != tt.want {
				t.Errorf("Want '%s', got '%s'", tt.want, responseRecorder.Body)
			}
		})
	}
}

func Test_apiStoreStats(t *testing.T) {
	tests := []struct {
		name       string
		method     string
		statusCode int
		key        string
		value      string
		want       string
	}{
		{
			name:       "api store stats action get request",
			method:     http.MethodGet,
			key:        "test",
			value:      "12",
			want:       `{"status":"success","message":"","result":"1"}`,
			statusCode: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			store = NewStore()
			store.Set(tt.key, tt.value)

			request := httptest.NewRequest(tt.method, "/api/stats", nil)
			responseRecorder := httptest.NewRecorder()

			apiStoreStats().ServeHTTP(responseRecorder, request)

			if responseRecorder.Code != tt.statusCode {
				t.Errorf("Want status '%d', got '%d'", tt.statusCode, responseRecorder.Code)
			}

			if strings.TrimSpace(responseRecorder.Body.String()) != tt.want {
				t.Errorf("Want '%s', got '%s'", tt.want, responseRecorder.Body)
			}
		})
	}
}

func Test_apiStoreFlush(t *testing.T) {
	tests := []struct {
		name       string
		method     string
		statusCode int
		key        string
		value      string
		want       string
	}{
		{
			name:       "api store flush action get request",
			method:     http.MethodGet,
			key:        "test",
			value:      "12",
			want:       `{"status":"success","message":"Flushed!","result":0}`,
			statusCode: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			store = NewStore()
			store.Set(tt.key, tt.value)

			request := httptest.NewRequest(tt.method, "/api/flush", nil)
			responseRecorder := httptest.NewRecorder()

			apiStoreFlush().ServeHTTP(responseRecorder, request)

			if responseRecorder.Code != tt.statusCode {
				t.Errorf("Want status '%d', got '%d'", tt.statusCode, responseRecorder.Code)
			}

			if strings.TrimSpace(responseRecorder.Body.String()) != tt.want {
				t.Errorf("Want '%s', got '%s'", tt.want, responseRecorder.Body)
			}
		})
	}
}

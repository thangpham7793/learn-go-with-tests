package urlshort

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCmd(t *testing.T) {
	yaml := []byte(`
- path: /redirect-me
  url: new-site-url
`)

	urlMap, _ := newUrlMapper(yaml)
	t.Run("should redirect if a match is found", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/redirect-me", nil)

		rr := httptest.NewRecorder()
		handler := YAMLHandler(urlMap, http.HandlerFunc(DefaultHandler))
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusMovedPermanently {
			t.Errorf("Wrong status code. Got %d, want %d", status, http.StatusMovedPermanently)
		}

		expected := "<a href=\"/new-site-url\">Moved Permanently</a>.\n\n"
		if body := rr.Body; body.String() != expected {
			t.Errorf("Wrong body. Got %q, want %q", body, expected)
		}
	})

	t.Run("should return bad reuqest if no matched url is found", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/redirect", nil)

		rr := httptest.NewRecorder()
		handler := YAMLHandler(urlMap, http.HandlerFunc(DefaultHandler))
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("Wrong status code. Got %d, want %d", status, http.StatusBadRequest)
		}

		expectedBody := "Please use a different shortened url"
		if body := rr.Body; body.String() != expectedBody {
			t.Errorf("Wrong body. Got %q, want %q", body, expectedBody)
		}
	})
}

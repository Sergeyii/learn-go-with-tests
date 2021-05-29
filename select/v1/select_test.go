package _select

import (
	"testing"
)

func TestWebsiteRacer(t *testing.T) {
	slowURL := "http://www.facebook.com"
	fastURL := "http://yandex.ru"

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got=%q, want=%q", got, want)
	}
}

package auth

import (
	"reflect"
	"testing"
)

func TestAuth(t *testing.T) {
	got, _ := GetAPIKey(map[string][]string{
		"Authorization": {"ApiKey a"},
	})
	want := "a"
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}

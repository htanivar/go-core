package sample

import "testing"

func TestWithHelper(t *testing.T) {
	helperFunction(t, "a parameter")
}

func helperFunction(t *testing.T, param string) string {
	t.Fatal("helperFunction setup failed")
	return param
}

func TestWithBetterHelper(t *testing.T) {
	betterHelperFunction(t, "file.json")
}

func betterHelperFunction(t *testing.T, s string) string {
	t.Helper()
	t.Fatal("betterHelperFunction setup failed")
	return ""
}

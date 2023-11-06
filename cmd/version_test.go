package cmd

import "testing"

func TestAppVersion(t *testing.T) {
	expected := "Version" // replace with your expected app version
	if appVersion != expected {
		t.Errorf("appVersion = %s; want %s", appVersion, expected)
	}
}

package conoha

import (
	"os"
	"testing"
)

func TestNewOpenStack(t *testing.T) {
	// 実際のOpenStackエンドポイントに接続するための環境変数が無い場合、
	// このテストは意味のある検証ができないのでスキップする。
	requiredEnv := []string{"OS_AUTH_URL", "OS_USERNAME", "OS_PASSWORD", "OS_TENANT_ID"}
	for _, key := range requiredEnv {
		if v := getenv(key); v == "" {
			t.Skipf("missing env %s; skipping live OpenStack test", key)
		}
	}

	os, err := NewOpenStack()
	if err != nil || os == nil {
		t.Errorf("%v", err)
	}

	if os.Compute == nil {
		t.Fatal("os.Compute should not be nil")
	}
	if os.Network == nil {
		t.Fatal("os.Network should not be nil")
	}
}

// getenv is defined to avoid importing os in production code paths unnecessarily in other files.
func getenv(key string) string {
	return os.Getenv(key)
}

//go:build unit_test
// +build unit_test

package data

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

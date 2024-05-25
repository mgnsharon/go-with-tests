package main

import (
	"fmt"
	"os"
	"path"
	"testing"
)

func TestCreatePackage(t *testing.T) {
	t.Run("create a folder", func(t *testing.T) {
		testPackage := "testfolder"
		defer cleanup(t, testPackage)

		err := CreatePackage(testPackage)
		if err != nil {
			t.Errorf("failed to create package %s: %v", testPackage, err)
		}
	})

	t.Run("create an empty go file", func(t *testing.T) {
		testPackage := "testfolder"
		defer cleanup(t, testPackage)

		err := CreatePackage(testPackage)
		if err != nil {
			t.Errorf("failed to create package %s: %v", testPackage, err)
		}
		assertFileExists(t, path.Join(testPackage, fmt.Sprintf("%s.go", testPackage)))
	})

	t.Run("create a test file", func(t *testing.T) {
		testPackage := "testfolder"
		defer cleanup(t, testPackage)

		err := CreatePackage(testPackage)
		if err != nil {
			t.Errorf("failed to create package %s: %v", testPackage, err)
		}
		assertFileExists(t, path.Join(testPackage, fmt.Sprintf("%s_test.go", testPackage)))
	})
}

func assertFileExists(t testing.TB, fp string) {
	_, err := os.Stat(fp)
	if err != nil {
		t.Errorf("expected %s to exist", fp)
	}
}

func cleanup(t *testing.T, f string) {
	t.Helper()
	err := os.RemoveAll(f)
	if err != nil {
		t.Errorf("unable to remove %v: %q", f, err)
	}
}

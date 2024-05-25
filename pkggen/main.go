package main

import (
	"fmt"
	"os"
	"path"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("package name is required")
		os.Exit(1)
	}
	pkg := os.Args[1]

	err := CreatePackage(pkg)
	if err != nil {
		fmt.Printf("unable to create package %s: %v", pkg, err)
		os.Exit(1)
	}
}

// CreatePackage creates a new package in a project if it doesn't exist
func CreatePackage(f string) error {
	err := os.Mkdir(f, 0o755)
	if err != nil {
		return fmt.Errorf("unable to create %s: %q", f, err)
	}
	c := fmt.Sprintf("package %s", f)
	err = os.WriteFile(path.Join(f, fmt.Sprintf("%s.go", f)), []byte(c), 0o755)
	if err != nil {
		return fmt.Errorf("unable to create %s: %q", f, err)
	}
	tfc := fmt.Sprintf(`package %s
import(
	"testing"
)

func Test%s(t *testing.T){
	t.Run("%s", func(t *testing.T){

	})
}`, f, cases.Title(language.AmericanEnglish).String(f), f)
	err = os.WriteFile(path.Join(f, fmt.Sprintf("%s_test.go", f)), []byte(tfc), 0o755)
	if err != nil {
		return fmt.Errorf("unable to create %s: %q", f, err)
	}
	return nil
}

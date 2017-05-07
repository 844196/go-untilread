package untilread

import (
	"errors"
	"reflect"
	"strings"
	"testing"
)

func TestDo(t *testing.T) {
	src := "こんにちは、ジャパリパーク。さようなら、世界。"
	sep := "。"

	expected := []string{"こんにちは、ジャパリパーク。", "さようなら、世界。"}
	var actual []string

	ir := strings.NewReader(src)
	err := Do(ir, sep, func(s string) error {
		actual = append(actual, s)
		return nil
	})

	if err != nil {
		t.Fatalf("unexpected error: %q", err)
	}

	if reflect.DeepEqual(expected, actual) == false {
		t.Fatalf("want: %q, got: %q", expected, actual)
	}
}

func TestDoContainesLf(t *testing.T) {
	src := "こんにちは、\nジャパリパーク。\nさようなら、世界。\n"
	sep := "。"

	expected := []string{"こんにちは、\nジャパリパーク。", "\nさようなら、世界。"}
	var actual []string

	ir := strings.NewReader(src)
	err := Do(ir, sep, func(s string) error {
		actual = append(actual, s)
		return nil
	})

	if err != nil {
		t.Fatalf("unexpected error: %q", err)
	}

	if reflect.DeepEqual(expected, actual) == false {
		t.Fatalf("want: %q, got: %q", expected, actual)
	}
}

func TestDoSeparatorOnly(t *testing.T) {
	src := "。"
	sep := "。"

	expected := []string{"。"}
	var actual []string

	ir := strings.NewReader(src)
	err := Do(ir, sep, func(s string) error {
		actual = append(actual, s)
		return nil
	})

	if err != nil {
		t.Fatalf("unexpected error: %q", err)
	}

	if reflect.DeepEqual(expected, actual) == false {
		t.Fatalf("want: %q, got: %q", expected, actual)
	}
}

func TestDoBufferEmpty(t *testing.T) {
	src := ""
	sep := "。"

	ir := strings.NewReader(src)
	isBlockCalled := false
	err := Do(ir, sep, func(s string) error {
		isBlockCalled = true
		return nil
	})

	if err != nil {
		t.Fatalf("unexpected error: %q", err)
	}

	if isBlockCalled {
		t.Fatal("expected anonymous function not called, but called")
	}
}

func TestDoNotExistsSeparator(t *testing.T) {
	src := "こんにちは、ジャパリパーク。さようなら、世界"
	sep := "。"

	expected := []string{"こんにちは、ジャパリパーク。"}
	var actual []string

	ir := strings.NewReader(src)
	err := Do(ir, sep, func(s string) error {
		actual = append(actual, s)
		return nil
	})

	if err != nil {
		t.Fatalf("unexpected error: %q", err)
	}

	if reflect.DeepEqual(expected, actual) == false {
		t.Fatalf("want: %q, got: %q", expected, actual)
	}
}

func TestDoFailed(t *testing.T) {
	src := "こんにちは、ジャパリパーク。さようなら、世界。"
	sep := "。"

	ir := strings.NewReader(src)
	expectedError := errors.New("TEST ERROR")
	err := Do(ir, sep, func(s string) error {
		return expectedError
	})

	if err == nil {
		t.Fatalf("expected error: %q", expectedError)
	}
}

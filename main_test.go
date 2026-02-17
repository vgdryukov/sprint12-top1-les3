package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	// Тестируем вывод main
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	io.Copy(&buf, r)

	expected := "Hello, word!\n"
	if buf.String() != expected {
		t.Errorf("Expected %q, got %q", expected, buf.String())
	}
}

func TestGetGreeting(t *testing.T) {
	// Тестируем логику
	expected := "Hello, word!"
	result := getGreeting()
	if result != expected {
		t.Errorf("getGreeting() = %q, want %q", result, expected)
	}
}

func BenchmarkGetGreeting(b *testing.B) {
	// Бенчмарк логики
	for i := 0; i < b.N; i++ {
		getGreeting()
	}
}

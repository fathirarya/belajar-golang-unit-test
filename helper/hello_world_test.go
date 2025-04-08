package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Table Test
func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Fathir",
			request:  "Fathir",
			expected: "Hello Fathir",
		},
		{
			name:     "Arya",
			request:  "Arya",
			expected: "Hello Arya",
		},
		{
			name:     "Nafis",
			request:  "Nafis",
			expected: "Hello Nafis",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

// Sub Test
func TestSubTest(t *testing.T) {
	t.Run("Fathir", func(t *testing.T) {
		result := HelloWorld("Fathir")
		require.Equal(t, "Hello Fathir", result, "Result must be 'Hello Fathir'")
	})

	t.Run("Arya", func(t *testing.T) {
		result := HelloWorld("Arya")
		require.Equal(t, "Hello Arya", result, "Result must be 'Hello Arya'")
	})
}

// Before and After test use Test Main
func TestMain(m *testing.M) {
	// before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	// after
	fmt.Println("AFTER UNIT TEST")
}

// Skip Test
func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Can not run on Windows")
	}

	result := HelloWorld("Fathir")
	require.Equal(t, "Hello Fathir", result, "Result must be 'Hello Fathir'")
}

// Test use Require
func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Fathir")
	require.Equal(t, "Hello Fathir", result, "Result must be 'Hello Fathir'")
	fmt.Println("TestHelloWorl with Require Done")
}

// Test use Assert
func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Fathir")
	assert.Equal(t, "Hello Fathir", result, "Result must be 'Hello Fathir'")
	fmt.Println("TestHelloWorl with Assert Done")
}

// Test use function Error
func TestHelloWorldFathir(t *testing.T) {
	result := HelloWorld("Fathir")

	if result != "Hello Fathir" {
		// error
		t.Error("Result Must Be 'Hello Fathir")
	}
	fmt.Println("TestHelloWorldFathir Done")
}

// Test use function Fatal
func TestHelloWorldArya(t *testing.T) {
	result := HelloWorld("Arya")

	if result != "Hello Arya" {
		// fatal
		t.Fatal("Result Must Be 'Hello Arya' ")
	}
	fmt.Println("TestHelloWorldArya Done")
}

func TestHelloWorldNafis(t *testing.T) {
	result := HelloWorld("Nafis")

	if result != "Hello Nafis" {
		// error
		t.Fatal("Result Must Be 'Hello Nafis' ")
	}
	fmt.Println("TestHelloWorldNafis Done")
}

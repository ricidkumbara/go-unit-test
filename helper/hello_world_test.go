package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	fmt.Println("Before Unit Test")

	m.Run()

	fmt.Println("After Unit Test")
}

func TestHelloWorldRicid(t *testing.T) {
	result := HelloWorld("Ricid")

	if result != "Hello Ricid" {
		t.Error("result must be 'Hello Ricid'")
	}

	fmt.Println("TestHelloWorldRicid done")
}

func TestHelloWorldKumbara(t *testing.T) {
	result := HelloWorld("Kumbara")

	if result != "Hello Kumbara" {
		// error
		t.Fatal("Result must be 'Hello Kumbara'")
	}

	fmt.Println("TestHelloWorldKumbara Done")
}

func TestHelloWorldAssertion(t *testing.T) {
	result := HelloWorld("World")

	assert.Equal(t, "Hello World", result, "Failed")
	fmt.Println("TestHelloWorldAssertion Done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("World")

	require.Equal(t, "Hello World", result, "Failed")
	fmt.Println("TestHelloWorldRequire Done")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Cannot run on Windows OS")
	}

	fmt.Println("This line will skipped if OS is Windows")
}

func TestSubtest(t *testing.T) {
	t.Run("Ricid", func(t *testing.T) {
		result := HelloWorld("Ricid")
		assert.Equal(t, "Hello Ricid", result, "Failed")
	})
	t.Run("Kumbara", func(t *testing.T) {
		result := HelloWorld("Kumbara")
		assert.Equal(t, "Hello Kumbara", result, "Failed")
	})
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Ricid",
			request:  "Ricid",
			expected: "Hello Ricid",
		},
		{
			name:     "Kumbara",
			request:  "Kumbara",
			expected: "Hello Kumbara",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

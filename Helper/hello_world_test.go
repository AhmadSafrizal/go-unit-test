package helper

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
)

func BenchmarkTable(b *testing.B) {
	benchmark := []struct {
		name string
		request string
	}{
		{
			name: "Saf",
			request: "Saf",
		},
		{
			name: "Ri",
			request: "Ri",
		},
		{
			name: "Zal",
			request: "Zal",
		},
		{
			name: "Ahmad Safrizal",
			request: "Ahmad Safrizal",
		},
	}

	for _, benchmark := range benchmark {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Ahmad", func (b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Ahmad")
		}
	})
	
	b.Run("Rizal", func (b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Rizal")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Saf")
	}
}

func BenchmarkHelloWorldRi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Ri")
	}
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name string
		request string
		expected string
	}{
		{
			name: "Saf",
			request: "Saf",
			expected: "Hello Saf",
		},
		{
			name: "Ri",
			request: "Ri",
			expected: "Hello Ri",
		},
		{
			name: "Zal",
			request: "Zal",
			expected: "Hello Zal",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Saf", func(t *testing.T) {
		result := HelloWorld("Saf")
		require.Equal(t, "Hello Saf", result, "Result must be 'Hello Saf'")
	})
	t.Run("Ri", func(t *testing.T) {
		result := HelloWorld("Ri")
		require.Equal(t, "Hello Ri", result, "Result must be 'Hello Ri'")
	})
}

func TestMain(m *testing.M) {
	// Before
	fmt.Println("Before Unit Test")

	m.Run()

	//After
	fmt.Println("After Unit Test")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Can't run on mac OS")
	}

	result := HelloWorld("Saf")
	require.Equal(t, "Hello Saf", result, "Result must be 'Hello Saf'")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Saf")
	require.Equal(t, "Hello Saf", result, "Result must be 'Hello Saf'")
	fmt.Println("TestHelloWorld with Assert Done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Saf")
	assert.Equal(t, "Hello Saf", result, "Result must be 'Hello Saf'")
	fmt.Println("TestHelloWorld with Assert Done")
}

func TestHelloWorldSaf(t *testing.T) {
	result := HelloWorld("Saf")

	if result != "Hello Saf" {
		// error
		t.Error("Result must be Saf")
	}

	fmt.Println("TestHelloWorldSaf Done")
}

func TestHelloWorldRi(t *testing.T) {
	result := HelloWorld("Ri")

	if result != "Hello Ri" {
		// error
		t.Fatal("Result must be Hello Ri")
	}

	fmt.Println("TestHelloWorldRi Done")
}

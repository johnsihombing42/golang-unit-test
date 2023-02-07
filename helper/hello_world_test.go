package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorld(t *testing.T){
	result:=HelloWorld("John")

	if result != "Hello John" {
		t.Error("Result must be Hello John")
	}
}


func TestHelloWorldAssert(t *testing.T){
	result:=HelloWorld("John")
	assert.Equal(t,"Hello John",result,"Result must be Hello John")
}

func TestSkip(t *testing.T){
	if runtime.GOOS=="darwin"{
		t.Skip("Can not run on mac os")
	}
	result:=HelloWorld("John")
	require.Equal(t,"Hello John",result,"Result must be Hello John")
}

func TestMain(m *testing.M){
	//before
	fmt.Println("Before Unit Test")
	m.Run()
	//After
	fmt.Println("After Unit Test")
}

func TestSubTest(t *testing.T){
	t.Run("John",func(t *testing.T) {
		result:=HelloWorld("John")
		require.Equal(t,"Hello John",result,"Result must be Hello John")
	})
	t.Run("Putra",func(t *testing.T) {
		result:=HelloWorld("Putra")
		require.Equal(t,"Hello Putra",result,"Result must be Hello Putra")
	})
}


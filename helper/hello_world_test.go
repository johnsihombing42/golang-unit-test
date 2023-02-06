package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
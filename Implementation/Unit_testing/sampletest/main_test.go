//reference := www.thepolyglotdeveloper.com
package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	total:= Add(11,22)
	assert.NotNil(t,total,"The `total` should not be `nil`")
	assert.Equal(t,33,total,"expecting 33")

}

func TestSub(t *testing.T) {
	total :=Sub(22,11)
	assert.NotNil(t,total,"The total should not be nil")
	assert.Equal(t,11,total,"ecpecting 11")
}